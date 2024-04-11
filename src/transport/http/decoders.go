package http

import (
	"context"
	"github.com/gorilla/mux"
	"net/http"
	"test/src/transport"
	"test/src/utils/errors"
)

func saveCurrencyRequestDecoder(_ context.Context, r *http.Request) (interface{}, error) {
	var request transport.SaveCurrenciesRequest

	vars := mux.Vars(r)
	date, ok := vars["date"]
	if !ok {
		return nil, errors.NotFound
	}

	request.Date = date

	err := request.Validate()
	if err != nil {
		return nil, err
	}

	return request, nil
}

func getCurrencyRateByDateAndCodeRequestDecoder(_ context.Context, r *http.Request) (interface{}, error) {
	var request transport.GetCurrencyByDateAndCodeRequest
	vars := mux.Vars(r)
	date, ok := vars["date"]
	if !ok {
		return nil, errors.NotFound
	}

	code, ok := vars["code"]
	if !ok {
		return nil, errors.NotFound
	}

	request.Date = date
	request.Code = code

	err := request.Validate()
	if err != nil {
		return nil, err
	}

	return request, nil
}
