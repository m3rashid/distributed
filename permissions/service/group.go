package service

import (
	"context"
	"fmt"
	"proto/generated/permissions"
)

type GroupServer struct {
	permissions.UnimplementedGroupServiceServer
}

func (s *GroupServer) DeleteGroup(
	ctx context.Context,
	msg *permissions.ID,
) (*permissions.Message, error) {
	fmt.Println("DeleteGroup received: ", msg)
	return nil, nil
}

func (s *GroupServer) CreateGroup(
	ctx context.Context,
	msg *permissions.CreateGroupRequest,
) (*permissions.Group, error) {
	fmt.Println("CreateGroup received: ", msg)
	fmt.Println(msg.Name)
	return nil, nil
}

func (s *GroupServer) GetAllGroups(
	ctx context.Context,
	msg *permissions.PaginationRequest,
) (*permissions.GroupPaginationResponse, error) {
	fmt.Println("GetAllGroups received: ", msg)
	return nil, nil
}
