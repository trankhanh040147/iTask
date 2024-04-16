package model

import (
	"errors"
	"strings"
)

var (
	ErrCreatedByCannotBeEmpty = errors.New("Created by cannot be empty")
	ErrMessageCannotBeEmpty   = errors.New("Message cannot be empty")
)

type TaskCommentCreation struct {
	Id              int    `json:"id" gorm:"column:id"`
	TaskId          int    `json:"task_id" gorm:"column:task_id"`
	CreatedBy       int    `json:"created_by" gorm:"column:created_by"`
	ParentCommentId int    `json:"parent_comment_id" gorm:"column:parent_comment_id"`
	Message         string `json:"message" gorm:"column:message"`
}

func (TaskCommentCreation) TableName() string { return TaskComment{}.TableName() }

func (a *TaskCommentCreation) Validate() error {
	a.Message = strings.TrimSpace(a.Message)

	if a.Message == "" {
		return ErrMessageCannotBeEmpty
	}

	if a.CreatedBy == 0 {
		return ErrCreatedByCannotBeEmpty
	}

	return nil
}
