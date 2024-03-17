package model

import (
	"iTask/common"
	"iTask/modules/project_members/model"
	"strings"
	"time"
)

type ProjectCreation struct {
	common.SQLModel
	Name                string               `json:"name" gorm:"column:name"`
	Description         string               `json:"description" gorm:"column:description"`
	Status              int                  `json:"status" gorm:"column:status"`
	Thumbnail           string               `json:"thumbnail_url" gorm:"column:thumbnail_url"`
	Privacy             ProjectPrivacy       `json:"privacy" gorm:"column:privacy"`
	Priority            int                  `json:"priority" gorm:"column:priority"`
	CreatedBy           int                  `json:"created_by" gorm:"column:created_by"`
	Deadline            string               `json:"deadline" gorm:"column:deadline"`
	StartedAt           *time.Time           `json:"started_at" gorm:"column:started_at"`
	TotalTasks          int                  `json:"total_tasks" gorm:"-"`
	TotalCompletedTasks int                  `json:"completed_tasks" gorm:"-"`
	Members             []model.SimpleMember `json:"members" gorm:"foreignKey:ProjectId"`
	Tags                string               `json:"tags" gorm:"-"`
}

func (i *ProjectCreation) Validate() error {
	i.Name = strings.TrimSpace(i.Name)

	if i.Name == "" {
		return ErrNameCannotBeEmpty
	}

	return nil
}

func (ProjectCreation) TableName() string { return Project{}.TableName() }

//json body:
//{
//	"name": "project name",
//	"description": "project description",
//	"status": 1,
//	"thumbnail_url": "https://www.google.com",
//	"privacy": "public",
//	"created_by": 1,
//	"deadline": "2021-12-12",
//	"started_at": "2021-12-12"
//}
