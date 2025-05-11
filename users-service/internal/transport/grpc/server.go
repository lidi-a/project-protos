package grpc

import (
	"net"
	"os"

	userpb "github.com/lidi-a/project-protos/proto/user"
	"github.com/lidi-a/project-protos/users-service/internal/user"
	"google.golang.org/grpc"
)

func RunGRPC(svc user.Service) error {
	// 1. net.Listen на ":50051"
	listener, err := net.Listen("tcp", os.Getenv("ADDR_GRPC_USERS"))
	if err != nil {
		return err
	}
	// 2.
	grpcSrv := grpc.NewServer()
	// 3.
	userpb.RegisterUserServiceServer(grpcSrv, NewHandler(svc))
	// 4.
	return grpcSrv.Serve(listener)
}
