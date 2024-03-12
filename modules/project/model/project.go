package model

import (
	"errors"
	"iTask/common"
	"iTask/modules/account/iomodel"
	"iTask/modules/project_members/model"
	"time"
)

var (
	ErrNameCannotBeEmpty = errors.New("name cannot be empty")
	ErrProjectIsDeleted  = errors.New("project is deleted")
)

type ProjectStatus int

const (
	EntityName = "Project"
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

type Project struct {
	common.SQLModel
	Name                string                 `json:"name" gorm:"column:name"`
	Description         string                 `json:"description" gorm:"column:description"`
	Status              ProjectStatus          `json:"status" gorm:"column:status"`
	Thumbnail           string                 `json:"thumbnail_url" gorm:"column:thumbnail_url"`
	Privacy             string                 `json:"privacy" gorm:"column:privacy"`
	CreatedBy           int                    `json:"-" gorm:"column:created_by"`
	Owner               *iomodel.SimpleAccount `json:"owner" gorm:"foreignKey:CreatedBy"`
	Deadline            string                 `json:"deadline" gorm:"column:deadline"`
	StartedAt           *time.Time             `json:"started_at" gorm:"column:started_at"`
	TotalTasks          int                    `json:"total_tasks" gorm:"-"`
	TotalCompletedTasks int                    `json:"completed_tasks" gorm:"-"`
	Members             *[]model.SimpleMember  `json:"members" gorm:"foreignKey:ProjectId"`
	Tags                string                 `json:"tags" gorm:"-"`
}

//Owner               iomodel.SimpleAccount `json:"owner" gorm:"foreignKey:CreatedBy"`

func (Project) TableName() string {
	return "Projects"
}

func (a *Project) GetStatus() ProjectStatus {
	return a.Status
}

func (a *Project) GetƠwner() *iomodel.SimpleAccount {
	return a.Owner
}

func (a *Project) GetID() int {
	return a.Id
}

var MapPriority map[int]string = map[int]string{
	1: "High",
	2: "Medium",
	3: "Low",
}

var MapProjectStatus map[int]string = map[int]string{
	3: "Deleted",
	2: "Completed",
	1: "Incomplete",
}
