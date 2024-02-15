package entities

import (
	"paradise-booking/common"
)

type Account struct {
	common.SQLModel
	Username        string `json:"username" gorm:"column:username"`
	Email           string `json:"email" gorm:"column:email"`
	FullName        string `json:"full_name" gorm:"column:full_name"`
	Role            int    `json:"role" gorm:"role"`
	Status          int    `json:"status" gorm:"column:status"`
	Password        string `json:"password" gorm:"column:password"`
	Address         string `json:"address" gorm:"column:address"`
	Phone           string `json:"phone" gorm:"column:phone"`
	Dob             string `json:"dob" gorm:"column:dob"`
	Avatar          string `json:"profile_ava_url" gorm:"profile_ava_url"`
	IsEmailVerified int    `json:"is_email_verified" gorm:"is_email_verified"`
	Bio             string `json:"bio" gorm:"bio"`
}

func (Account) TableName() string {
	return "Users"
}

func (a *Account) GetRole() int {
	return a.Role
}

func (a *Account) GetEmail() string {
	return a.Email
}

func (a *Account) GetID() int {
	return a.Id
}

var MapRole map[int]string = map[int]string{
	1: "Member",
	2: "Owner",
	3: "Admin",
	4: "Observer",
}

var MapStatus map[int]string = map[int]string{
	2: "Active",
	1: "Inactive",
}
