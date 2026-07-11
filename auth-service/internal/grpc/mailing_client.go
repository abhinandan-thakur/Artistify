package grpc

import (
	"context"
	"time"

	pb "github.com/abhinandan-thakur/Artistify/auth-service/proto"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"github.com/abhinandan-thakur/Artistify/auth-service/internal/config"

)

type MailingClient struct {
	Client pb.MailingServiceClient
}

func NewMailingClient() (*MailingClient, error,) {
	config := config.LoadConfig()

	conn, err := grpc.NewClient(config.MailingGRPCHost+":"+config.MailingGRPCPort,grpc.WithTransportCredentials(insecure.NewCredentials(),),)

	if err != nil {
		return nil, err
	}

	client := pb.NewMailingServiceClient(conn,)

	return &MailingClient{
		Client: client,
	}, nil
}

func (m *MailingClient) SendMail(receiverEmail string, otp string,) error {

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second,)

	defer cancel()

	_, err := m.Client.SendMail(ctx, &pb.MailingServiceRequest{
				ReceiverEmail:	receiverEmail,
				Otp: otp,
			},
		)

	return err
}