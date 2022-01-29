package endpoint

import (
	"context"

	service "git.tmuyu.com.cn/demo/srv/login/pkg/service"
	endpoint "github.com/go-kit/kit/endpoint"
)

// LoginRequest collects the request parameters for the Login method.
type LoginRequest struct {
	Name string `json:"name"`
	Pwd  string `json:"pwd"`
}

// LoginResponse collects the response parameters for the Login method.
type LoginResponse struct {
	Token string `json:"token"`
	Err   error  `json:"err"`
}

// MakeLoginEndpoint returns an endpoint that invokes Login on the service.
func MakeLoginEndpoint(s service.LoginService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(LoginRequest)
		token, err := s.Login(ctx, req.Name, req.Pwd)
		return LoginResponse{
			Err:   err,
			Token: token,
		}, nil
	}
}

// Failed implements Failer.
func (r LoginResponse) Failed() error {
	return r.Err
}

// RefreshTokenRequest collects the request parameters for the RefreshToken method.
type RefreshTokenRequest struct {
	OldToken string `json:"old_token"`
}

// RefreshTokenResponse collects the response parameters for the RefreshToken method.
type RefreshTokenResponse struct {
	Token string `json:"token"`
}

// MakeRefreshTokenEndpoint returns an endpoint that invokes RefreshToken on the service.
func MakeRefreshTokenEndpoint(s service.LoginService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(RefreshTokenRequest)
		token := s.RefreshToken(ctx, req.OldToken)
		return RefreshTokenResponse{Token: token}, nil
	}
}

// Failure is an interface that should be implemented by response types.
// Response encoders can check if responses are Failer, and if so they've
// failed, and if so encode them using a separate write path based on the error.
type Failure interface {
	Failed() error
}

// Login implements Service. Primarily useful in a client.
func (e Endpoints) Login(ctx context.Context, name string, pwd string) (token string, err error) {
	request := LoginRequest{
		Name: name,
		Pwd:  pwd,
	}
	response, err := e.LoginEndpoint(ctx, request)
	if err != nil {
		return
	}
	return response.(LoginResponse).Token, response.(LoginResponse).Err
}

// RefreshToken implements Service. Primarily useful in a client.
func (e Endpoints) RefreshToken(ctx context.Context, oldToken string) (token string) {
	request := RefreshTokenRequest{OldToken: oldToken}
	response, err := e.RefreshTokenEndpoint(ctx, request)
	if err != nil {
		return
	}
	return response.(RefreshTokenResponse).Token
}
