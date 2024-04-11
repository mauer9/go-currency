package http

import (
	kittransport "github.com/go-kit/kit/transport"
	kithttp "github.com/go-kit/kit/transport/http"
	"github.com/go-kit/log"
	"net/http"
	"test/src/middleware"
)

// HTTP service constructor
func NewHTTPService(svcEndpoints *middleware.Endpoints, options []kithttp.ServerOption, logger log.Logger) http.Handler {
	errorEncoder := kithttp.ServerErrorEncoder(
		EncodeErrorResponse,
	)

	errorLogger := kithttp.ServerErrorHandler(
		kittransport.NewLogErrorHandler(logger),
	)

	options = append(options, errorEncoder, errorLogger)

	return initializeRoutes(svcEndpoints, options)
}
