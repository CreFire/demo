package endpoint

import (
	"context"
	service "git.tmuyu.com.cn/demo/srv/date/pkg/service"
	endpoint "github.com/go-kit/kit/endpoint"
	"time"
)

// ReportDateRequest collects the request parameters for the ReportDate method.
type ReportDateRequest struct{}

// ReportDateResponse collects the response parameters for the ReportDate method.
type ReportDateResponse struct {
	T0 time.Time `json:"t0"`
}

// MakeReportDateEndpoint returns an endpoint that invokes ReportDate on the service.
func MakeReportDateEndpoint(s service.DateService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		t0 := s.ReportDate(ctx)
		return ReportDateResponse{T0: t0}, nil
	}
}

// ReportDate implements Service. Primarily useful in a client.
func (e Endpoints) ReportDate(ctx context.Context) (t0 time.Time) {
	request := ReportDateRequest{}
	response, err := e.ReportDateEndpoint(ctx, request)
	if err != nil {
		return
	}
	return response.(ReportDateResponse).T0
}
