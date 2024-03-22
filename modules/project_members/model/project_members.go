package model

import (
	"errors"
	"time"
)

var (
	ErrNameCannotBeEmpty = errors.New("name cannot be empty")
	ErrProjectIsDeleted  = errors.New("project is deleted")
)

const (
	EntityName = "ProjectMember"
)

type ProjectMember struct {
	Id        int        `json:"id" gorm:"column:id"`
	ProjectId int        `json:"project_id" gorm:"column:project_id"`
	UserId    int        `json:"user_id" gorm:"column:user_id"`
	AddedAt   *time.Time `json:"added_at" gorm:"column:added_at"`
	Role      int        `json:"role" gorm:"column:role"`
}

func (ProjectMember) TableName() string {
	return "ProjectMembers"
}

func (a *ProjectMember) GetUserID() int {
	return a.UserId
}

func (a *ProjectMember) GetProjectID() int {
	return a.ProjectId
}

var MapProjectStatus map[int]string = map[int]string{
	3: "Deleted",
	2: "Completed",
	1: "Incomplete",
}
