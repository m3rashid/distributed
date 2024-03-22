package models

import "errors"

type Relation int

const (
	READER Relation = 1
	EDITOR Relation = 2
	ADMIN  Relation = 4
	OWNER  Relation = 8
)

var ALL_RELATIONS = []Relation{READER, EDITOR, ADMIN, OWNER}

func (r Relation) ToRelationName() (string, error) {
	switch r {
	case READER:
		return "reader", nil
	case EDITOR:
		return "editor", nil
	case ADMIN:
		return "admin", nil
	case OWNER:
		return "owner", nil
	default:
		return "", errors.New("invalid relation")
	}
}

func (r Relation) IsValid() error {
	for _, relation := range ALL_RELATIONS {
		if r == relation {
			return nil
		}
	}
	return errors.New("invalid relation")
}
