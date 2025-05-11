package grpc

import (
	"context"

	"github.com/google/uuid"
	userpb "github.com/lidi-a/project-protos/proto/user"
	"github.com/lidi-a/project-protos/users-service/internal/user"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
)

type Handler struct {
	svc user.Service
	userpb.UnimplementedUserServiceServer
}

func NewHandler(svc user.Service) *Handler {
	return &Handler{svc: svc}
}

func (h *Handler) CreateUser(ctx context.Context, req *userpb.CreateUserRequest) (*userpb.CreateUserResponse, error) {

	user := user.User{
		ID:       uuid.NewString(),
		Email:    req.Email,
		Password: req.Password,
	}

	user, err := h.svc.CreateUser(user)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &userpb.CreateUserResponse{
		User: &userpb.User{
			Id:       user.ID,
			Email:    user.Email,
			Password: user.Password,
		}}, nil

}
func (h *Handler) GetUser(ctx context.Context, req *userpb.GetUserRequest) (*userpb.GetUserResponse, error) {

	user, err := h.svc.GetUserByID(req.Id)
	if err != nil {
		return nil, status.Errorf(codes.NotFound, "resource with ID %s not found", req.Id)
	}

	return &userpb.GetUserResponse{
		User: &userpb.User{
			Id: user.ID,
			Email: user.Email,
			Password: user.Password,
		}}, nil

}
func (h *Handler) ListUsers(ctx context.Context, e *emptypb.Empty) (*userpb.ListUsersResponse, error) {

	domainUsers, err := h.svc.GetAllUsers()
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}
	var users []*userpb.User
	for _, domainUser := range domainUsers {
		user := &userpb.User{
			Id: domainUser.ID,
			Email: domainUser.Email,
			Password: domainUser.Password,
		}
		users = append(users, user)
	}

	return &userpb.ListUsersResponse{Users: users}, nil
}

func (h *Handler) UpdateUser(ctx context.Context, req *userpb.UpdateUserRequest) (*userpb.UpdateUserResponse, error) {

	user := user.User{
		Email: req.Email,
		Password: req.Password,
	}

	user, err := h.svc.UpdateUserByID(req.Id, user)
	if err != nil {
		return nil, status.Errorf(codes.NotFound, "resource with ID %s not found", req.Id)
	}

	return &userpb.UpdateUserResponse{
		User: &userpb.User{
			Id: user.ID,
			Email: user.Email,
			Password: user.Password,
		},
	}, nil

}
func (h *Handler) DeleteUser(ctx context.Context, req *userpb.DeleteUserRequest) (*userpb.DeleteUserResponse, error) {

	err := h.svc.DeleteUserByID(req.Id)
	if err != nil {
		return nil, status.Errorf(codes.NotFound, "resource with ID %s not found", req.Id)
	}

	return &userpb.DeleteUserResponse{}, nil
}


	