package transport

import "test/src/domain"

type SaveCurrenciesRequest struct {
	Date string `json:"-"`
}

type SaveCurrenciesResponse struct {
	Success bool   `json:"success"`
	Msg     string `json:"msg"`
}

type GetCurrencyByDateAndCodeRequest struct {
	Date string `json:"-"`
	Code string `json:"-"`
}

type GetCurrencyByDateAndCodeResponse struct {
	Currencies []domain.Currency `json:"currencies"`
}
