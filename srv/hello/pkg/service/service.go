package service

import "context"

// HelloService describes the service.
type HelloService interface {
	SayHello(ctx context.Context, s string) (rs string, err error)
}
type Hello struct {
}

func (he Hello) SayHello(ctx context.Context, s string) (rs string, err error) {
	rs = s
	return
}

type basicHelloService struct{}

func (b *basicHelloService) SayHello(ctx context.Context, s string) (rs string, err error) {
	rs = s
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
