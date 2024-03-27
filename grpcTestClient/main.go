package main

import (
	"context"
	"log"
	"proto/generated/permissions"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	var conn *grpc.ClientConn
	conn, err := grpc.Dial(":4001", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}

	defer conn.Close()

	groupClient := permissions.NewGroupServiceClient(conn)
	// permissionClient := permissions.NewPermissionServiceClient(conn)

	res, err := groupClient.CreateGroup(context.Background(), &permissions.CreateGroupRequest{
		Name: "test",
	})

	if err != nil {
		log.Fatalf("Error when calling GetAllGroups: %v", err)
	}
	log.Printf("Response from GetAllGroups: %v", res)
}
