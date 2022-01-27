package http

import (
	"context"
	"encoding/json"
	endpoint "git.tmuyu.com.cn/demo/srv/date/pkg/endpoint"
	http1 "github.com/go-kit/kit/transport/http"
	"net/http"
)

// makeReportDateHandler creates the handler logic
func makeReportDateHandler(m *http.ServeMux, endpoints endpoint.Endpoints, options []http1.ServerOption) {
	m.Handle("/report-date", http1.NewServer(endpoints.ReportDateEndpoint, decodeReportDateRequest, encodeReportDateResponse, options...))
}

// decodeReportDateRequest is a transport/http.DecodeRequestFunc that decodes a
// JSON-encoded request from the HTTP request body.
func decodeReportDateRequest(_ context.Context, r *http.Request) (interface{}, error) {
	req := endpoint.ReportDateRequest{}
	err := json.NewDecoder(r.Body).Decode(&req)
	return req, err
}

// encodeReportDateResponse is a transport/http.EncodeResponseFunc that encodes
// the response as JSON to the response writer
func encodeReportDateResponse(ctx context.Context, w http.ResponseWriter, response interface{}) (err error) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	err = json.NewEncoder(w).Encode(response)
	return
}
