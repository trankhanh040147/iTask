package model

import (
	"iTask/common"
	"time"
)

type TaskCreation struct {
	common.SQLModel
	Name        string       `json:"name" gorm:"column:name"`
	Description string       `json:"description" gorm:"column:description"`
	Status      TaskStatus   `json:"status" gorm:"column:status"`
	ProjectId   int          `json:"project_id" gorm:"column:project_id"`
	CreatedBy   int          `json:"created_by" gorm:"column:created_by"`
	ParentTask  int          `json:"parent_task_id" gorm:"column:parent_task_id"`
	Position    float64      `json:"position" gorm:"column:position"`
	Priority    TaskPriority `json:"priority" gorm:"column:priority"`
	DueDate     *time.Time   `json:"due_date" gorm:"column:due_date"`
	StartedAt   *time.Time   `json:"started_at" gorm:"column:started_at"`
	CompletedAt *time.Time   `json:"completed_at" gorm:"column:completed_at"`
}

func (TaskCreation) TableName() string {
	return Task{}.TableName()
}

func (t *TaskCreation) GetStatus() string {
	return t.Status.String()
}

func (t *TaskCreation) GetCreatedBy() int {
	return t.CreatedBy
}

func (t *TaskCreation) GetID() int {
	return t.Id
}
