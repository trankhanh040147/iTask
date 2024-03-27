package model

import (
	"iTask/modules/account/iomodel"
)

type SimpleAssignee struct {
	TaskId   int                    `json:"-" gorm:"column:task_id"`
	UserId   int                    `json:"-" gorm:"column:user_id"`
	UserInfo *iomodel.SimpleAccount `json:"user_info" gorm:"foreignKey:UserId"`
}

func (SimpleAssignee) TableName() string {
	return TaskAssignee{}.TableName()
}

func (SimpleAssignee) GetEntityName() string {
	return EntityName
}
