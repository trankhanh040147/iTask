package iomodel

type SimpleAccount struct {
	//common.SQLModel
	Id       int    `json:"id" gorm:"column:id"`
	FullName string `json:"full_name" gorm:"column:full_name"`
	Status   int    `json:"-" gorm:"column:status"`
	Avatar   string `json:"profile_ava_url" gorm:"profile_ava_url"`
	Title    string `json:"title" gorm:"title"`
}

func (SimpleAccount) TableName() string {
	return "Users"
}
