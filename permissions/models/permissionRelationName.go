package models

import "errors"

type RelationName string

const (
	READER_NAME RelationName = "reader"
	EDITOR_NAME RelationName = "editor"
	ADMIN_NAME  RelationName = "admin"
	OWNER_NAME  RelationName = "owner"
)

var ALL_RELATION_NAMES = []RelationName{READER_NAME, EDITOR_NAME, ADMIN_NAME, OWNER_NAME}

func (rn RelationName) ToRelation() Relation {
	switch rn {
	case READER_NAME:
		return READER
	case EDITOR_NAME:
		return EDITOR
	case ADMIN_NAME:
		return ADMIN
	case OWNER_NAME:
		return OWNER
	default:
		return 0
	}
}

func (rn RelationName) TSName() Relation {
	return rn.ToRelation()
}

func (rn RelationName) IsValid() error {
	for _, relationName := range ALL_RELATION_NAMES {
		if rn == relationName {
			return nil
		}
	}
	return errors.New("invalid relation name")
}
