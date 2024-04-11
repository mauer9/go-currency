package middleware

import (
	"context"
	"github.com/go-kit/log"
	"test/src/service"
	"test/src/transport"
	"time"
)

// Logging middleware entity
type loggingMiddleware struct {
	next   service.MainService
	logger log.Logger
}

// Logging middleware private method
func (lm *loggingMiddleware) logging(begin time.Time, method string, err error) {
	_ = lm.logger.Log("method", method, "took", time.Since(begin), "err", err)
}

// Logging middleware constructor
func NewLoggingMiddleware(logger log.Logger) Middleware {
	return func(next service.MainService) service.MainService {
		return &loggingMiddleware{
			next:   next,
			logger: logger,
		}
	}
}

func (lm *loggingMiddleware) SaveCurrencyRates(ctx context.Context, req *transport.SaveCurrenciesRequest) (_ *transport.SaveCurrenciesResponse, err error) {
	defer lm.logging(time.Now(), "saveCurrencyRates", err)
	return lm.next.SaveCurrencyRates(ctx, req)
}

func (lm *loggingMiddleware) GetCurrencyRateByDateAndCode(ctx context.Context, req *transport.GetCurrencyByDateAndCodeRequest) (_ *transport.GetCurrencyByDateAndCodeResponse, err error) {
	defer lm.logging(time.Now(), "getCurrencyRateByDateAndCode", err)
	return lm.next.GetCurrencyRateByDateAndCode(ctx, req)
}
