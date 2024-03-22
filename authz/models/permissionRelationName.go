package models

import "errors"

type RelationName string

const (
	READER_NAME RelationName = "reader"
	EDITOR_NAME RelationName = "editor"
	ADMIN_NAME  RelationName = "admin"
	OWNER_NAME  RelationName = "owner"
)

func (rn RelationName) ToRelation() (Relation, error) {
	switch rn {
	case READER_NAME:
		return READER, nil
	case EDITOR_NAME:
		return EDITOR, nil
	case ADMIN_NAME:
		return ADMIN, nil
	case OWNER_NAME:
		return OWNER, nil
	default:
		return 0, errors.New("invalid relation name")
	}
}
