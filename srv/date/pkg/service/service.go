package service

import (
	"context"
	"time"
)

// DateService describes the service.
type DateService interface {
	ReportDate(ctx context.Context) time.Time
	// Add your methods here
	// e.x: Foo(ctx context.Context,s string)(rs string, err error)
}

type basicDateService struct{}

func (b *basicDateService) ReportDate(ctx context.Context) (t0 time.Time) {
	date := time.Now()
	return date
}

// NewBasicDateService returns a naive, stateless implementation of DateService.
func NewBasicDateService() DateService {
	return &basicDateService{}
}

// New returns a DateService with all of the expected middleware wired in.
func New(middleware []Middleware) DateService {
	var svc DateService = NewBasicDateService()
	for _, m := range middleware {
		svc = m(svc)
	}
	return svc
}
