package main

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/lidi-a/project-protos/tasks-service/internal/database"
	"github.com/lidi-a/project-protos/tasks-service/internal/task"
	transportgrpc "github.com/lidi-a/project-protos/tasks-service/internal/transport/grpc"
)

func main() {
	// 0.
	if err := godotenv.Load("../../.env"); err != nil {
		log.Fatalf("failed to load environmental variables: %v", err)
	}
	// 1. Инициализация БД
	db, err := database.InitDB()
	if err != nil {
		log.Fatalf("Could not connect to database: %v", err)
	}

	// 2. Репозиторий и сервис задач
	repo := task.NewRepository(db)
	svc := task.NewService(repo)

	// 3. Клиент к Users-сервису
	userClient, conn, err := transportgrpc.NewUserClient(os.Getenv("ADDR_GRPC_USERS"))
	if err != nil {
		log.Fatalf("failed to connect to users: %v", err)
	}
	defer conn.Close()

	// 4. Запуск gRPC Tasks-сервиса
	log.Println("Starting Tasks gRPC server")
	if err := transportgrpc.RunGRPC(svc, userClient); err != nil {
		log.Fatalf("Tasks gRPC server error: %v", err)
	}
}
