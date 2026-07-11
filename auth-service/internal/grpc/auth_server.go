package grpc

import (
	"context"
	pb "github.com/abhinandan-thakur/Artistify/auth-service/proto"
	"github.com/golang-jwt/jwt/v5"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"log"
	"github.com/abhinandan-thakur/Artistify/auth-service/internal/config"

)

type AuthServer struct{
	pb.UnimplementedAuthenticateServiceServer
	Config *config.Config
}

func (server *AuthServer) IsValidToken(ctx context.Context, request *pb.IsValidTokenRequest) (*pb.IsValidTokenResponse, error) {
	return nil, status.Error(codes.Unimplemented, "method IsValidToken not implemented")
}
func (server *AuthServer) GetRole(ctx context.Context, request *pb.GetRoleRequest) (*pb.GetRoleResponse, error) {
	return nil, status.Error(codes.Unimplemented, "method GetRole not implemented")
}
func (server *AuthServer) IsValidAndGetRole(ctx context.Context, request *pb.IsValidAndGetRoleRequest) (*pb.IsValidAndGetRoleResponse, error) {

	jwtSecret := []byte(server.Config.JWTSecret)
	log.Println("The jwt secre")

	tokenString := request.TokenString
	log.Println("the jwt secret is:", jwtSecret)
	claims := jwt.MapClaims{}

	token, err := jwt.ParseWithClaims(
		tokenString,
		claims,
		func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, jwt.ErrTokenSignatureInvalid
			}
			return jwtSecret, nil
		},
	)

	log.Println("The token is:", token)

	if err != nil || !token.Valid {
		return &pb.IsValidAndGetRoleResponse{
			Valid: false,
			Role: "user",
		}, nil
	}

	claims = token.Claims.(jwt.MapClaims)

	log.Println("The claim is:", claims)

	role, _ := claims["role"].(string)

	return &pb.IsValidAndGetRoleResponse{
		Valid: true,
		Role: role,
	}, nil

}