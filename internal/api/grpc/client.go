package grpcclient

import (
    // "context"
    "log"
    "sync"
    "time"

    "google.golang.org/grpc"
    "github.com/sxd0/SweetTweet/proto/authpb"
    "github.com/sxd0/SweetTweet/proto/userpb"
)

var (
    once         sync.Once
    authClient   authpb.AuthServiceClient
    userClient   userpb.UserServiceClient
)

func ConnectAuth() authpb.AuthServiceClient {
    once.Do(func() {
        conn, err := grpc.Dial("auth:8081", grpc.WithInsecure(), grpc.WithBlock(), grpc.WithTimeout(3*time.Second))
        if err != nil {
            log.Fatalf("Failed to connect to auth service: %v", err)
        }
        authClient = authpb.NewAuthServiceClient(conn)
    })
    return authClient
}

func ConnectUser() userpb.UserServiceClient {
    once.Do(func() {
        conn, err := grpc.Dial("user:8082", grpc.WithInsecure(), grpc.WithBlock(), grpc.WithTimeout(3*time.Second))
        if err != nil {
            log.Fatalf("Failed to connect to user service: %v", err)
        }
        userClient = userpb.NewUserServiceClient(conn)
    })
    return userClient
}
