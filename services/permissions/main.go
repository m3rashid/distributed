package main

import (
	"fmt"
	"io"
	"net"
	"os"
	"permissions/models"
	"permissions/rpcs"
	"permissions/utils"
	"proto/generated"

	"google.golang.org/grpc"
	"google.golang.org/grpc/grpclog"
)

func main() {
	// Adds gRPC internal logs. This is quite verbose, so adjust as desired!
	log := grpclog.NewLoggerV2(os.Stdout, io.Discard, io.Discard)
	grpclog.SetLoggerV2(log)

	err := utils.ParseEnv()
	if err != nil {
		panic(err)
	}

	db := utils.GetDB()
	db.AutoMigrate(&models.Group{}, &models.Permission{})

	lis, err := net.Listen("tcp", fmt.Sprintf(":%s", os.Getenv("SERVER_PORT")))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()

	groupServer := rpcs.GroupServer{}
	generated.RegisterGroupServiceServer(grpcServer, &groupServer)

	permissionServer := rpcs.PermissionServer{}
	generated.RegisterPermissionServiceServer(grpcServer, &permissionServer)

	log.Infoln("grpc server running on port", os.Getenv("SERVER_PORT"))
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %s", err)
	}

	log.Infoln("Server is running")
}
