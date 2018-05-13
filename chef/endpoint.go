package chef

import (
	"context"
	"github.com/go-kit/kit/endpoint"
)

type AddChefClientRequest struct {
	Name, Key, URL string
}

//find nodes rest service request
type FindNodesRequest struct {
	Recipe      string `json:"recipe"`
	Environment string `json:"environment"`
}

//find nodes rest service response
type FindNodesResponse struct {
	Nodes []Node `json:"nodes"`
}

//send keyboard service request
type SendKeyboardRequest struct {
	Message string `json:"message"`
}

//send keyboard node service request
type SendKeyboardNodeRequest struct {
	Recipe      string `json:"recipe"`
	Environment string `json:"environment"`
	Message     string `json:"message"`
}

func makeChefDeliveryAlertEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(string)
		return nil, s.sendDeliveryAlert(ctx, req)
	}
}
func makeFindNodesHandler(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(FindNodesRequest)
		nodes, _ := s.FindNodesFromFriendlyNames(req.Recipe, req.Environment)
		return FindNodesResponse{Nodes: nodes}, nil
	}
}
func makeSendKeyboardRecipeHandler(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(SendKeyboardRequest)
		return nil, s.SendKeyboardRecipe(ctx, req.Message)
	}
}
func makeSendKeyboardEnvironmentHandler(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(SendKeyboardRequest)
		return nil, s.SendKeyboardEnvironment(ctx, req.Message)
	}
}
func makeSendKeyboardNodesHandler(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(SendKeyboardNodeRequest)
		return nil, s.SendKeyboardNodes(ctx, req.Recipe, req.Environment, req.Message)
	}
}
