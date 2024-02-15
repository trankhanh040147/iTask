package entities

import "iTask/common"

type User struct {
	common.SQLModel
	AccountId int `json:"account_id" gorm:"column:account_id"`
}