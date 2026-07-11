package grpc

import (
	"log"
	"net"

	pb "github.com/abhinandan-thakur/Artistify/auth-service/proto"

	"github.com/abhinandan-thakur/Artistify/auth-service/internal/config"
	"google.golang.org/grpc"
)

func StartGRPCServer(config *config.Config) error {

	grpcServer := grpc.NewServer()

	authServer := &AuthServer{Config: config}

	pb.RegisterAuthenticateServiceServer(grpcServer, authServer)

	listener, err := net.Listen("tcp", ":"+config.GRPCPort)

	if err != nil {
		return err
	}

	log.Println("gRPC running on :", config.GRPCPort)

	// go grpcServer.Serve(listener)

	go func() {
		err := grpcServer.Serve(listener)

		if err != nil {
			log.Println("GRPC Server Stopped", err)
		}
	}()

	return nil
}
