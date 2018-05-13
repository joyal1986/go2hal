package chef

import (
	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"
	"golang.org/x/net/context"
	"os"

	"github.com/go-kit/kit/circuitbreaker"
	"github.com/go-kit/kit/endpoint"
	kitprometheus "github.com/go-kit/kit/metrics/prometheus"
	"github.com/go-kit/kit/ratelimit"
	"github.com/go-kit/kit/transport/http"
	stdprometheus "github.com/prometheus/client_golang/prometheus"
	"github.com/sony/gobreaker"
	"github.com/weAutomateEverything/go2hal/alert"
	"github.com/weAutomateEverything/go2hal/gokit"
	"golang.org/x/time/rate"
	"time"
)

type chefProxy struct {
	findNodesEndpoint               endpoint.Endpoint
	sendDeliveryEndpoint            endpoint.Endpoint
	sendKeyboardRecipeEndpoint      endpoint.Endpoint
	sendKeyboardEnvironmentEndpoint endpoint.Endpoint
	sendKeyboardNodesEndpoint       endpoint.Endpoint
}

// NewChefProxy will create a HTTP Rest client to easily invoke the chef service offered by HAL. The HAL Service
// endpoint needs to be set in a Environment Variable named HAL_ENDPOINT
func NewChefProxy() Service {
	if getHalUrl() == "" {
		panic("No HAL Endpoint set. Please set the environment variable HAL_ENDPOINT with the http address of the chef service")
	}
	logger := log.NewLogfmtLogger(log.NewSyncWriter(os.Stderr))
	logger = level.NewFilter(logger, level.AllowAll())
	logger = log.With(logger, "ts", log.DefaultTimestamp)

	return newProxy("", logger)

}

// NewKubernetesChefProxy creates a HTTP Rest client to easily call the Chef service by using the Service discovery
// mechaisms build into kubernetes. The funciton does assume that the deployment was done using the templates provided
// within the kuberets folder of the project, and that a service names hal was created. If you application resides in the
// same namespace as HAL, then the namespace can be left as a empty string, else provide the namespace that contains
// the hal deployment.
func NewKubernetesChefProxy(namespace string) Service {
	fieldKeys := []string{"method"}

	logger := log.NewLogfmtLogger(log.NewSyncWriter(os.Stderr))
	logger = level.NewFilter(logger, level.AllowAll())
	logger = log.With(logger, "ts", log.DefaultTimestamp)

	service := newProxy(namespace, logger)
	service = NewLoggingService(log.With(logger, "component", "chef_proxy"), service)
	service = NewInstrumentService(kitprometheus.NewCounterFrom(stdprometheus.CounterOpts{
		Namespace: "proxy",
		Subsystem: "callout_service",
		Name:      "request_count",
		Help:      "Number of requests received.",
	}, fieldKeys),
		kitprometheus.NewCounterFrom(stdprometheus.CounterOpts{
			Namespace: "proxy",
			Subsystem: "callout_service",
			Name:      "error_count",
			Help:      "Number of errors.",
		}, fieldKeys),
		kitprometheus.NewSummaryFrom(stdprometheus.SummaryOpts{
			Namespace: "proxy",
			Subsystem: "callout_service",
			Name:      "request_latency_microseconds",
			Help:      "Total duration of requests in microseconds.",
		}, fieldKeys), service)

	return service
}

