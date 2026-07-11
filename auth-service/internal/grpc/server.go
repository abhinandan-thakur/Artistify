package grpc

import (
	"log"
	"net"

	pb "github.com/abhinandan-thakur/Artistify/auth-service/proto"

	"google.golang.org/grpc"
	"github.com/abhinandan-thakur/Artistify/auth-service/internal/config"
)


func StartGRPCServer(config *config.Config) error {

	grpcServer := grpc.NewServer()

	authServer := &AuthServer{Config: config}

	pb.RegisterAuthenticateServiceServer(grpcServer, authServer)

	listener, err := net.Listen("tcp", ":"+config.GRPCPort)

	if err != nil {	return err}

	log.Println("gRPC running on :50051")

	go grpcServer.Serve(listener)

	return nil
}