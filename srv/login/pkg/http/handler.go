package http

import (
	"context"
	"encoding/json"
	"errors"
	endpoint "git.tmuyu.com.cn/demo/srv/login/pkg/endpoint"
	http1 "github.com/go-kit/kit/transport/http"
	"net/http"
)

// makeLoginHandler creates the handler logic
func makeLoginHandler(m *http.ServeMux, endpoints endpoint.Endpoints, options []http1.ServerOption) {
	m.Handle("/login", http1.NewServer(endpoints.LoginEndpoint, decodeLoginRequest, encodeLoginResponse, options...))
}

// decodeLoginRequest is a transport/http.DecodeRequestFunc that decodes a
// JSON-encoded request from the HTTP request body.
func decodeLoginRequest(_ context.Context, r *http.Request) (interface{}, error) {
	req := endpoint.LoginRequest{}
	err := json.NewDecoder(r.Body).Decode(&req)
	return req, err
}

// encodeLoginResponse is a transport/http.EncodeResponseFunc that encodes
// the response as JSON to the response writer
func encodeLoginResponse(ctx context.Context, w http.ResponseWriter, response interface{}) (err error) {
	if f, ok := response.(endpoint.Failure); ok && f.Failed() != nil {
		ErrorEncoder(ctx, f.Failed(), w)
		return nil
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	err = json.NewEncoder(w).Encode(response)
	return
}

// makeRefreshTokenHandler creates the handler logic
func makeRefreshTokenHandler(m *http.ServeMux, endpoints endpoint.Endpoints, options []http1.ServerOption) {
	m.Handle("/refresh-token", http1.NewServer(endpoints.RefreshTokenEndpoint, decodeRefreshTokenRequest, encodeRefreshTokenResponse, options...))
}

// decodeRefreshTokenRequest is a transport/http.DecodeRequestFunc that decodes a
// JSON-encoded request from the HTTP request body.
func decodeRefreshTokenRequest(_ context.Context, r *http.Request) (interface{}, error) {
	req := endpoint.RefreshTokenRequest{}
	err := json.NewDecoder(r.Body).Decode(&req)
	return req, err
}

// encodeRefreshTokenResponse is a transport/http.EncodeResponseFunc that encodes
// the response as JSON to the response writer
func encodeRefreshTokenResponse(ctx context.Context, w http.ResponseWriter, response interface{}) (err error) {
	if f, ok := response.(endpoint.Failure); ok && f.Failed() != nil {
		ErrorEncoder(ctx, f.Failed(), w)
		return nil
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	err = json.NewEncoder(w).Encode(response)
	return
}
func ErrorEncoder(_ context.Context, err error, w http.ResponseWriter) {
	w.WriteHeader(err2code(err))
	json.NewEncoder(w).Encode(errorWrapper{Error: err.Error()})
}
func ErrorDecoder(r *http.Response) error {
	var w errorWrapper
	if err := json.NewDecoder(r.Body).Decode(&w); err != nil {
		return err
	}
	return errors.New(w.Error)
}

// This is used to set the http status, see an example here :
// https://github.com/go-kit/kit/blob/master/examples/addsvc/pkg/addtransport/http.go#L133
func err2code(err error) int {
	return http.StatusInternalServerError
}

type errorWrapper struct {
	Error string `json:"error"`
}
