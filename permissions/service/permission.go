package service

import (
	"context"
	"fmt"
	"proto/generated/permissions"
)

type PermissionServer struct {
	permissions.UnimplementedPermissionServiceServer
}

func (s *PermissionServer) DeletePermission(
	ctx context.Context,
	msg *permissions.ID,
) (*permissions.Message, error) {
	fmt.Println("DeletePermission received: ", msg)
	return nil, nil
}

func (s *PermissionServer) CreatePermission(
	ctx context.Context,
	msg *permissions.CreatePermissionRequest,
) (*permissions.Permission, error) {
	fmt.Println("CreatePermission received: ", msg)
	return nil, nil
}

func (s *PermissionServer) CheckPermission(
	ctx context.Context,
	msg *permissions.CreatePermissionRequest,
) (*permissions.CheckPermissionResponse, error) {
	fmt.Println("CheckPermission received: ", msg)
	return nil, nil
}
