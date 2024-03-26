package model

type Filter struct {
	TaskID int `json:"task_id" gorm:"column:task_id"`
}
