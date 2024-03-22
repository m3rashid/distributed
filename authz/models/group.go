package models

import (
	"errors"

	"github.com/go-playground/validator/v10"
)

type Group struct {
	ID   SqlID  `json:"id" gorm:"column:id;primary_key;not null" validate:"required"`
	Name string `json:"name" gorm:"column:name" validate:"required"`
}

const GROUP_MODEL_NAME = "groups"

func (*Group) TableName() string {
	return GROUP_MODEL_NAME
}

func (g *Group) IsValid() error {
	if g.Name == "" {
		return errors.New("invalid group name")
	}
	return nil
}

func (g *Group) Validate() error {
	validate := validator.New()
	if err := validate.Struct(g); err != nil {
		return err
	}

	if err := g.IsValid(); err != nil {
		return err
	}

	return nil
}
