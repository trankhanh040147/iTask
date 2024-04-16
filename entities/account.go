package entities

import (
	"iTask/common"
)

const (
	EntityName = "Accounts"
)

type Account struct {
	common.SQLModel
	Username        string     `json:"username" gorm:"column:username"`
	Email           string     `json:"email" gorm:"column:email"`
	FullName        string     `json:"full_name" gorm:"column:full_name"`
	Role            UserRole   `json:"role" gorm:"role"`
	Title           string     `json:"title" gorm:"column:title"`
	Status          UserStatus `json:"status" gorm:"column:status"`
	Password        string     `json:"password" gorm:"column:password"`
	Address         string     `json:"address" gorm:"column:address"`
	Phone           string     `json:"phone" gorm:"column:phone"`
	Dob             string     `json:"dob" gorm:"column:dob"`
	Avatar          string     `json:"profile_ava_url" gorm:"column:avatar"`
	Cover           string     `json:"profile_cover_url" gorm:"column:profile_cover_url"`
	IsEmailVerified int        `json:"is_email_verified" gorm:"column:is_email_verified"`
	Bio             string     `json:"bio" gorm:"column:bio"`
	IsInvited       bool       `json:"is_invited"`
}

func (Account) TableName() string {
	return "Users"
}

func (a *Account) GetRole() string {
	return a.Role.String()
}

func (a *Account) GetEmail() string {
	return a.Email
}

func (a *Account) GetUserId() int {
	return a.Id
}

type UserRole int

const (
	RoleMember UserRole = 1 + iota
	RoleOwner
	RoleAdmin
	RoleObserver
)

func (role UserRole) String() string {
	switch role {
	case RoleAdmin:
		return "admin"
	case RoleMember:
		return "member"
	case RoleOwner:
		return "owner"
	default:
		return "observer"
	}
}

type UserStatus int

const (
	StatusActive UserStatus = 1 + iota
	StatusInactive
)

func (status UserStatus) String() string {
	switch status {
	case StatusActive:
		return "active"
	default:
		return "inactive"
	}
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

//var (
//	ErrEmailOrPasswordInvalid = common.NewCustomError(
//		errors.New("email or password invalid"),
//		"email or password invalid",
//		"ErrUsernameOrPasswordInvalid",
//	)
//
//	ErrEmailExisted = common.NewCustomError(
//		errors.New("email has already existed"),
//		"email has already existed",
//		"ErrEmailExisted",
//	)
//)
