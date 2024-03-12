package model

import (
	"errors"
	"iTask/common"
)

var (
	ErrNameCannotBeEmpty = errors.New("name cannot be empty")
	ErrTagIsDeleted      = errors.New("Tag is deleted")
)

type TagStatus int
type TagType int

const (
	EntityName = "Tag"
)

const (
	StatusPending TagStatus = 1 + iota
	StatusInProgress
	StatusDone
	StatusDeleted
)

const (
	TypeProject TagType = 1 + iota
	TypeTask
)

func (status TagStatus) String() string {
	switch status {
	case StatusPending:
		return "Pending"
	case StatusInProgress:
		return "In Progress"
	case StatusDone:
		return "Done"
	case StatusDeleted:
		return "Deleted"
	default:
		return "Unknown"
	}
}

func (tagType TagType) String() string {
	switch tagType {
	case TypeProject:
		return "Project"
	case TypeTask:
		return "Task"
	default:
		return "Unknown"
	}
}

type Tag struct {
	common.SQLModel
	Name        string  `json:"name" gorm:"column:name"`
	Description string  `json:"description" gorm:"column:description"`
	Position    float64 `json:"position" gorm:"column:position"`
	TagType     TagType `json:"tag_type" gorm:"column:tag_type"`
}

//Owner               iomodel.SimpleAccount `json:"owner" gorm:"foreignKey:CreatedBy"`

func (Tag) TableName() string {
	return "Tags"
}

func (a *Tag) GetType() TagType {
	return a.TagType
}

func (a *Tag) GetID() int {
	return a.Id
}
