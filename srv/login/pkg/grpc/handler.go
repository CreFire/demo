package grpc

import (
	"context"
	"errors"
	endpoint "git.tmuyu.com.cn/demo/srv/login/pkg/endpoint"
	pb "git.tmuyu.com.cn/demo/srv/login/pkg/grpc/pb"
	grpc "github.com/go-kit/kit/transport/grpc"
	context1 "golang.org/x/net/context"
)

// makeLoginHandler creates the handler logic
func makeLoginHandler(endpoints endpoint.Endpoints, options []grpc.ServerOption) grpc.Handler {
	return grpc.NewServer(endpoints.LoginEndpoint, decodeLoginRequest, encodeLoginResponse, options...)
}

// decodeLoginResponse is a transport/grpc.DecodeRequestFunc that converts a
// gRPC request to a user-domain Login request.
// TODO implement the decoder
func decodeLoginRequest(_ context.Context, r interface{}) (interface{}, error) {
	return nil, errors.New("'Login' Decoder is not impelemented")
}

// encodeLoginResponse is a transport/grpc.EncodeResponseFunc that converts
// a user-domain response to a gRPC reply.
// TODO implement the encoder
func encodeLoginResponse(_ context.Context, r interface{}) (interface{}, error) {
	return nil, errors.New("'Login' Encoder is not impelemented")
}
func (g *grpcServer) Login(ctx context1.Context, req *pb.LoginRequest) (*pb.LoginReply, error) {
	_, rep, err := g.login.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return rep.(*pb.LoginReply), nil
}

// makeRefreshTokenHandler creates the handler logic
func makeRefreshTokenHandler(endpoints endpoint.Endpoints, options []grpc.ServerOption) grpc.Handler {
	return grpc.NewServer(endpoints.RefreshTokenEndpoint, decodeRefreshTokenRequest, encodeRefreshTokenResponse, options...)
}

// decodeRefreshTokenResponse is a transport/grpc.DecodeRequestFunc that converts a
// gRPC request to a user-domain RefreshToken request.
// TODO implement the decoder
func decodeRefreshTokenRequest(_ context.Context, r interface{}) (interface{}, error) {
	if r != nil {
		req := endpoint.RefreshTokenRequest{OldToken: r.(*pb.RefreshTokenRequest).OldToken}

		return req,nil
	}
	return nil, errors.New("'Login' Decoder is not impelemented")
}

// encodeRefreshTokenResponse is a transport/grpc.EncodeResponseFunc that converts
// a user-domain response to a gRPC reply.
// TODO implement the encoder
func encodeRefreshTokenResponse(_ context.Context, r interface{}) (interface{}, error) {
	if r != nil {
		return &pb.RefreshTokenReply{Token: r.(endpoint.RefreshTokenResponse).Token},nil
	}
	return nil, errors.New("'Login' Encoder is not impelemented")
}
func (g *grpcServer) RefreshToken(ctx context1.Context, req *pb.RefreshTokenRequest) (*pb.RefreshTokenReply, error) {
	_, rep, err := g.refreshToken.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return rep.(*pb.RefreshTokenReply), nil
}
