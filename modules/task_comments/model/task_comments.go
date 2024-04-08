package model

import (
	"errors"
	"iTask/common"
)

var (
	ErrCreatedByCannotBeEmpty = errors.New("Created by cannot be empty")
	ErrMessageCannotBeEmpty   = errors.New("Message cannot be empty")
)

const (
	EntityName = "TaskComments"
)

type TaskComments struct {
	common.SQLModel
	TaskId          int    `json:"task_id" gorm:"column:task_id"`
	CreatedBy       int    `json:"created_by" gorm:"column:created_by"`
	ParentCommentId int    `json:"parent_comment_id" gorm:"column:parent_comment_id"`
	IsPinned        bool   `json:"is_pinned" gorm:"column:pinned"`
	Message         string `json:"message" gorm:"column:message"`
}

func (TaskComments) TableName() string { return "TaskComments" }

func (a *TaskComments) GetCreatedBy() int { return a.CreatedBy }
