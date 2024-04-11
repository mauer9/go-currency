package service

import (
	"context"
	"test/src/transport"
)

type MainService interface {
	SaveCurrencyRates(ctx context.Context, req *transport.SaveCurrenciesRequest) (*transport.SaveCurrenciesResponse, error)
	GetCurrencyRateByDateAndCode(ctx context.Context, req *transport.GetCurrencyByDateAndCodeRequest) (*transport.GetCurrencyByDateAndCodeResponse, error)
}
