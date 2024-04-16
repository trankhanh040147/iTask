package model

import (
	"iTask/modules/account/iomodel"
)

type SimpleMember struct {
	ProjectId   int                    `json:"project_id" gorm:"column:project_id;"`
	UserId      int                    `json:"-" gorm:"column:user_id;"`
	AccountInfo *iomodel.SimpleAccount `json:"account_info" gorm:"foreignKey:UserId"`
}

func (SimpleMember) TableName() string {
	return "ProjectMembers"
}
