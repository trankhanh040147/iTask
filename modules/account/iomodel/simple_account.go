package iomodel

import "iTask/common"

type SimpleAccount struct {
	common.SQLModel
	FullName string `json:"full_name" gorm:"column:full_name"`
	Status   int    `json:"status" gorm:"column:status;"`
	Avatar   string `json:"profile_ava_url" gorm:"profile_ava_url"`
}

func (SimpleAccount) TableName() string {
	return "Users"
}
