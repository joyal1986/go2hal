package callout

import (
	"github.com/go-kit/kit/metrics"
	"golang.org/x/net/context"
	"time"
)

type instrumentingService struct {
	requestCount   metrics.Counter
	requestLatency metrics.Histogram
	errorCount     metrics.Counter
	Service
}

func NewInstrumentService(counter metrics.Counter, errorCount metrics.Counter, latency metrics.Histogram,
	s Service) Service {
	return &instrumentingService{
		requestCount:   counter,
		requestLatency: latency,
		errorCount:     errorCount,
		Service:        s,
	}
}

func (s instrumentingService) InvokeCallout(ctx context.Context, title, message string) (err error) {
	defer func(begin time.Time) {
		s.requestCount.With("method", "InvokeCallout").Add(1)
		s.requestLatency.With("method", "InvokeCallout").Observe(time.Since(begin).Seconds())
		if err != nil {
			s.errorCount.With("method", "InvokeCallout").Add(1)
		}
	}(time.Now())
	return s.Service.InvokeCallout(ctx, title, message)
}

func (s instrumentingService) getFirstCallName(ctx context.Context) (name string, phone string, err error) {
	defer func(begin time.Time) {
		s.requestCount.With("method", "getFirstCallName").Add(1)
		if err != nil {
			s.errorCount.With("method", "getFirstCallName").Add(1)

		}
	}(time.Now())
	return s.Service.getFirstCall(ctx)
}
