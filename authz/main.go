package main

import (
	"authz/service"
	"fmt"
	"log"
	"net"
	"os"
	"proto/generated/authz"

	"github.com/joho/godotenv"
	"google.golang.org/grpc"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// db := utils.GetDB()
	// db.AutoMigrate(&models.Group{}, &models.Permission{})

	lis, err := net.Listen("tcp", fmt.Sprintf(":%s", os.Getenv("SERVER_PORT")))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()

	groupServer := service.GroupServer{}
	authz.RegisterGroupServiceServer(grpcServer, &groupServer)

	permissionServer := service.PermissionServer{}
	authz.RegisterPermissionServiceServer(grpcServer, &permissionServer)

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %s", err)
	}

	log.Println("Server is running")
}
