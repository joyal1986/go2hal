package alert

import (
	"context"
	"github.com/go-kit/kit/endpoint"
	"github.com/pkg/errors"
)



//alert group rest service response
type AlertGroupResponse struct {
	Groupid int64 `json:"groupid"`
}

func makeAlertEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(string)
		return nil, s.SendAlert(ctx, req)
	}
}

func makeImageAlertEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.([]byte)
		return nil, s.SendImageToAlertGroup(ctx, req)
	}
}

func makeHeartbeatMessageEncpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(string)
		return nil, s.SendHeartbeatGroupAlert(ctx, req)
	}
}

func makeImageHeartbeatEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.([]byte)
		return nil, s.SendImageToHeartbeatGroup(ctx, req)
	}
}

func makeBusinessAlertEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(string)
		return nil, s.SendNonTechnicalAlert(ctx, req)
	}
}

func makeAlertErrorHandler(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(string)
		return nil, s.SendError(ctx, errors.New(req))

	}

}

func makeAlertGroupHandler(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		groupid, err := s.AlertGroup(ctx)
		if err != nil {
			return nil, err
		}
		return AlertGroupResponse{Groupid: groupid}, nil
	}
}
