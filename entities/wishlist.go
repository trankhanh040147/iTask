package entities

import "paradise-booking/common"

type WishList struct {
	common.SQLModel
	UserID int    `gorm:"column:user_id"`
	Title  string `gorm:"column:title"`
}

func (WishList) TableName() string {
	return "wishlists"
}
