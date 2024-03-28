package model

import (
	"iTask/common"
	"time"
)

type TaskUpdate struct {
	common.SQLModel
	Name        string     `json:"name" gorm:"column:name"`
	Description string     `json:"description" gorm:"column:description"`
	Status      TaskStatus `json:"status" gorm:"column:status"`
	ProjectId   int        `json:"-" gorm:"column:project_id"`
	//CreatedBy   int                      `json:"-" gorm:"column:created_by"`
	//ParentTask  int                      `json:"-" gorm:"column:parent_task_id"`
	Position    float64    `json:"position" gorm:"column:position"`
	Priority    int        `json:"priority" gorm:"column:priority"`
	Completed   bool       `json:"completed" gorm:"column:completed"`
	DueDate     *time.Time `json:"due_date" gorm:"column:due_date"`
	StartedAt   *time.Time `json:"started_at" gorm:"column:started_at"`
	CompletedAt *time.Time `json:"completed_at" gorm:"column:completed_at"`
	//Owner       *userModel.SimpleAccount `json:"-" gorm:"foreignKey:CreatedBy"`
}

func (TaskUpdate) TableName() string {
	return Task{}.TableName()
}

func (a *TaskUpdate) GetStatus() string {
	return a.Status.String()
}

func (a *TaskUpdate) GetID() int {
	return a.Id
}
