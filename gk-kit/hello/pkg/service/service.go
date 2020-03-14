package service

import "context"

// HelloService describes the service.
type HelloService interface {
	Foo(ctx context.Context, s string) (rs string, err error)
}

type basicHelloService struct{}

func (b *basicHelloService) Foo(ctx context.Context, s string) (rs string, err error) {
	// TODO implement the business logic of Foo
	return rs, err
}

// NewBasicHelloService returns a naive, stateless implementation of HelloService.
func NewBasicHelloService() HelloService {
	return &basicHelloService{}
}

// New returns a HelloService with all of the expected middleware wired in.
func New(middleware []Middleware) HelloService {
	var svc HelloService = NewBasicHelloService()
	for _, m := range middleware {
		svc = m(svc)
	}
	return svc
}
