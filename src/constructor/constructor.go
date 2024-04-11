package constructor

import (
	kitprometheus "github.com/go-kit/kit/metrics/prometheus"
	"github.com/go-kit/log"
	stdprometheus "github.com/prometheus/client_golang/prometheus"
	"test/src/middleware"
	"test/src/repository"
	"test/src/rpc"
	"test/src/service"
)

// Auth service constructor
func NewService(mainRepo repository.MainStore, rpc rpc.RPC, logger log.Logger) service.MainService {
	svc := service.NewService(rpc, logger, mainRepo)
	svc = middleware.NewLoggingMiddleware(logger)(svc)
	svc = middleware.NewInstrumentingMiddleware(
		kitprometheus.NewCounterFrom(stdprometheus.CounterOpts{
			Namespace: "api",
			Subsystem: "auth_service",
			Name:      "request_count",
			Help:      "Number of requests received.",
		}, []string{"method"}),
		kitprometheus.NewCounterFrom(stdprometheus.CounterOpts{
			Namespace: "api",
			Subsystem: "auth_service",
			Name:      "error_count",
			Help:      "Number of error requests received.",
		}, []string{"method"}),
		kitprometheus.NewSummaryFrom(stdprometheus.SummaryOpts{
			Namespace: "api",
			Subsystem: "auth_service",
			Name:      "request_latency_microseconds",
			Help:      "Total duration of requests in microseconds.",
		}, []string{"method"}),
	)(svc)

	return svc
}
