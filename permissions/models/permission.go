package models

import (
	"errors"

	"github.com/go-playground/validator/v10"
)

/**
 * Permission definition
 *
 * A permission consists of 3 objects
 * <user>  ::  <relation>  ::  <resource>
 * user_id  	 	Relation       resource_id
 * group_id 		  enum				 resource_parent_id
 */

type SqlID uint

type Permission struct {
	ID               SqlID    `json:"id" gorm:"column:id;primary_key;not null" validate:"required"`
	ResourceID       SqlID    `json:"resource_id" gorm:"column:resource_id;not null" validate:"required"`
	ResourceParentID SqlID    `json:"resource_parent_id" gorm:"column:resource_parent_id" validate:""`
	Relation         Relation `json:"relation" gorm:"column:relation;not null" validate:"required"`
	UserID           SqlID    `json:"user_id" gorm:"column:user_id" validate:""`
	GroupID          SqlID    `json:"group_id" gorm:"column:group_id" validate:""`
}

const PERMISSION_MODEL_NAME = "permissions"

func (*Permission) TableName() string {
	return PERMISSION_MODEL_NAME
}

func (p *Permission) Validate() error {
	validate := validator.New()
	if err := validate.Struct(p); err != nil {
		return err
	}

	if err := p.Relation.IsValid(); err != nil {
		return err
	}

	if p.UserID == 0 && p.GroupID == 0 {
		return errors.New("user_id or group_id is required")
	}

	if p.ResourceID == 0 {
		return errors.New("resource_id is required")
	}

	return nil
}
