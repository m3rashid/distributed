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

func (r Relation) ToRelationName() string {
	switch r {
	case READER:
		return "reader"
	case EDITOR:
		return "editor"
	case ADMIN:
		return "admin"
	case OWNER:
		return "owner"
	default:
		return ""
	}
}

func (r Relation) TSName() string {
	return r.ToRelationName()
}

func (r Relation) IsValid() error {
	for _, relation := range ALL_RELATIONS {
		if r == relation {
			return nil
		}
	}
	return errors.New("invalid relation")
}
