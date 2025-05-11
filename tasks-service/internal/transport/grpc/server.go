package grpc

import (
	"net"
	"os"

	userpb "github.com/lidi-a/project-protos/proto/user"
	taskpb "github.com/lidi-a/project-protos/proto/task"
	"github.com/lidi-a/project-protos/tasks-service/internal/task"
	"google.golang.org/grpc"
)

func RunGRPC(svc task.Service, uc userpb.UserServiceClient) error {

	listener, err := net.Listen("tcp", os.Getenv("ADDR_GRPC_TASKS"))
	if err != nil {
		return err
	}

	grpcSrv := grpc.NewServer()

	handler := NewHandler(svc, uc)
	taskpb.RegisterTaskServiceServer(grpcSrv, handler)
	return grpcSrv.Serve(listener)
}