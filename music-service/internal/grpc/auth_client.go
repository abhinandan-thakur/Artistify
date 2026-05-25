package grpc

import (
	"log"

	pb "github.com/abhinandan-thakur/Artistify/music-service/proto"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func NewAuthClient() pb.AuthenticateServiceClient {
	log.Println("Entered NewAuthClient()")

	conn, err := grpc.NewClient(
		"localhost:50051",
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

	client :=pb.NewAuthenticateServiceClient(conn)

	log.Println("exited NewAuthClient()", client)

	return client
}