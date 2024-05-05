package grpc

import (
	"context"
	"fmt"
	"net"

	"github.com/cmclaughlin24/soteria-be/apps/iam/internal/core/ports"
	"github.com/cmclaughlin24/soteria-be/pkg/iam"
	"google.golang.org/grpc"
)

type IamGRPCServer struct {
	iam.UnimplementedIamServiceServer
	services *ports.Services
}

func NewIamGRPCServer(services *ports.Services) *IamGRPCServer {
	return &IamGRPCServer{
		services: services,
	}
}

func (s *IamGRPCServer) Listen(port string) {
	listen, err := net.Listen("tcp", fmt.Sprintf(":%s", port))

	if err != nil {
		panic(err)
	}

	defer listen.Close()

	server := grpc.NewServer()

	iam.RegisterIamServiceServer(server, s)

	if err := server.Serve(listen); err != nil {
		panic(err)
	}
}

func (s *IamGRPCServer) VerifyAccessToken(ctx context.Context, req *iam.VerifyTokenRequest) (*iam.AccessTokenClaimsReponse, error) {
	token := req.GetToken()

	claims, err := s.services.Authentication.VerifyAccessToken(ctx, token)

	if err != nil {
		return nil, err
	}

	return &iam.AccessTokenClaimsReponse{
		Sub:                  claims.Subject,
		Name:                 claims.Name,
		AuthorizationDetails: claims.AuthorizationDetails,
		ExpiresAt:            claims.ExpiresAt.Unix(),
	}, nil
}

func (s *IamGRPCServer) VerifyApiKey(ctx context.Context, req *iam.VerifyApiKeyRequest) (*iam.ApiKeyClaimsReponse, error) {
	apiKey := req.GetKey()

	claims, err := s.services.ApiKey.VerifyApiKey(ctx, apiKey)

	if err != nil {
		return nil, err
	}

	return &iam.ApiKeyClaimsReponse{
		Sub:                  claims.Sub,
		Name:                 claims.Name,
		AuthorizationDetails: claims.AuthorizationDetails,
		ExpiresAt:            claims.ExpiresAt,
	}, nil
}
