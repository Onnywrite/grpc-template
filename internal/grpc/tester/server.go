package tester

import (
	"context"
	"fmt"

	"github.com/Onnywrite/grpc-template/gen"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
)

type GRPCServer struct {
	gen.UnimplementedTesterServer
}

func (s *GRPCServer) Add(ctx context.Context, req *gen.AddRequest) (*gen.AddResponse, error) {
	return &gen.AddResponse{Result: req.X + req.Y}, nil
}

func (s *GRPCServer) GetError(context.Context, *emptypb.Empty) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.InvalidArgument, `{
		"Code": "666999",
		"Errors": [
			"Field is invalid"
			]
		}`)
}
func (s *GRPCServer) SayHello(c context.Context, name *gen.Name) (*gen.Hello, error) {
	return &gen.Hello{
		Message: fmt.Sprintf("Hello, %s!", name.Name),
	}, nil
}
