package grpcclient

import (
	"context"
	"log"
	"os"
	"time"

	userpb "github.com/sxd0/SweetTweet/proto/userpb"
	"google.golang.org/grpc"
)

var UserClient userpb.UserServiceClient

func InitUserClient() {
	addr := os.Getenv("USER_SERVICE_ADDR")
	if addr == "" {
		addr = "user:8082"
		log.Println("[WARN] USER_SERVICE_ADDR not set, fallback to", addr)
	}

	conn, err := grpc.Dial(addr, grpc.WithInsecure(), grpc.WithBlock(), grpc.WithTimeout(5*time.Second))
	if err != nil {
		log.Fatalf("failed to connect to user service: %v", err)
	}

	UserClient = userpb.NewUserServiceClient(conn)
}

func GetProfile(ctx context.Context, id int64) (*userpb.Profile, error) {
    return UserClient.GetProfile(ctx, &userpb.ProfileRequest{UserId: id})
}


