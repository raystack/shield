package v1beta1

import (
	"context"

	"github.com/raystack/frontier/pkg/metadata"

	grpczap "github.com/grpc-ecosystem/go-grpc-middleware/logging/zap/ctxzap"
	"github.com/raystack/frontier/billing/product"
	frontierv1beta1 "github.com/raystack/frontier/proto/v1beta1"
)

type ProductService interface {
	GetByID(ctx context.Context, id string) (product.Product, error)
	Create(ctx context.Context, product product.Product) (product.Product, error)
	Update(ctx context.Context, product product.Product) (product.Product, error)
	List(ctx context.Context, filter product.Filter) ([]product.Product, error)
}

func (h Handler) CreateProduct(ctx context.Context, request *frontierv1beta1.CreateProductRequest) (*frontierv1beta1.CreateProductResponse, error) {
	logger := grpczap.Extract(ctx)

	metaDataMap := metadata.Build(request.GetBody().GetMetadata().AsMap())
	// parse price
	var productPrices []product.Price
	for _, v := range request.GetBody().GetPrices() {
		productPrices = append(productPrices, product.Price{
			Name:             v.GetName(),
			Amount:           v.GetAmount(),
			Currency:         v.GetCurrency(),
			UsageType:        product.BuildPriceUsageType(v.GetUsageType()),
			BillingScheme:    product.BuildBillingScheme(v.GetBillingScheme()),
			MeteredAggregate: v.GetMeteredAggregate(),
			Interval:         v.GetInterval(),
			Metadata:         metadata.Build(v.GetMetadata().AsMap()),
		})
	}
	// parse features
	var productFeatures []product.Feature
	for _, v := range request.GetBody().GetFeatures() {
		productFeatures = append(productFeatures, product.Feature{
			Name:       v.GetName(),
			ProductIDs: v.GetProductIds(),
			Metadata:   metadata.Build(v.GetMetadata().AsMap()),
		})
	}

	newProduct, err := h.productService.Create(ctx, product.Product{
		PlanIDs:      []string{request.GetBody().GetPlanId()},
		Name:         request.GetBody().GetName(),
		Title:        request.GetBody().GetTitle(),
		Description:  request.GetBody().GetDescription(),
		Prices:       productPrices,
		CreditAmount: request.GetBody().GetCreditAmount(),
		Behavior:     product.Behavior(request.GetBody().GetBehavior()),
		Features:     productFeatures,
		Metadata:     metaDataMap,
	})
	if err != nil {
		logger.Error(err.Error())
		return nil, grpcInternalServerError
	}

	productPB, err := transformProductToPB(newProduct)
	if err != nil {
		logger.Error(err.Error())
		return nil, grpcInternalServerError
	}

	return &frontierv1beta1.CreateProductResponse{
		Product: productPB,
	}, nil
}

func (h Handler) UpdateProduct(ctx context.Context, request *frontierv1beta1.UpdateProductRequest) (*frontierv1beta1.UpdateProductResponse, error) {
	logger := grpczap.Extract(ctx)

	metaDataMap := metadata.Build(request.GetBody().GetMetadata().AsMap())
	// parse price
	var productPrices []product.Price
	for _, v := range request.GetBody().GetPrices() {
		productPrices = append(productPrices, product.Price{
			Name:     v.GetName(),
			Metadata: metadata.Build(v.GetMetadata().AsMap()),
		})
	}
	// parse features
	var productFeatures []product.Feature
	for _, v := range request.GetBody().GetFeatures() {
		productFeatures = append(productFeatures, product.Feature{
			Name:       v.GetName(),
			ProductIDs: v.GetProductIds(),
			Metadata:   metadata.Build(v.GetMetadata().AsMap()),
		})
	}
	updatedProduct, err := h.productService.Update(ctx, product.Product{
		ID:          request.GetId(),
		Name:        request.GetBody().GetName(),
		Title:       request.GetBody().GetTitle(),
		Description: request.GetBody().GetDescription(),
		Behavior:    product.Behavior(request.GetBody().GetBehavior()),
		Prices:      productPrices,
		Features:    productFeatures,
		Metadata:    metaDataMap,
	})
	if err != nil {
		logger.Error(err.Error())
		return nil, grpcInternalServerError
	}
	productPb, err := transformProductToPB(updatedProduct)
	if err != nil {
		logger.Error(err.Error())
		return nil, grpcInternalServerError
	}

	return &frontierv1beta1.UpdateProductResponse{
		Product: productPb,
	}, nil
}

func (h Handler) ListProducts(ctx context.Context, request *frontierv1beta1.ListProductsRequest) (*frontierv1beta1.ListProductsResponse, error) {
	logger := grpczap.Extract(ctx)

	var products []*frontierv1beta1.Product
	productsList, err := h.productService.List(ctx, product.Filter{})
	if err != nil {
		logger.Error(err.Error())
		return nil, grpcInternalServerError
	}
	for _, v := range productsList {
		productPB, err := transformProductToPB(v)
		if err != nil {
			logger.Error(err.Error())
			return nil, grpcInternalServerError
		}
		products = append(products, productPB)
	}

	return &frontierv1beta1.ListProductsResponse{
		Products: products,
	}, nil
}

func (h Handler) GetProduct(ctx context.Context, request *frontierv1beta1.GetProductRequest) (*frontierv1beta1.GetProductResponse, error) {
	logger := grpczap.Extract(ctx)

	product, err := h.productService.GetByID(ctx, request.GetId())
	if err != nil {
		logger.Error(err.Error())
		return nil, grpcInternalServerError
	}

	productPB, err := transformProductToPB(product)
	if err != nil {
		logger.Error(err.Error())
		return nil, grpcInternalServerError
	}

	return &frontierv1beta1.GetProductResponse{
		Product: productPB,
	}, nil
}
