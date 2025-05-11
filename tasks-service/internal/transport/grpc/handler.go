package grpc

import (
	"context"

	taskpb "github.com/lidi-a/project-protos/proto/task"
	userpb "github.com/lidi-a/project-protos/proto/user"
	"github.com/lidi-a/project-protos/tasks-service/internal/task"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
)

type Handler struct {
	svc        task.Service
	userClient userpb.UserServiceClient
	taskpb.UnimplementedTaskServiceServer
}

func NewHandler(svc task.Service, uc userpb.UserServiceClient) *Handler {
	return &Handler{svc: svc, userClient: uc}
}

func (h *Handler) CreateTask(ctx context.Context, req *taskpb.CreateTaskRequest) (*taskpb.CreateTaskResponse, error) {

	// 1. Проверить пользователя:
	if _, err := h.userClient.GetUser(ctx, &userpb.GetUserRequest{Id: req.UserId}); err != nil {
		return nil, status.Errorf(codes.NotFound, "user %v not found", req.UserId)
	}

	// 2. Внутренняя логика:
	t, err := h.svc.CreateTask(task.Task{UserID: req.UserId, Title: req.Title})
	if err != nil {
		return nil, err
	}

	// 3. Ответ:
	return &taskpb.CreateTaskResponse{Task: &taskpb.Task{Id: t.ID, UserId: t.UserID, Title: t.Title, IsDone: t.IsDone}}, nil
}

func (h *Handler) GetTask(ctx context.Context, req *taskpb.GetTaskRequest) (*taskpb.GetTaskResponse, error) {

	task, err := h.svc.GetTaskByID(req.Id)
	if err != nil {
		return nil, status.Errorf(codes.NotFound, "task with ID %s not found", req.Id)
	}

	return &taskpb.GetTaskResponse{
		Task: &taskpb.Task{
			Id:     task.ID,
			Title:  task.Title,
			IsDone: task.IsDone,
			UserId: task.UserID,
		}}, nil
}

func (h *Handler) ListTasks(ctx context.Context, e *emptypb.Empty) (*taskpb.ListTasksResponse, error) {

	domainTasks, err := h.svc.GetAllTasks()
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}
	var tasks []*taskpb.Task
	for _, domainTask := range domainTasks {
		task := &taskpb.Task{
			Id:     domainTask.ID,
			Title:  domainTask.Title,
			IsDone: domainTask.IsDone,
			UserId: domainTask.UserID,
		}
		tasks = append(tasks, task)
	}

	return &taskpb.ListTasksResponse{Tasks: tasks}, nil
}

func (h *Handler) ListTasksByUser(ctx context.Context, req *taskpb.ListTasksByUserRequest) (*taskpb.ListTasksResponse, error) {

	domainTasks, err := h.svc.GetTasksByUserID(req.UserId)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}
	var tasks []*taskpb.Task
	for _, domainTask := range domainTasks {
		task := &taskpb.Task{
			Id:     domainTask.ID,
			Title:  domainTask.Title,
			IsDone: domainTask.IsDone,
			UserId: domainTask.UserID,
		}
		tasks = append(tasks, task)
	}

	return &taskpb.ListTasksResponse{Tasks: tasks}, nil
}

func (h *Handler) UpdateTask(ctx context.Context, req *taskpb.UpdateTaskRequest) (*taskpb.UpdateTaskResponse, error) {

	// Проверяем пользователя
	if req.UserId != nil {
		if _, err := h.userClient.GetUser(ctx, &userpb.GetUserRequest{Id: *req.UserId}); err != nil {
			return nil, status.Errorf(codes.NotFound, "user %v not found", req.UserId)
		}
	}

	task := task.Task{
		Title: req.Title,
		IsDone: req.IsDone,
	}

	if req.UserId != nil {
		task.UserID = *req.UserId
	}

	task, err := h.svc.UpdateTaskByID(req.Id, task)
	if err != nil {
		return nil, status.Errorf(codes.NotFound, "task with ID %s not found", req.Id)
	}

	return &taskpb.UpdateTaskResponse{
		Task: &taskpb.Task{
			Id:     task.ID,
			Title:  task.Title,
			IsDone: task.IsDone,
			UserId: task.UserID,
		},
	}, nil
}

func (h *Handler) DeleteTask(ctx context.Context, req *taskpb.DeleteTaskRequest) (*taskpb.DeleteTaskResponse, error) {

	err := h.svc.DeleteTaskByID(req.Id)
	if err != nil {
		return nil, status.Errorf(codes.NotFound, "task with ID %s not found", req.Id)
	}

	return &taskpb.DeleteTaskResponse{}, nil
}