func newProxy(namespace string, logger log.Logger) Service {
	findNodes := makeFindNodesHttpProxy(namespace, logger)
	findNodes = circuitbreaker.Gobreaker(gobreaker.NewCircuitBreaker(gobreaker.Settings{}))(findNodes)
	findNodes = ratelimit.NewErroringLimiter(rate.NewLimiter(rate.Every(time.Second), 10))(findNodes)

	sendDelivery := makeDeliveryHttpProxy(namespace, logger)
	sendDelivery = circuitbreaker.Gobreaker(gobreaker.NewCircuitBreaker(gobreaker.Settings{}))(sendDelivery)
	sendDelivery = ratelimit.NewErroringLimiter(rate.NewLimiter(rate.Every(time.Second), 10))(sendDelivery)

	sendKeyboardRecipe := makeSendKeyboardRecipeHttpProxy(namespace, logger)
	sendKeyboardRecipe = circuitbreaker.Gobreaker(gobreaker.NewCircuitBreaker(gobreaker.Settings{}))(sendKeyboardRecipe)
	sendKeyboardRecipe = ratelimit.NewErroringLimiter(rate.NewLimiter(rate.Every(time.Second), 10))(sendKeyboardRecipe)

	sendKeyboardEnvironment := makeSendKeyboardEnvironmentHttpProxy(namespace, logger)
	sendKeyboardEnvironment = circuitbreaker.Gobreaker(gobreaker.NewCircuitBreaker(gobreaker.Settings{}))(sendKeyboardEnvironment)
	sendKeyboardEnvironment = ratelimit.NewErroringLimiter(rate.NewLimiter(rate.Every(time.Second), 10))(sendKeyboardEnvironment)

	sendKeyboardNode := makeSendKeyboardNodeHttpProxy(namespace, logger)
	sendKeyboardNode = circuitbreaker.Gobreaker(gobreaker.NewCircuitBreaker(gobreaker.Settings{}))(sendKeyboardNode)
	sendKeyboardNode = ratelimit.NewErroringLimiter(rate.NewLimiter(rate.Every(time.Second), 10))(sendKeyboardNode)

	return &chefProxy{findNodesEndpoint: findNodes, sendDeliveryEndpoint: sendDelivery, sendKeyboardRecipeEndpoint: sendKeyboardRecipe, sendKeyboardEnvironmentEndpoint: sendKeyboardEnvironment, sendKeyboardNodesEndpoint: sendKeyboardNode}
}

func (s *chefProxy) FindNodesFromFriendlyNames(recipe, environment string) ([]Node, error) {
	ctx := context.Background()
	_, err := s.findNodesEndpoint(ctx, &FindNodesRequest{Recipe: recipe, Environment: environment})
	return nil, err
}
func (s *chefProxy) sendDeliveryAlert(ctx context.Context, message string) error {
	_, err := s.sendDeliveryEndpoint(ctx, message)
	return err
}

func (s *chefProxy) SendKeyboardRecipe(ctx context.Context, message string) error {
	_, err := s.sendKeyboardRecipeEndpoint(ctx, SendKeyboardRequest{Message: message})
	return err
}
func (s *chefProxy) SendKeyboardEnvironment(ctx context.Context, message string) error {
	_, err := s.sendKeyboardEnvironmentEndpoint(ctx, SendKeyboardRequest{Message: message})
	return err
}
func (s *chefProxy) SendKeyboardNodes(ctx context.Context, recipe, environment, message string) error {
	_, err := s.sendKeyboardNodesEndpoint(ctx, SendKeyboardNodeRequest{Recipe: recipe, Environment: environment, Message: message})
	return err
}
func getHalUrl() string {
	return os.Getenv("HAL_ENDPOINT")
}

func makeFindNodesHttpProxy(namespace string, logger log.Logger) endpoint.Endpoint {
	return http.NewClient(
		"POST",
		alert.GetURL(namespace, "chef/nodes"),
		gokit.EncodeJsonRequest,
		gokit.DecodeResponse,
		gokit.GetClientOpts(logger)...,
	).Endpoint()
}
func makeSendKeyboardRecipeHttpProxy(namespace string, logger log.Logger) endpoint.Endpoint {
	return http.NewClient(
		"POST",
		alert.GetURL(namespace, "chef/keyboard/recipes"),
		gokit.EncodeJsonRequest,
		gokit.DecodeResponse,
		gokit.GetClientOpts(logger)...,
	).Endpoint()
}
func makeSendKeyboardEnvironmentHttpProxy(namespace string, logger log.Logger) endpoint.Endpoint {
	return http.NewClient(
		"POST",
		alert.GetURL(namespace, "chef/keyboard/environments"),
		gokit.EncodeJsonRequest,
		gokit.DecodeResponse,
		gokit.GetClientOpts(logger)...,
	).Endpoint()
}
func makeSendKeyboardNodeHttpProxy(namespace string, logger log.Logger) endpoint.Endpoint {
	return http.NewClient(
		"POST",
		alert.GetURL(namespace, "chef/keyboard/nodes"),
		gokit.EncodeJsonRequest,
		gokit.DecodeResponse,
		gokit.GetClientOpts(logger)...,
	).Endpoint()
}
func makeDeliveryHttpProxy(namespace string, logger log.Logger) endpoint.Endpoint {
	return http.NewClient(
		"POST",
		alert.GetURL(namespace, "/delivery"),
		gokit.EncodeRequest,
		gokit.DecodeResponse,
		gokit.GetClientOpts(logger)...,
	).Endpoint()
}
