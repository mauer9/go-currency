package http

import (
	"context"
	"encoding/json"
	"net/http"
	"test/src/utils/errors"
)

// Errors interface
type errorer interface {
	error() error
}

func EncodeResponse(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	err, ok := response.(errorer)
	if ok && err.error() != nil {
		EncodeErrorResponse(ctx, err.error(), w)
		return nil
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	return json.NewEncoder(w).Encode(response)
}

// EncodeErrorResponse
// errors encoder
func EncodeErrorResponse(ctx context.Context, err error, w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(codeFrom(ctx, err))
	_ = json.NewEncoder(w).Encode(err)
}

// Determine error code
func codeFrom(ctx context.Context, err error) int {
	switch err {
	case errors.AccessDenied:
		return http.StatusUnauthorized
	case errors.InvalidCharacter:
		return http.StatusBadRequest
	case errors.IncorrectRequest:
		return http.StatusBadRequest
	case errors.NotFound:
		return http.StatusNotFound
	default:
		return http.StatusInternalServerError
	}
}
