package client

import (
	"context"
	"git.tmuyu.com.cn/demo/srv/login/pkg/grpc/pb"
	"google.golang.org/grpc"
	"log"
	"testing"
)

func TestGrpc(t *testing.T) {
	address := "localhost:8082"
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewLoginClient(conn)

	// Contact the server and print out its response.
	token := "admin||1643445402"

	r, err := c.RefreshToken(context.Background(), &pb.RefreshTokenRequest{OldToken: token})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}

	log.Printf("####### get server Greeting response: %s", r.GetToken())
}
