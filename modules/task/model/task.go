package model

import (
	"errors"
	"iTask/common"
	"iTask/entities"
	"time"
)

var (
	ErrNameCannotBeEmpty = errors.New("name cannot be empty")
	ErrTaskIsDeleted     = errors.New("task is deleted")
)

const (
	EntityName = "Task"
)

type Task struct {
	common.SQLModel
	Name        string            `json:"name" gorm:"column:name"`
	Description string            `json:"description" gorm:"column:description"`
	Status      int               `json:"status" gorm:"column:status"`
	ProjectId   int               `json:"project_id" gorm:"column:project_id"`
	CreatedBy   int               `json:"created_by" gorm:"column:created_by"`
	ParentTask  int               `json:"parent_task_id" gorm:"column:parent_task_id"`
	Position    float64           `json:"position" gorm:"column:position"`
	Priority    int               `json:"priority" gorm:"column:priority"`
	Completed   bool              `json:"completed" gorm:"column:completed"`
	DueDate     *time.Time        `json:"due_date" gorm:"column:due_date"`
	StartedAt   *time.Time        `json:"started_at" gorm:"column:started_at"`
	CompletedAt *time.Time        `json:"completed_at" gorm:"column:completed_at"`
	Owner       *entities.Account `json:"owner" gorm:"foreignKey:CreatedBy"`
}

func (Task) TableName() string {
	return "Tasks"
}

func (a *Task) GetStatus() int {
	return a.Status
}

func (a *Task) GetCreatedBy() int {
	return a.CreatedBy
}

func (a *Task) GetID() int {
	return a.Id
}

var MapPriority map[int]string = map[int]string{
	1: "High",
	2: "Medium",
	3: "Low",
}

var MapTaskStatus map[int]string = map[int]string{
	3: "Deleted",
	2: "Completed",
	1: "Incomplete",
}
