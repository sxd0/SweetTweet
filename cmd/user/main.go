package main

import (
    "context"
    "log"
    "net"
    "os"

	"google.golang.org/grpc/reflection"
    "google.golang.org/grpc"
    "github.com/sxd0/SweetTweet/internal/user/handler"
    "github.com/sxd0/SweetTweet/internal/user/repository"
    "github.com/sxd0/SweetTweet/internal/user/service"
    pb "github.com/sxd0/SweetTweet/proto/userpb"
    "go.mongodb.org/mongo-driver/mongo"
    "go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
    ctx := context.Background()

    mongoURI := os.Getenv("MONGO_URI")
    if mongoURI == "" {
        mongoURI = "mongodb://localhost:27017"
    }

    client, err := mongo.Connect(ctx, options.Client().ApplyURI(mongoURI))
    if err != nil {
        log.Fatalf("Mongo connect error: %v", err)
    }

    repo := repository.NewProfileRepository(client)
    svc := service.NewProfileService(repo)
    grpcHandler := handler.NewGRPCHandler(svc)

    lis, err := net.Listen("tcp", ":8082")
    if err != nil {
        log.Fatalf("Failed to listen: %v", err)
    }

	grpcServer := grpc.NewServer()
	pb.RegisterUserServiceServer(grpcServer, grpcHandler)

	reflection.Register(grpcServer)

	log.Println("gRPC UserService running on :8082")
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
