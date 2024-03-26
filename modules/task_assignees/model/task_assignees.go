package model

import (
	"iTask/modules/account/iomodel"
	"time"
)

const (
	EntityName = "task_assignees"
)

type TaskAssignee struct {
	Id           int                    `json:"id" gorm:"column:id"`
	TaskId       int                    `json:"task_id" gorm:"column:task_id"`
	UserId       int                    `json:"-" gorm:"column:user_id"`
	AssignedDate *time.Time             `json:"assigned_date" gorm:"column:assigned_date"`
	UserInfo     *iomodel.SimpleAccount `json:"user_info" gorm:"foreignKey:UserId"`
}

func (TaskAssignee) TableName() string {
	return "TaskAssigned"
}

func (TaskAssignee) GetEntityName() string {
	return EntityName
}

func (a *TaskAssignee) GetId() int {
	return a.Id
}
