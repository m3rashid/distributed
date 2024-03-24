package main

import (
	"fmt"
	"log"
	"net"
	"os"
	"permissions/models"
	"permissions/service"
	"permissions/utils"
	"proto/generated/permissions"

	"github.com/joho/godotenv"
	"google.golang.org/grpc"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	db := utils.GetDB()
	db.AutoMigrate(&models.Group{}, &models.Permission{})

	lis, err := net.Listen("tcp", fmt.Sprintf(":%s", os.Getenv("SERVER_PORT")))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()

	groupServer := service.GroupServer{}
	permissions.RegisterGroupServiceServer(grpcServer, &groupServer)

	permissionServer := service.PermissionServer{}
	permissions.RegisterPermissionServiceServer(grpcServer, &permissionServer)

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %s", err)
	}

	log.Println("Server is running")
}
