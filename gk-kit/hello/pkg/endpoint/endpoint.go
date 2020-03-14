package endpoint

import (
	"context"
	service "github.com/OahcUil94/go-kit-training/gk-kit/hello/pkg/service"
	endpoint "github.com/go-kit/kit/endpoint"
)

// FooRequest collects the request parameters for the Foo method.
type FooRequest struct {
	S string `json:"s"`
}

// FooResponse collects the response parameters for the Foo method.
type FooResponse struct {
	Rs  string `json:"rs"`
	Err error  `json:"err"`
}

// MakeFooEndpoint returns an endpoint that invokes Foo on the service.
func MakeFooEndpoint(s service.HelloService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(FooRequest)
		rs, err := s.Foo(ctx, req.S)
		return FooResponse{
			Err: err,
			Rs:  rs,
		}, nil
	}
}

// Failed implements Failer.
func (r FooResponse) Failed() error {
	return r.Err
}

// Failure is an interface that should be implemented by response types.
// Response encoders can check if responses are Failer, and if so they've
// failed, and if so encode them using a separate write path based on the error.
type Failure interface {
	Failed() error
}

// Foo implements Service. Primarily useful in a client.
func (e Endpoints) Foo(ctx context.Context, s string) (rs string, err error) {
	request := FooRequest{S: s}
	response, err := e.FooEndpoint(ctx, request)
	if err != nil {
		return
	}
	return response.(FooResponse).Rs, response.(FooResponse).Err
}
