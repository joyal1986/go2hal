package chef

import (
	"context"
	"github.com/go-kit/kit/log"
	"time"
)

type loggingService struct {
	logger log.Logger
	Service
}

func NewLoggingService(logger log.Logger, s Service) Service {
	return &loggingService{logger, s}
}

func (s *loggingService) sendDeliveryAlert(ctx context.Context, message string) error {
	defer func(begin time.Time) {
		s.logger.Log(
			"method", "sendDeliveryAlert",
			"message", message,
			"took", time.Since(begin),
		)
	}(time.Now())

	return s.Service.sendDeliveryAlert(ctx, message)
}
func (s *loggingService) FindNodesFromFriendlyNames(recipe, environment string) ([]Node, error) {
	defer func(begin time.Time) {
		s.logger.Log(
			"method", "findNodesFromFriendlyNames",
			"recipe", recipe,
			"environment", environment,
			"took", time.Since(begin),
		)
	}(time.Now())
	return s.Service.FindNodesFromFriendlyNames(recipe, environment)

}
func (s *loggingService) SendKeyboardRecipe(ctx context.Context, message string) error {
	defer func(begin time.Time) {
		s.logger.Log(
			"method", "SendKeyboardRecipe",
			"message", message,
			"took", time.Since(begin),
		)
	}(time.Now())
	return s.Service.SendKeyboardRecipe(ctx, message)

}
func (s *loggingService) SendKeyboardEnvironment(ctx context.Context, message string) error {
	defer func(begin time.Time) {
		s.logger.Log(
			"method", "SendKeyboardEnvironment",
			"message", message,
			"took", time.Since(begin),
		)
	}(time.Now())
	return s.Service.SendKeyboardEnvironment(ctx, message)

}
func (s *loggingService) SendKeyboardNodes(ctx context.Context, recipe, environment, message string) error {
	defer func(begin time.Time) {
		s.logger.Log(
			"method", " SendKeyboardNodes",
			"recipe", recipe,
			"environment", environment,
			"message", message,
			"took", time.Since(begin),
		)
	}(time.Now())
	return s.Service.SendKeyboardNodes(ctx, recipe, environment, message)

}
