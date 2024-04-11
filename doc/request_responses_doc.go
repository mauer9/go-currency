package doc

import t "test/src/transport"

type DefaultHeader struct {
	// in:header
	ContentType string `json:"Content-Type"`
}

// swagger:parameters SaveCurrencyRequest
type SaveCurrencyRequest struct {
	DefaultHeader

	t.SaveCurrenciesRequest
}

// swagger:response SaveCurrencyResponse
type SaveCurrencyResponse struct {
	// in:body
	Body t.SaveCurrenciesResponse
}

// swagger:parameters GetCurrencyRateByDateAndCodeRequest
type GetCurrencyRateByDateAndCodeRequest struct {
	DefaultHeader

	t.GetCurrencyByDateAndCodeRequest
}

// swagger:response GetCurrencyRateByDateAndCodeResponse
type GetCurrencyRateByDateAndCodeResponse struct {
	// in:body
	Body t.GetCurrencyByDateAndCodeResponse
}
