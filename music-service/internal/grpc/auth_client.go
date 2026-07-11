package grpc

import (
	"log"

	"github.com/abhinandan-thakur/Artistify/music-service/internal/config"
	pb "github.com/abhinandan-thakur/Artistify/music-service/proto"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func NewAuthClient(
	config *config.Config,
) pb.AuthenticateServiceClient {

	log.Println("Entered NewAuthClient()")

	conn, err := grpc.NewClient(
		config.GRPCHost+":"+config.GRPCPort,
		grpc.WithTransportCredentials(
			insecure.NewCredentials(),
		),
	)

	if err != nil {
		log.Fatal(
			"failed to connect to auth-service:",
			err,
		)
	}

	client := pb.NewAuthenticateServiceClient(conn)

	log.Println(
		"Connected to:",
		config.GRPCHost+":"+config.GRPCPort,
	)

	return client
}