package rpcs

import (
	"context"
	"fmt"
	"permissions/models"
	"permissions/utils"
	"proto/generated"
)

type PermissionServer struct {
	generated.UnimplementedPermissionServiceServer
}

func (s *PermissionServer) DeletePermission(
	ctx context.Context,
	msg *generated.ID,
) (*generated.Message, error) {
	db := utils.GetDB()
	if err := db.Where("id = ?", msg.Id).Update("deleted_at", true).Error; err != nil {
		return nil, err
	}
	return &generated.Message{Message: "successfully deleted"}, nil
}

func (s *PermissionServer) CreatePermission(
	ctx context.Context,
	msg *generated.CreatePermissionRequest,
) (*generated.Permission, error) {
	db := utils.GetDB()
	permission := models.Permission{
		ResourceID:       models.SqlID(msg.ResourceId),
		ResourceParentID: models.SqlID(*msg.ResourceParentId),
		Relation:         models.Relation(msg.Relation),
		UserID:           models.SqlID(msg.UserId),
		GroupID:          models.SqlID(*msg.GroupId),
	}
	if err := db.Create(&permission).Error; err != nil {
		return nil, err
	}

	return &generated.Permission{
		Id:               uint32(permission.ID),
		ResourceId:       uint32(permission.ResourceID),
		UserId:           uint32(permission.UserID),
		GroupId:          uint32(permission.GroupID),
		Relation:         uint32(permission.Relation),
		ResourceParentId: uint32(permission.ResourceParentID),
		IsDeleted:        permission.IsDeleted,
	}, nil
}

func (s *PermissionServer) CheckPermission(
	ctx context.Context,
	msg *generated.CreatePermissionRequest,
) (*generated.CheckPermissionResponse, error) {
	fmt.Println("CheckPermission called", msg)
	// Dummy response
	return &generated.CheckPermissionResponse{Res: true}, nil
}
