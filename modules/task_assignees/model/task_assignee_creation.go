package model

import (
	"errors"
)

var (
	ErrTaskIdCannotBeEmpty = errors.New("task id cannot be empty")
	ErrUserIdCannotBeEmpty = errors.New("user id cannot be empty")
)

type TaskAssigneeCreation struct {
	Id     int `json:"id" gorm:"column:id"`
	TaskId int `json:"task_id" gorm:"column:task_id"`
	UserId int `json:"user_id" gorm:"column:user_id"`
}

func (TaskAssigneeCreation) TableName() string {
	return TaskAssignee{}.TableName()
}

func (TaskAssigneeCreation) GetEntityName() string {
	return EntityName
}

func (a *TaskAssigneeCreation) GetId() int {
	return a.Id
}

func (a *TaskAssigneeCreation) Validate() error {
	if a.TaskId == 0 {
		return ErrTaskIdCannotBeEmpty
	}
	if a.UserId == 0 {
		return ErrUserIdCannotBeEmpty
	}
	return nil
}
