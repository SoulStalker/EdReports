package auth

import (
	"context"

	asv1 "github.com/SoulStalker/EdReports/protos/gen/go/as"
	"google.golang.org/grpc"
)

type serverAPI struct {
	asv1.UnimplementedAuthServer
}

func Register(gRPC *grpc.Server) {
	asv1.RegisterAuthServer(gRPC, &serverAPI{})
}

func (s *serverAPI) Login(ctx context.Context, req *asv1.LoginRequest) (*asv1.LoginResponse, error) {
	return &asv1.LoginResponse{
		Token: req.GetEmail(),
	}, nil
}

func (s *serverAPI) Register(ctx context.Context, req *asv1.RegisterRequest) (*asv1.RegisterResponse, error) {
	panic("unimplemented")
}

func (s *serverAPI) IsAdmin(ctx context.Context, req *asv1.IsAdminRequest) (*asv1.IsAdminResponse, error) {
	panic("unimplemented")
}
