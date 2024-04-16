package model

import "strings"

type TaskCommentUpdate struct {
	Message string `json:"message" gorm:"column:message"`
}

func (TaskCommentUpdate) TableName() string {
	return TaskComment{}.TableName()
}

func (a *TaskCommentUpdate) Validate() error {
	a.Message = strings.TrimSpace(a.Message)
	if a.Message == "" {
		return ErrMessageCannotBeEmpty
	}

	return nil
}
