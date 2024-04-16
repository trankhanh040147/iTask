package model

import (
	"time"
)

type ProjectMemberCreation struct {
	ProjectId int        `json:"project_id" gorm:"column:project_id"`
	UserId    int        `json:"user_id" gorm:"column:user_id"`
	AddedAt   *time.Time `json:"added_at" gorm:"column:added_at"`
	Role      int        `json:"role" gorm:"column:role"`
}

func (ProjectMemberCreation) TableName() string {
	return "ProjectMembers"
}

func (a *ProjectMemberCreation) GetUserID() int {
	return a.UserId
}

func (a *ProjectMemberCreation) GetProjectID() int {
	return a.ProjectId
}
