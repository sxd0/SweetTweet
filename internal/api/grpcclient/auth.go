package grpcclient

import (
	"context"
	"log"
	"os"
	"time"

	authpb "github.com/sxd0/SweetTweet/proto/authpb"
	"google.golang.org/grpc"
)

var AuthClient authpb.AuthServiceClient

func InitAuthClient() {
	addr := os.Getenv("AUTH_SERVICE_ADDR")
	if addr == "" {
		addr = "auth:8081"
		log.Println("[WARN] AUTH_SERVICE_ADDR not set, fallback to", addr)
	}

	conn, err := grpc.Dial(addr, grpc.WithInsecure(), grpc.WithBlock(), grpc.WithTimeout(5*time.Second))
	if err != nil {
		log.Fatalf("failed to connect to auth service: %v", err)
	}

	AuthClient = authpb.NewAuthServiceClient(conn)
}

func Register(ctx context.Context, email, password string) (*authpb.RegisterResponse, error) {
	return AuthClient.Register(ctx, &authpb.RegisterRequest{Email: email, Password: password})
}

func Login(ctx context.Context, email, password string) (*authpb.LoginResponse, error) {
	return AuthClient.Login(ctx, &authpb.LoginRequest{Email: email, Password: password})
}

func Me(ctx context.Context, token string) (*authpb.MeResponse, error) {
	return AuthClient.Me(ctx, &authpb.MeRequest{Token: token})
}