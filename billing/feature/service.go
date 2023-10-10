package feature

import (
	"context"
	"fmt"

	"github.com/mcuadros/go-defaults"

	"github.com/google/uuid"
	"github.com/raystack/frontier/pkg/utils"
	"github.com/stripe/stripe-go/v75"
	"github.com/stripe/stripe-go/v75/client"
)

type Repository interface {
	GetByID(ctx context.Context, id string) (Feature, error)
	GetByName(ctx context.Context, name string) (Feature, error)
	Create(ctx context.Context, feature Feature) (Feature, error)
	UpdateByName(ctx context.Context, feature Feature) (Feature, error)
	List(ctx context.Context, flt Filter) ([]Feature, error)
}

type PriceRepository interface {
	GetByID(ctx context.Context, id string) (Price, error)
	GetByName(ctx context.Context, name string) (Price, error)
	Create(ctx context.Context, price Price) (Price, error)
	UpdateByID(ctx context.Context, price Price) (Price, error)
	List(ctx context.Context, flt Filter) ([]Price, error)
}

type Service struct {
	stripeClient    *client.API
	repository      Repository
	priceRepository PriceRepository
}

func NewService(stripeClient *client.API, repository Repository,
	priceRepository PriceRepository) *Service {
	return &Service{
		stripeClient:    stripeClient,
		priceRepository: priceRepository,
		repository:      repository,
	}
}

func (s *Service) Create(ctx context.Context, feature Feature) (Feature, error) {
	// create a product in stripe for each feature in plan
	if feature.ID == "" {
		feature.ID = uuid.New().String()
	}
	defaults.SetDefaults(&feature)
	_, err := s.stripeClient.Products.New(&stripe.ProductParams{
		Params: stripe.Params{
			Context: ctx,
		},
		ID:          &feature.ID,
		Name:        &feature.Title,
		Description: &feature.Description,
		Metadata: map[string]string{
			"name": feature.Name,
		},
	})
	if err != nil {
		return Feature{}, err
	}

	return s.repository.Create(ctx, feature)
}

func (s *Service) GetByID(ctx context.Context, id string) (Feature, error) {
	var fetchedFeature Feature
	var err error
	if utils.IsValidUUID(id) {
		fetchedFeature, err = s.repository.GetByID(ctx, id)
		if err != nil {
			return Feature{}, err
		}
	}
	fetchedFeature, err = s.repository.GetByName(ctx, id)
	if err != nil {
		return Feature{}, err
	}

	if fetchedFeature.Prices, err = s.priceRepository.List(ctx, Filter{
		FeatureID: fetchedFeature.ID,
	}); err != nil {
		return Feature{}, fmt.Errorf("failed to fetch prices for feature %s: %w", fetchedFeature.ID, err)
	}
	return fetchedFeature, nil
}

// Update updates a feature, but it doesn't update all fields
// ideally we should keep it immutable and create a new feature
func (s *Service) Update(ctx context.Context, feature Feature) (Feature, error) {
	// update product in stripe
	_, err := s.stripeClient.Products.Update(feature.ID, &stripe.ProductParams{
		Params: stripe.Params{
			Context: ctx,
		},
		Name:        &feature.Title,
		Description: &feature.Description,
		Metadata: map[string]string{
			"name":       feature.Name,
			"managed_by": "frontier",
		},
	})
	if err != nil {
		return Feature{}, err
	}
	return s.repository.UpdateByName(ctx, feature)
}

func (s *Service) CreatePrice(ctx context.Context, price Price) (Price, error) {
	// set defaults
	if price.Type == "" {
		price.Type = PriceTypeOnetime
	}
	if price.BillingScheme == "" {
		price.BillingScheme = BillingSchemeFlat
	}
	if price.Currency == "" {
		price.Currency = "usd"
	}
	if price.UsageType == "" {
		price.UsageType = PriceUsageTypeLicensed
	}

	providerParams := &stripe.PriceParams{
		Params: stripe.Params{
			Context: ctx,
		},
		Product:       &price.FeatureID,
		Nickname:      &price.Name,
		BillingScheme: stripe.String(price.BillingScheme.ToStripe()),
		Currency:      &price.Currency,
		UnitAmount:    &price.Amount,
		Metadata: map[string]string{
			"title":      price.Title,
			"managed_by": "frontier",
		},
	}
	if price.RecurringInterval != "" {
		providerParams.Recurring = &stripe.PriceRecurringParams{
			Interval:  stripe.String(price.RecurringInterval),
			UsageType: stripe.String(price.UsageType.ToStripe()),
		}
		if price.UsageType == PriceUsageTypeMetered {
			providerParams.Recurring.AggregateUsage = stripe.String(price.MeteredAggregate)
		}
	}
	stripePrice, err := s.stripeClient.Prices.New(providerParams)
	if err != nil {
		return Price{}, err
	}

	price.ProviderID = stripePrice.ID
	return s.priceRepository.Create(ctx, price)
}

func (s *Service) GetPriceByID(ctx context.Context, id string) (Price, error) {
	if utils.IsValidUUID(id) {
		return s.priceRepository.GetByID(ctx, id)
	}
	return s.priceRepository.GetByName(ctx, id)
}

// UpdatePrice updates a price, but it doesn't update all fields
// ideally we should keep it immutable and create a new price
func (s *Service) UpdatePrice(ctx context.Context, price Price) (Price, error) {
	_, err := s.stripeClient.Prices.Update(price.ProviderID, &stripe.PriceParams{
		Params: stripe.Params{
			Context: ctx,
		},
		Nickname: &price.Name,
		Metadata: map[string]string{
			"title": price.Title,
		},
	})
	if err != nil {
		return Price{}, err
	}
	return s.priceRepository.UpdateByID(ctx, price)
}

func (s *Service) List(ctx context.Context, flt Filter) ([]Feature, error) {
	listedFeatures, err := s.repository.List(ctx, flt)
	if err != nil {
		return nil, err
	}

	// enrich with prices
	for i, listedFeature := range listedFeatures {
		prices, err := s.priceRepository.List(ctx, Filter{
			FeatureID: listedFeature.ID,
		})
		if err != nil {
			return nil, err
		}
		listedFeatures[i].Prices = prices
	}
	return listedFeatures, nil
}
