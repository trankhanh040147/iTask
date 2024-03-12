package model

import (
	"errors"
	"iTask/common"
)

var (
	ErrNameCannotBeEmpty = errors.New("name cannot be empty")
	ErrProjectIsDeleted  = errors.New("project is deleted")
)

type ProjectStatus int

const (
	EntityName = "project_tags"
)

const (
	StatusPending ProjectStatus = 1 + iota
	StatusInProgress
	StatusDone
	StatusDeleted
)

func (status ProjectStatus) String() string {
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

type ProjectTag struct {
	common.SQLModel
	ProjectId int `json:"project_id" gorm:"column:project_id"`
	TagId     int `json:"tag_id" gorm:"column:tag_id"`
}

func (ProjectTag) TableName() string {
	return "ProjectTags"
}

func (ProjectTag) GetEntityName() string {
	return EntityName
}

func (a *ProjectTag) GetId() int {
	return a.Id
}
