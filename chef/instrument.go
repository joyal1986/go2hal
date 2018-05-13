package chef

import (
	"context"
	"github.com/go-kit/kit/metrics"
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

func (s *instrumentingService) sendDeliveryAlert(ctx context.Context, message string) error {
	defer func(begin time.Time) {
		s.requestCount.With("method", "sendDeliveryAlert").Add(1)
		s.requestLatency.With("method", "sendDeliveryAlert").Observe(time.Since(begin).Seconds())
	}(time.Now())
	return s.Service.sendDeliveryAlert(ctx, message)
}
func (s *instrumentingService) FindNodesFromFriendlyNames(recipe, environment string) ([]Node, error) {
	defer func(begin time.Time) {
		s.requestCount.With("method", "findNodesFromFriendlyNames").Add(1)
		s.requestLatency.With("method", "findNodesFromFriendlyNames").Observe(time.Since(begin).Seconds())
	}(time.Now())
	return s.Service.FindNodesFromFriendlyNames(recipe, environment)
}
func (s *instrumentingService) SendKeyboardRecipe(ctx context.Context, message string) error {
	defer func(begin time.Time) {
		s.requestCount.With("method", "send keyboard recipe").Add(1)
		s.requestLatency.With("method", "send keyboard recipe").Observe(time.Since(begin).Seconds())
	}(time.Now())
	return s.Service.SendKeyboardRecipe(ctx, message)
}
func (s *instrumentingService) SendKeyboardEnvironment(ctx context.Context, message string) error {
	defer func(begin time.Time) {
		s.requestCount.With("method", "send keyboard environment").Add(1)
		s.requestLatency.With("method", "send keyboard environment").Observe(time.Since(begin).Seconds())
	}(time.Now())
	return s.Service.SendKeyboardEnvironment(ctx, message)
}

func (s *instrumentingService) SendKeyboardNodes(ctx context.Context, recipe, environment, message string) error {
	defer func(begin time.Time) {
		s.requestCount.With("method", "send keyboard nodes").Add(1)
		s.requestLatency.With("method", "send keyboard nodes").Observe(time.Since(begin).Seconds())
	}(time.Now())
	return s.Service.SendKeyboardNodes(ctx, recipe, environment, message)
}
