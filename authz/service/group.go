package service

import (
	"context"
	"fmt"
	"proto/generated/authz"
)

type GroupServer struct {
	authz.UnimplementedGroupServiceServer
}

func (s *GroupServer) DeleteGroup(
	ctx context.Context,
	msg *authz.ID,
) (*authz.Message, error) {
	fmt.Println("DeleteGroup received: ", msg)
	return nil, nil
}

func (s *GroupServer) CreateGroup(
	ctx context.Context,
	msg *authz.CreateGroupRequest,
) (*authz.Group, error) {
	fmt.Println("CreateGroup received: ", msg)
	fmt.Println(msg.Name)
	return nil, nil
}

func (s *GroupServer) GetAllGroups(
	ctx context.Context,
	msg *authz.PaginationRequest,
) (*authz.GroupPaginationResponse, error) {
	fmt.Println("GetAllGroups received: ", msg)
	return nil, nil
}
