package http

import (
	kithttp "github.com/go-kit/kit/transport/http"
	"github.com/gorilla/mux"
	"test/src/middleware"
)

func initializeRoutes(endpoints *middleware.Endpoints, options []kithttp.ServerOption) *mux.Router {
	saveCurrency := kithttp.NewServer(
		endpoints.SaveCurrencies,
		saveCurrencyRequestDecoder,
		EncodeResponse,
		options...,
	)

	getCurrencyRateByDateAndCode := kithttp.NewServer(
		endpoints.GetCurrencyByDateAndCode,
		getCurrencyRateByDateAndCodeRequestDecoder,
		EncodeResponse,
		options...,
	)

	r := mux.NewRouter()

	// swagger:route GET /currency/{date} Test SaveCurrencyRequest
	// Save currency
	// responses:
	//   200: SaveCurrencyResponse
	r.Path("/currency/{date}").
		Methods("GET").
		Handler(saveCurrency)

	// swagger:route GET /currency/{date}/{code} Test GetCurrencyRateByDateAndCodeRequest
	// Get currency
	// responses:
	//   200: GetCurrencyRateByDateAndCodeResponse
	r.Path("/currency/{date}/{code}").
		Methods("GET").
		Handler(getCurrencyRateByDateAndCode)

	return r
}
