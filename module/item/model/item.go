package model

import (
	"errors"
	"social-todo-list/common"
	"strings"
)

var (
	ErrTitleCannotBeEmpty = errors.New("title cannot be empty")
	ErrItemIsDeleted      = errors.New("item is deleted")
)

const (
	EntityName = "Item"
)

type TodoItem struct {
	common.SQLModel                    // embed struct
	UserId          int                `json:"-" gorm:"column:user_id;"`
	Title           string             `json:"title" gorm:"column:title;"`
	Description     string             `json:"description" gorm:"column:;description"`
	Status          string             `json:"status" gorm:"column:status;"`
	Image           *common.Image      `json:"image" gorm:"column:image;"`
	LikedCount      int                `json:"liked_count" gorm:"liked_count"`
	Owner           *common.SimpleUser `json:"owner" gorm:"foreignKey:UserId;"`
}

// >> Why it do not have receiver like (t TodoItem) ? --> it apply for all TodoItem objects
func (TodoItem) TableName() string { return "todo_items" }

// fea_FakeID

func (i *TodoItem) Mask() {
	i.SQLModel.Mask(common.DbTypeItem)

	if v := i.Owner; v != nil {
		v.Mask()
	}
}

type TodoItemCreation struct {
	Id          int           `json:"id" gorm:"column:id;"`
	UserId      int           `json:"-" gorm:"column:user_id;"`
	Title       string        `json:"title" gorm:"column:title;"`
	Description string        `json:"description" gorm:"column:;description"`
	Image       *common.Image `json:"image" gorm:"column:image;"`
}

func (i *TodoItemCreation) Validate() error {
	i.Title = strings.TrimSpace(i.Title)

	if i.Title == "" {
		return ErrTitleCannotBeEmpty
	}

	return nil
}

func (TodoItemCreation) TableName() string { return TodoItem{}.TableName() }

type TodoItemUpdate struct {
	// use pointer(*) to allow update data to "", 0, false... except nil
	Title       *string `json:"title" gorm:"column:title;"`
	Description *string `json:"description" gorm:"column:;description"`
	Status      *string `json:"status" gorm:"column:status;"`
}

func (TodoItemUpdate) TableName() string { return TodoItem{}.TableName() }
