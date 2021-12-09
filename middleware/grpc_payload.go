package middleware

import (
	"bytes"
	"encoding/binary"
	"errors"
	"fmt"
	"github.com/golang/protobuf/proto"
	"github.com/golang/protobuf/protoc-gen-go/descriptor"
	"io"
	"io/ioutil"
	"net/http"
	"regexp"
	"strconv"
	"strings"

	"github.com/jhump/protoreflect/desc"
	"github.com/jhump/protoreflect/desc/builder"
	"github.com/jhump/protoreflect/desc/protoprint"
	"github.com/jhump/protoreflect/dynamic"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// GRPCPayloadCompressionFormat tells by reading the first byte compressed or not
type GRPCPayloadCompressionFormat uint8

const (
	compressionNone GRPCPayloadCompressionFormat = 0
	compressionMade GRPCPayloadCompressionFormat = 1
	maxInt                                       = int(^uint(0) >> 1)

	NestedArray = "MessageArray"
	StringArray = "StringArray"
	String      = "String"
	Message     = "Message"
)

type GRPCPayloadHandler struct{}


func (b GRPCPayloadHandler) Extract(req *http.Request, protoIndex string) (string, error) {
	reqBody, err := ioutil.ReadAll(req.Body)
	if err != nil {
		return "", err
	}
	defer req.Body.Close()
	// repopulate body
	req.Body = ioutil.NopCloser(bytes.NewBuffer(reqBody))

	return b.extractFromRequest(reqBody, protoIndex)
}

func (b GRPCPayloadHandler) extractFromRequest(body []byte, protoIndex string) (string, error) {
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
//4.2[1]
func fieldFromProtoMessage(msg []byte, tagIndex string) (string, error) {
	q, err := NewQuery(tagIndex)
	if err != nil {
		return "", err
	}

	msgDescriptor, err := buildPayloadProto(q)
	if err != nil {
		return "", err
	}

	// populate message
	dynamicMsgKey := dynamic.NewMessage(msgDescriptor)
	protoPrinter := &protoprint.Printer{}
	str, err := protoPrinter.PrintProtoToString(msgDescriptor.GetFile())
	fmt.Println(str)
	fmt.Println(err)
	if err := dynamicMsgKey.Unmarshal(msg); err != nil {
		return "", err
	}

	val, err := QueryValue(dynamicMsgKey, q)
	if err != nil {
		return "", err
	}

	return val.(string), nil
}

func buildPayloadProto(queries []Query) (*desc.MessageDescriptor, error) {
	builderMsg := builder.NewMessage("shield")
	fieldName := "_"
	var lastBuilderMsg *builder.MessageBuilder
	for queryListIndex := len(queries) - 1; queryListIndex > 0; queryListIndex-- {
		subQuery := queries[queryListIndex]
		if subQuery.DataType == String {
			lastBuilderMsg = builder.NewMessage(fmt.Sprintf("msg_%s%d", fieldName, subQuery.Field)).AddField(builder.NewField(fmt.Sprintf("%s%d", fieldName, subQuery.Field),
				builder.FieldTypeScalar(descriptor.FieldDescriptorProto_TYPE_STRING)).SetNumber(int32(subQuery.Field)))
		} else if subQuery.DataType == StringArray {
			lastBuilderMsg = builder.NewMessage(fmt.Sprintf("msg_%s%d", fieldName, subQuery.Field)).AddField(builder.NewField(fmt.Sprintf("%s%d", fieldName, subQuery.Field),
				builder.FieldTypeScalar(descriptor.FieldDescriptorProto_TYPE_STRING)).SetNumber(int32(subQuery.Field)).SetRepeated())
		} else if subQuery.DataType == Message {
			lastBuilderMsg = builder.NewMessage(fmt.Sprintf("msg_%s%d", fieldName, subQuery.Field)).AddField(builder.NewField(fmt.Sprintf("%s%d", fieldName, subQuery.Field),
				builder.FieldTypeMessage(lastBuilderMsg)).SetNumber(int32(subQuery.Field)))
		} else if subQuery.DataType == NestedArray {
			lastBuilderMsg = builder.NewMessage(fmt.Sprintf("msg_%s%d", fieldName, subQuery.Field)).AddField(builder.NewField(fmt.Sprintf("%s%d", fieldName, subQuery.Field),
				builder.FieldTypeMessage(lastBuilderMsg)).SetNumber(int32(subQuery.Field)).SetRepeated())
		}
		fieldName = "_" + fieldName
	}
	//
	//if queries[0].DataType == Message {
	//	builderMsg.AddField(builder.NewField(fmt.Sprintf("%s%d", fieldName, queries[0].Field),
	//		builder.FieldTypeMessage(lastBuilderMsg)).SetNumber(int32(queries[0].Field)))
	//
	//} else if queries[0].DataType == NestedArray {
	//	builderMsg.AddField(builder.NewField(fmt.Sprintf("%s%d", fieldName, queries[0].Field),
	//		builder.FieldTypeMessage(lastBuilderMsg)).SetNumber(int32(queries[0].Field)).SetRepeated())
	//
	//}

	nestedMsg := builder.NewMessage("sub_message").AddField(builder.NewField("b1", builder.FieldTypeScalar(descriptor.FieldDescriptorProto_TYPE_STRING)).SetNumber(1))

	msgD, err:= nestedMsg.Build()
	fmt.Println(err)
	protoPrinter := &protoprint.Printer{}
	str, err := protoPrinter.PrintProtoToString(msgD.GetFile())
	fmt.Println(str)
	fmt.Println(err)

	builderMsg = builderMsg.AddField(builder.NewField("a2", builder.FieldTypeMessage(nestedMsg)).SetNumber(4))

	//builderMsg.AddField(builder.NewField("_1", builder.FieldTypeScalar(descriptor.FieldDescriptorProto_TYPE_STRING)).SetNumber(1))
	//builderMsg.AddField(builder.NewField("_2", builder.FieldTypeScalar(descriptor.FieldDescriptorProto_TYPE_STRING)).SetRepeated().SetNumber(2))

	msgDescriptor, err := builderMsg.Build()
	if err != nil {
		return nil, err
	}
	return msgDescriptor, nil
}

//--------------------------------------------------------

type Query struct {
	Field    int         `json:"field"`
	Index    int         `json:"index"`
	MapKey   interface{} `json:"map_key"`
	DataType string      `json:"data_type"`
}

func QueryValue(msg *dynamic.Message, query []Query) (interface{}, error) {
	first := query[0]
	rest := query[1:]

	var val interface{}
	var err error
	if first.Index != -1 {
		val, err = msg.TryGetRepeatedFieldByNumber(first.Field, first.Index)
		fmt.Println("here1")
	} else if first.MapKey != nil {
		val, err = msg.TryGetMapFieldByNumber(first.Field, first.MapKey)
		fmt.Println("here2")
	} else {
		val, err = msg.TryGetFieldByNumber(first.Field)
	}
	if err != nil {
		return nil, err
	}

	if len(rest) == 0 {
		return val, nil
	}

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

	return QueryValue(dm, rest)
}

func NewQuery(query string) ([]Query, error) {
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
				processingQuery.DataType = NestedArray
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
