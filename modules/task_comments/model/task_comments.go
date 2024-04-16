package model

import (
	"iTask/common"
	"iTask/modules/account/iomodel"
)

const (
	EntityName = "TaskComments"
)

type TaskComment struct {
	common.SQLModel
	TaskId          int                    `json:"task_id" gorm:"column:task_id"`
	CreatedBy       int                    `json:"-" gorm:"column:created_by"`
	ParentCommentId int                    `json:"parent_comment_id" gorm:"column:parent_comment_id"`
	IsPinned        bool                   `json:"is_pinned" gorm:"column:pinned"`
	Message         string                 `json:"message" gorm:"column:message"`
	Owner           *iomodel.SimpleAccount `json:"owner" gorm:"foreignKey:CreatedBy"`
}

func (TaskComment) TableName() string { return "TaskComments" }

func (a *TaskComment) GetCreatedBy() int { return a.CreatedBy }
