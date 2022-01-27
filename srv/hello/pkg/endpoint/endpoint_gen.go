// THIS FILE IS AUTO GENERATED BY GK-CLI DO NOT EDIT!!
package endpoint

import (
	service "git.tmuyu.com.cn/demo/srv/hello/pkg/service"
	endpoint "github.com/go-kit/kit/endpoint"
)

// Endpoints collects all of the endpoints that compose a profile service. It's
// meant to be used as a helper struct, to collect all of the endpoints into a
// single parameter.
type Endpoints struct {
	SayHelloEndpoint endpoint.Endpoint
}

// New returns a Endpoints struct that wraps the provided service, and wires in all of the
// expected endpoint middlewares
func New(s service.HelloService, mdw map[string][]endpoint.Middleware) Endpoints {
	eps := Endpoints{SayHelloEndpoint: MakeSayHelloEndpoint(s)}
	for _, m := range mdw["SayHello"] {
		eps.SayHelloEndpoint = m(eps.SayHelloEndpoint)
	}
	return eps
}
