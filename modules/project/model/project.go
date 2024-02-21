package model

import (
	"errors"
	"iTask/common"
	"time"
)

var (
	ErrNameCannotBeEmpty = errors.New("name cannot be empty")
	ErrProjectIsDeleted  = errors.New("project is deleted")
)

const (
	EntityName    = "Project"
	StatusDeleted = 3
)

type Project struct {
	common.SQLModel
	Name        string     `json:"name" gorm:"column:name"`
	Description string     `json:"description" gorm:"column:description"`
	Status      int        `json:"status" gorm:"column:status"`
	Thumbnail   string     `json:"thumbnail_url" gorm:"column:thumbnail_url"`
	Privacy     string     `json:"privacy" gorm:"column:privacy"`
	CreatedBy   int        `json:"created_by" gorm:"column:created_by"`
	Deadline    string     `json:"deadline" gorm:"column:deadline"`
	StartedAt   *time.Time `json:"started_at" gorm:"column:started_at"`
	TotalTasks  int        `json:"total_tasks" gorm:"-"`
	//Owner       Account    `json:"owner" gorm:"foreignKey:CreatedBy"`
}

func (Project) TableName() string {
	return "Projects"
}

func (a *Project) GetStatus() int {
	return a.Status
}

func (a *Project) GetCreatedBy() int {
	return a.CreatedBy
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
