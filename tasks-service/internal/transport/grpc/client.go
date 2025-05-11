package grpc

import (
	"log"

	userpb "github.com/lidi-a/project-protos/proto/user"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func NewUserClient(addr string) (userpb.UserServiceClient, *grpc.ClientConn, error) {
	
	conn, err := grpc.NewClient(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatal("Could not establish connection with gRPC users service")
	}
	client := userpb.NewUserServiceClient(conn)
	return client, conn, err
}
