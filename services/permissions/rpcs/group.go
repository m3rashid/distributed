package rpcs

import (
	"context"
	"permissions/models"
	"permissions/utils"
	"proto/generated"
)

type GroupServer struct {
	generated.UnimplementedGroupServiceServer
}

func (s *GroupServer) DeleteGroup(
	ctx context.Context,
	msg *generated.ID,
) (*generated.Message, error) {
	db := utils.GetDB()
	if err := db.Where("id = ?", msg.Id).Update("deleted_at", true).Error; err != nil {
		return nil, err
	}
	return &generated.Message{Message: "successfully deleted"}, nil
}

func (s *GroupServer) CreateGroup(
	ctx context.Context,
	msg *generated.CreateGroupRequest,
) (*generated.Group, error) {
	db := utils.GetDB()
	group := models.Group{Name: msg.Name}
	if err := db.Create(&group).Error; err != nil {
		return nil, err
	}
	return &generated.Group{
		Name:      group.Name,
		Id:        uint32(group.ID),
		IsDeleted: group.IsDeleted,
	}, nil
}

func (s *GroupServer) GetAllGroups(
	ctx context.Context,
	msg *generated.PaginationRequest,
) (*generated.GroupPaginationResponse, error) {
	db := utils.GetDB()
	filter := models.Group{IsDeleted: false}

	var groups []models.Group
	if err := db.Limit(int(msg.Limit)).Offset(int(msg.Page) - 1).Where(filter).Find(&groups).Error; err != nil {
		return nil, err
	}

	var docsCount int64
	if err := db.Model(&models.Group{}).Where(filter).Count(&docsCount).Error; err != nil {
		return nil, err
	}

	var groupResponses []*generated.Group
	for _, group := range groups {
		groupResponses = append(groupResponses, &generated.Group{
			Name:      group.Name,
			Id:        uint32(group.ID),
			IsDeleted: group.IsDeleted,
		})
	}

	return &generated.GroupPaginationResponse{
		Data: groupResponses,
		Meta: &generated.PaginationResponse{
			Total:       uint32(len(groups)),
			Page:        msg.Page,
			Limit:       msg.Limit,
			HasNextPage: docsCount > int64(msg.Page)*int64(msg.Limit),
			HasPrevPage: msg.Page > 1,
		},
	}, nil
}
