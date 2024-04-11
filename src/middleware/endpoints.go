package middleware

import (
	"context"
	"github.com/go-kit/kit/endpoint"
	"test/src/service"
	"test/src/transport"
)

// Endpoints holds all go-kit endpoints for the service
type Endpoints struct {
	SaveCurrencies           endpoint.Endpoint
	GetCurrencyByDateAndCode endpoint.Endpoint
}

// MakeEndpoints initializes all go-kit endpoints for the service
func MakeEndpoints(s service.MainService) *Endpoints {
	return &Endpoints{
		SaveCurrencies:           makeSaveCurrencies(s),
		GetCurrencyByDateAndCode: makeGetCurrencyByDateAndCode(s),
	}
}

func makeSaveCurrencies(s service.MainService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(transport.SaveCurrenciesRequest)
		return s.SaveCurrencyRates(ctx, &req)
	}
}

func makeGetCurrencyByDateAndCode(s service.MainService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(transport.GetCurrencyByDateAndCodeRequest)
		return s.GetCurrencyRateByDateAndCode(ctx, &req)
	}
}
