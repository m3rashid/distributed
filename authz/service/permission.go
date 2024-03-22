package service

import (
	"context"
	"fmt"
	"proto/generated/authz"
)

type PermissionServer struct{}

func (s *PermissionServer) DeletePermission(
	ctx context.Context,
	msg *authz.ID,
) (*authz.Message, error) {
	fmt.Println("DeletePermission received: ", msg)
	return nil, nil
}

func (s *PermissionServer) CreatePermission(
	ctx context.Context,
	msg *authz.CreatePermissionRequest,
) (*authz.Permission, error) {
	fmt.Println("CreatePermission received: ", msg)
	return nil, nil
}

func (s *PermissionServer) CheckPermission(
	ctx context.Context,
	msg *authz.CreatePermissionRequest,
) (*authz.CheckPermissionResponse, error) {
	fmt.Println("CheckPermission received: ", msg)
	return nil, nil
}
