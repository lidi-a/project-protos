package main

import (
	"log"

	"github.com/joho/godotenv"
	"github.com/lidi-a/project-protos/users-service/internal/database"
	transportgrpc "github.com/lidi-a/project-protos/users-service/internal/transport/grpc"
	"github.com/lidi-a/project-protos/users-service/internal/user"
)

func main() {

	if err := godotenv.Load("../../.env"); err != nil {
		log.Fatalf("failed to load environmental variables: %v", err)
	}

	db, err := database.InitDB()
	if err != nil {
		log.Fatalf("Could not connect to database: %v", err)
	}

	repo := user.NewRepository(db)
	svc := user.NewService(repo)

	log.Println("Starting Users gRPC server")
	if err := transportgrpc.RunGRPC(svc); err != nil {
		log.Fatalf("gRPC сервер завершился с ошибкой: %v", err)
	}
}
