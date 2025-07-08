package handler

import (
    "github.com/sxd0/SweetTweet/internal/auth/service"
    pb "github.com/sxd0/SweetTweet/proto/authpb"
)

func RegisterGRPC(server pb.AuthServiceServer) pb.AuthServiceServer {
    return server
}

func NewGRPCHandler(svc *service.AuthService) pb.AuthServiceServer {
    return svc
}
