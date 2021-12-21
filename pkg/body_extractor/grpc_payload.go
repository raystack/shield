package body_extractor

import (
	"bytes"
	"encoding/binary"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"reflect"
	"regexp"
	"strconv"
	"strings"
	"sync"

	"github.com/golang/protobuf/protoc-gen-go/descriptor"

	"github.com/golang/protobuf/proto"
	"github.com/jhump/protoreflect/desc"
	"github.com/jhump/protoreflect/desc/builder"
	"github.com/jhump/protoreflect/dynamic"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// GRPCPayloadCompressionFormat tells by reading the first byte compressed or not
type GRPCPayloadCompressionFormat uint8

const (
	compressionNone GRPCPayloadCompressionFormat = 0
	compressionMade GRPCPayloadCompressionFormat = 1

	maxInt = int(^uint(0) >> 1)

	MessageArray = "MessageArray"
	StringArray  = "StringArray"
	String       = "String"
	Message      = "Message"
)

type queryCacheMutex struct {
	m          sync.Mutex
	queryCache map[string][]Query
}

func (q *queryCacheMutex) Get(key string) ([]Query, bool) {
	q.m.Lock()
	defer q.m.Unlock()
	val, ok := q.queryCache[key]
	return val, ok
}

func (q *queryCacheMutex) Set(key string, query []Query) {
	q.m.Lock()
	defer q.m.Unlock()
	q.queryCache[key] = query
}

type payloadProtoCacheMutex struct {
	m                 sync.Mutex
	payloadProtoCache map[string]*desc.MessageDescriptor
}

func (p *payloadProtoCacheMutex) Get(key string) (*desc.MessageDescriptor, bool) {
	p.m.Lock()
	defer p.m.Unlock()
	val, ok := p.payloadProtoCache[key]
	return val, ok
}

func (p *payloadProtoCacheMutex) Set(key string, msgDescriptor *desc.MessageDescriptor) {
	p.m.Lock()
	defer p.m.Unlock()
	p.payloadProtoCache[key] = msgDescriptor
}

var (
	queryCache        = queryCacheMutex{queryCache: make(map[string][]Query)}
	payloadProtoCache = payloadProtoCacheMutex{payloadProtoCache: make(map[string]*desc.MessageDescriptor)}
)

type Query struct {
	Field    int    `json:"field"`
	Index    int    `json:"index"`
	DataType string `json:"data_type"`
}

type GRPCPayloadHandler struct {
	grpcDisabled bool
}

func (b GRPCPayloadHandler) Extract(body *io.ReadCloser, protoIndex string) (interface{}, error) {
	reqBody, err := ioutil.ReadAll(*body)
	if err != nil {
		return "", err
	}
	defer (*body).Close()
	// repopulate body
	*body = ioutil.NopCloser(bytes.NewBuffer(reqBody))

	return b.extractFromRequest(reqBody, protoIndex)
}

func (b GRPCPayloadHandler) extractFromRequest(body []byte, protoIndex string) (interface{}, error) {
	if b.grpcDisabled {
		return fieldFromProtoMessage(body, protoIndex)
	}

	reqParser := grpcRequestParser{
		r:      bytes.NewBuffer(body),
		header: [5]byte{},
	}
	pf, msg, err := reqParser.Parse()
	if err != nil {
		return "", err
	}
	if pf == compressionMade {
		// unsupported for now
		return "", errors.New("compressed message, unsupported grpc feature")
	}

	return fieldFromProtoMessage(msg, protoIndex)
}

// grpcRequestParser reads complete gRPC messages from the underlying reader
type grpcRequestParser struct {
	// r is the underlying reader
	r io.Reader

	// The header of a gRPC message. Find more detail at
	// https://github.com/grpc/grpc/blob/master/doc/PROTOCOL-HTTP2.md
	header [5]byte
}

// Parse reads a complete gRPC message from the stream.
//
// It returns the message and its payload (compression/encoding)
// format. The caller owns the returned msg memory.
//
// If there is an error, possible values are:
//   * io.EOF, when no messages remain
//   * io.ErrUnexpectedEOF
//   * of type transport.ConnectionError
//   * an error from the status package
// No other error values or types must be returned, which also means
// that the underlying io.Reader must not return an incompatible
// error.
func (p *grpcRequestParser) Parse() (pf GRPCPayloadCompressionFormat, msg []byte, err error) {
	if _, err := p.r.Read(p.header[:]); err != nil {
		return 0, nil, err
	}
	// first byte is for compressed or not
	pf = GRPCPayloadCompressionFormat(p.header[0])
	// next 4 bytes is for length of message
	length := binary.BigEndian.Uint32(p.header[1:])
	if length == 0 {
		return pf, nil, nil
	}
	if int64(length) > int64(maxInt) {
		return 0, nil, status.Errorf(codes.ResourceExhausted, "grpc: received message larger than max length allowed on current machine (%d vs. %d)", length, maxInt)
	}

	msg = make([]byte, int(length))
	if _, err := p.r.Read(msg); err != nil {
		if err == io.EOF {
			err = io.ErrUnexpectedEOF
		}
		return 0, nil, err
	}
	return pf, msg, nil
}

func fieldFromProtoMessage(msg []byte, tagIndex string) (interface{}, error) {
	parsedQuery, err := ParseQuery(tagIndex)
	if err != nil {
		return "", err
	}

	msgDesc, err := buildPayloadProto(tagIndex, parsedQuery)
	if err != nil {
		return "", err
	}

	// populate message
	dynamicMsgKey := dynamic.NewMessage(msgDesc)
	if err := dynamicMsgKey.Unmarshal(msg); err != nil {
		return "", err
	}

	val, err := RunQuery(dynamicMsgKey, parsedQuery)
	if err != nil {
		return "", err
	}
	return val, nil
}

func buildPayloadProto(query string, queries []Query) (*desc.MessageDescriptor, error) {
	if val, ok := payloadProtoCache.Get(query); ok {
		return val, nil
	}

	builderMsg := builder.NewMessage("shield")
	fieldName := "_"
	var lastBuilderMsg *builder.MessageBuilder
	for queryListIndex := len(queries) - 1; queryListIndex > 0; queryListIndex-- {
		subQuery := queries[queryListIndex]
		if subQuery.DataType == String {
			lastBuilderMsg = builder.NewMessage(fmt.Sprintf("msg%s%d", fieldName, subQuery.Field)).AddField(builder.NewField(fmt.Sprintf("field%s%d", fieldName, subQuery.Field),
				builder.FieldTypeScalar(descriptor.FieldDescriptorProto_TYPE_STRING)).SetNumber(int32(subQuery.Field)))
		} else if subQuery.DataType == StringArray {
			lastBuilderMsg = builder.NewMessage(fmt.Sprintf("msg%s%d", fieldName, subQuery.Field)).AddField(builder.NewField(fmt.Sprintf("field%s%d", fieldName, subQuery.Field),
				builder.FieldTypeScalar(descriptor.FieldDescriptorProto_TYPE_STRING)).SetNumber(int32(subQuery.Field)).SetRepeated())
		} else if subQuery.DataType == Message {
			lastBuilderMsg = builder.NewMessage(fmt.Sprintf("msg%s%d", fieldName, subQuery.Field)).AddField(builder.NewField(fmt.Sprintf("field%s%d", fieldName, subQuery.Field),
				builder.FieldTypeMessage(lastBuilderMsg)).SetNumber(int32(subQuery.Field)))
		} else if subQuery.DataType == MessageArray {
			lastBuilderMsg = builder.NewMessage(fmt.Sprintf("msg%s%d", fieldName, subQuery.Field)).AddField(builder.NewField(fmt.Sprintf("field%s%d", fieldName, subQuery.Field),
				builder.FieldTypeMessage(lastBuilderMsg)).SetNumber(int32(subQuery.Field)).SetRepeated())
		}
		fieldName = "_" + fieldName
	}

	if queries[0].DataType == Message {
		builderMsg.AddField(builder.NewField(fmt.Sprintf("field%s%d", fieldName, queries[0].Field),
			builder.FieldTypeMessage(lastBuilderMsg)).SetNumber(int32(queries[0].Field)))
	} else if queries[0].DataType == MessageArray {
		builderMsg.AddField(builder.NewField(fmt.Sprintf("field%s%d", fieldName, queries[0].Field),
			builder.FieldTypeMessage(lastBuilderMsg)).SetNumber(int32(queries[0].Field)).SetRepeated())
	} else if queries[0].DataType == String {
		builderMsg.AddField(builder.NewField(fmt.Sprintf("field%s%d", fieldName, queries[0].Field),
			builder.FieldTypeScalar(descriptor.FieldDescriptorProto_TYPE_STRING)).SetNumber(int32(queries[0].Field)))
	} else if queries[0].DataType == StringArray {
		builderMsg.AddField(builder.NewField(fmt.Sprintf("field%s%d", fieldName, queries[0].Field),
			builder.FieldTypeScalar(descriptor.FieldDescriptorProto_TYPE_STRING)).SetNumber(int32(queries[0].Field)).SetRepeated())
	}

	msgDescriptor, err := builderMsg.Build()
	if err != nil {
		return nil, err
	}

	payloadProtoCache.Set(query, msgDescriptor)
	return msgDescriptor, nil
}

func RunQuery(msg *dynamic.Message, query []Query) (interface{}, error) {
	first := query[0]
	rest := query[1:]

	var val interface{}
	var err error
	if first.Index != -1 {
		if first.DataType == StringArray {
			val, err = msg.TryGetFieldByNumber(first.Field)
		} else if first.DataType == MessageArray {
			v, err := msg.TryGetFieldByNumber(first.Field)
			if err != nil {
				return nil, err
			}

			nestedMessageArray := v.([]interface{})
			valueList := make([]interface{}, 0)
			for _, j := range nestedMessageArray {
				dm, err := getNextDynamicMessage(j, first)
				if err != nil {
					return nil, err
				}
				valuesFromSubQuery, err := RunQuery(dm, rest)
				if err != nil {
					return nil, err
				}

				if reflect.TypeOf(valuesFromSubQuery).String() == "[]interface {}" {
					for _, interfaceValue := range valuesFromSubQuery.([]interface{}) {
						valueList = append(valueList, interfaceValue)
					}
				} else {
					valueList = append(valueList, valuesFromSubQuery)
				}
			}
			val = valueList
			return val, nil
		}
	} else {
		val, err = msg.TryGetFieldByNumber(first.Field)
	}
	if err != nil {
		return nil, err
	}

	if len(rest) == 0 {
		return val, nil
	}

	dm, err := getNextDynamicMessage(val, first)
	if err != nil {
		return nil, err
	}

	return RunQuery(dm, rest)
}

func getNextDynamicMessage(val interface{}, first Query) (*dynamic.Message, error) {
	dm, ok := val.(*dynamic.Message)
	if !ok {
		pm, ok := val.(proto.Message)
		if !ok {
			return nil, fmt.Errorf("cannot query field from non-message value %q", first.Field)
		}
		md, err := desc.LoadMessageDescriptorForMessage(pm)
		if err != nil {
			return nil, err
		}
		dm = dynamic.NewMessage(md)
		if err := dm.ConvertFrom(pm); err != nil {
			return nil, err
		}
	}
	return dm, nil
}

func ParseQuery(query string) ([]Query, error) {
	if queries, ok := queryCache.Get(query); ok {
		return queries, nil
	}

	arrayBrackets := []string{"[", "]"}
	arrayRegexBuild, err := regexp.Compile("\\[([0-9]*)]")
	if err != nil {
		return []Query{}, err
	}

	queries := make([]Query, 0)
	subQueries := strings.Split(query, ".")
	for queryIndex, subQueryString := range subQueries {
		processingQuery := Query{}
		if strings.Contains(subQueryString, arrayBrackets[0]) {
			r := removeBrackets(arrayRegexBuild.FindString(subQueryString), arrayBrackets)
			parsedIndex, err := strconv.Atoi(r)
			if err != nil {
				return []Query{}, err
			}
			processingQuery.Index = parsedIndex

			parsedfield, err := strconv.Atoi(subQueryString[:strings.Index(subQueryString, "[")])
			if err != nil {
				return []Query{}, err
			}
			processingQuery.Field = parsedfield

			if queryIndex == len(subQueries)-1 {
				processingQuery.DataType = StringArray
			} else {
				processingQuery.DataType = MessageArray
			}
		} else {
			processingQuery.Index = -1

			parsedfield, err := strconv.Atoi(subQueryString)
			if err != nil {
				return []Query{}, err
			}
			processingQuery.Field = parsedfield

			if queryIndex == len(subQueries)-1 {
				processingQuery.DataType = String
			} else {
				processingQuery.DataType = Message
			}
		}
		queries = append(queries, processingQuery)
	}

	queryCache.Set(query, queries)
	return queries, nil
}

func removeBrackets(capture string, bracket []string) string {
	return strings.Replace(
		strings.Replace(capture, bracket[0], "", 1),
		bracket[1],
		"",
		1,
	)
}
