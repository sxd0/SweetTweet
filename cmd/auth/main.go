package main

import (
    "fmt"
    "log"
    "net"
    "os"

    "google.golang.org/grpc"
    "google.golang.org/grpc/reflection"

    "github.com/sxd0/SweetTweet/internal/auth/repository"
    "github.com/sxd0/SweetTweet/internal/auth/service"
    pb "github.com/sxd0/SweetTweet/proto/authpb"
    "github.com/sxd0/SweetTweet/pkg/config"
    "github.com/sxd0/SweetTweet/pkg/db"
    "github.com/sxd0/SweetTweet/pkg/logger"
)

func main() {
    cfg := config.Load()
    logg := logger.NewLogger()

    database, err := db.Connect()
    if err != nil {
        logg.Error("DB connection failed", "error", err)
        return
    }
    logg.Info("Connected to DB")

    jwtSecret := os.Getenv("JWT_SECRET")
    if jwtSecret == "" {
        jwtSecret = "supersecret"
        log.Println("[WARN] JWT_SECRET not set, fallback to", jwtSecret)
    }

    repo := repository.NewUserRepository(database)
    authSvc := service.NewAuthService(repo, jwtSecret)

    lis, err := net.Listen("tcp", fmt.Sprintf(":%s", cfg.Port))
    if err != nil {
        logg.Error("Failed to listen", "error", err)
        return
    }

    grpcServer := grpc.NewServer()
    pb.RegisterAuthServiceServer(grpcServer, authSvc)

    reflection.Register(grpcServer)

    logg.Info("Auth gRPC service running", "port", cfg.Port)
    if err := grpcServer.Serve(lis); err != nil {
        logg.Error("gRPC server error", "error", err)
    }
}
