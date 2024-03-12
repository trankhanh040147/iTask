package model

import (
	"iTask/common"
	"strings"
)

type TagCreation struct {
	common.SQLModel
	Name string  `json:"name" gorm:"column:name"`
	Type TagType `json:"tag_type" gorm:"column:tag_type"`
}

func (TagCreation) TableName() string {
	return Tag{}.TableName()
}

func (t *TagCreation) Validate() error {
	t.Name = strings.TrimSpace(t.Name)

	if t.Name == "" {
		return ErrNameCannotBeEmpty
	}

	return nil
}
