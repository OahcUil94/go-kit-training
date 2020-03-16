package endpoints

import (
	"context"

	"gk-kitgen/service/service"
	"github.com/go-kit/kit/endpoint"
)

type PostProfileRequest struct {
	P service.Profile
}
type PostProfileResponse struct {
	Err error
}

func MakePostProfileEndpoint(s service.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(PostProfileRequest)
		err := s.PostProfile(ctx, req.P)
		return PostProfileResponse{Err: err}, nil
	}
}

type Endpoints struct {
	PostProfile endpoint.Endpoint
}
