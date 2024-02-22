package model

import (
	"iTask/common"
	"strings"
)

type ProjectCreation struct {
	Id          int           `json:"id" gorm:"column:id;"`
	UserId      int           `json:"-" gorm:"column:user_id;"`
	Title       string        `json:"title" gorm:"column:title;"`
	Description string        `json:"description" gorm:"column:;description"`
	Image       *common.Image `json:"image" gorm:"column:image;"`
}

func (i *ProjectCreation) Validate() error {
	i.Title = strings.TrimSpace(i.Title)

	if i.Title == "" {
		return ErrNameCannotBeEmpty
	}

	return nil
}

func (ProjectCreation) TableName() string { return Project{}.TableName() }
