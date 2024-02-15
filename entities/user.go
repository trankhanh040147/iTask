package entities

import "paradise-booking/common"

type User struct {
	common.SQLModel
	AccountId int `json:"account_id" gorm:"column:account_id"`
}