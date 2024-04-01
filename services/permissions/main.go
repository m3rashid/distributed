package main

import (
	"context"
	"fmt"
	"io"
	"net"
	"net/http"
	"os"
	"permissions/models"
	"permissions/rpcs"
	"permissions/utils"
	"proto/generated"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/grpclog"
	"google.golang.org/grpc/reflection"
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

	go func() {
		if err := runHttpServer(); err != nil {
			grpclog.Fatal(err)
		}
	}()

	if err := runGrpcServer(); err != nil {
		grpclog.Fatal(err)
	}
}

func runGrpcServer() error {
	grpcPort := fmt.Sprintf(":%s", os.Getenv("GRPC_SERVER_PORT"))
	lis, err := net.Listen("tcp", grpcPort)
	if err != nil {
		grpclog.Infoln("failed to listen: %v", err)
		return err
	}

	grpcServer := grpc.NewServer()

	groupServer := rpcs.GroupServer{}
	generated.RegisterGroupServiceServer(grpcServer, &groupServer)

	permissionServer := rpcs.PermissionServer{}
	generated.RegisterPermissionServiceServer(grpcServer, &permissionServer)

	// Enable reflection to allow clients to query the server's services
	reflection.Register(grpcServer)

	grpclog.Infoln("grpc server running on port", grpcPort)
	if err := grpcServer.Serve(lis); err != nil {
		grpclog.Infoln("failed to serve: %s", err)
		return err
	}
	return nil
}

func runHttpServer() error {
	httpPort := fmt.Sprintf(":%s", os.Getenv("HTTP_SERVER_PORT"))
	grpcPort := fmt.Sprintf(":%s", os.Getenv("GRPC_SERVER_PORT"))

	ctx := context.Background()
	mux := runtime.NewServeMux()

	conn, err := grpc.DialContext(ctx, "localhost"+grpcPort, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		grpclog.Infoln("failed to dial server: %v", err)
		return err
	}

	if err := generated.RegisterGroupServiceHandler(ctx, mux, conn); err != nil {
		grpclog.Infoln("failed to register group handler: %v", err)
		return err
	}

	if err := generated.RegisterPermissionServiceHandler(ctx, mux, conn); err != nil {
		grpclog.Infoln("failed to register permissions handler: %v", err)
		return err
	}

	grpclog.Infoln("http server running on port", httpPort)
	if err := http.ListenAndServe(httpPort, mux); err != nil {
		grpclog.Infoln("failed to serve: %s", err)
		return err
	}

	return nil
}
