package middleware

import (
	"context"
	"github.com/go-kit/kit/metrics"
	"test/src/service"
	"test/src/transport"
	"time"
)

// Instrumenting middleware entity
type instrumentingMiddleware struct {
	next           service.MainService
	requestCount   metrics.Counter
	requestError   metrics.Counter
	requestLatency metrics.Histogram
}

// Instrumenting middleware private method
func (im *instrumentingMiddleware) instrumenting(begin time.Time, method string, err error) {
	im.requestCount.With("method", method).Add(1)
	if err != nil {
		im.requestError.With("method", method).Add(1)
	}
	im.requestLatency.With("method", method).Observe(time.Since(begin).Seconds())
}

// Instrumenting middleware constructor
func NewInstrumentingMiddleware(counter, counterErr metrics.Counter, latency metrics.Histogram) Middleware {
	return func(next service.MainService) service.MainService {
		return &instrumentingMiddleware{
			next:           next,
			requestCount:   counter,
			requestError:   counterErr,
			requestLatency: latency,
		}
	}
}

func (im *instrumentingMiddleware) SaveCurrencyRates(ctx context.Context, req *transport.SaveCurrenciesRequest) (_ *transport.SaveCurrenciesResponse, err error) {
	defer im.instrumenting(time.Now(), "saveCurrencyRates", err)
	return im.next.SaveCurrencyRates(ctx, req)
}

func (im *instrumentingMiddleware) GetCurrencyRateByDateAndCode(ctx context.Context, req *transport.GetCurrencyByDateAndCodeRequest) (_ *transport.GetCurrencyByDateAndCodeResponse, err error) {
	defer im.instrumenting(time.Now(), "getCurrencyByDateAndCode", err)
	return im.next.GetCurrencyRateByDateAndCode(ctx, req)
}
