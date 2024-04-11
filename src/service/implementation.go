package service

import (
	"context"
	"test/src/transport"
)

func (s *service) SaveCurrencyRates(ctx context.Context, req *transport.SaveCurrenciesRequest) (*transport.SaveCurrenciesResponse, error) {
	currencies, err := s.rpc.GetCurrencyRate(ctx, req.Date)
	if err != nil {
		return nil, err
	}

	go s.mainStore.Mssql().UpdateCurrencyByDate(ctx, currencies)

	return &transport.SaveCurrenciesResponse{
		Success: true,
		Msg:     "started async saving",
	}, nil
}

func (s *service) GetCurrencyRateByDateAndCode(ctx context.Context, req *transport.GetCurrencyByDateAndCodeRequest) (*transport.GetCurrencyByDateAndCodeResponse, error) {
	crs, err := s.mainStore.Mssql().GetCurrenciesByDateAndCode(ctx, req.Date, req.Code)
	if err != nil {
		return nil, err
	}

	return &transport.GetCurrencyByDateAndCodeResponse{
		Currencies: crs,
	}, nil
}
