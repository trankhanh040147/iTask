package model

import "time"

type ProjectUpdate struct {
	// use pointer(*) to allow update data to "", 0, false... except nil
	Name        *string    `json:"name" gorm:"column:name"`
	Description *string    `json:"description" gorm:"column:description"`
	Status      *int       `json:"status" gorm:"column:status"`
	Thumbnail   *string    `json:"thumbnail_url" gorm:"column:thumbnail_url"`
	Privacy     *string    `json:"privacy" gorm:"column:privacy"`
	Deadline    *string    `json:"deadline" gorm:"column:deadline"`
	StartedAt   *time.Time `json:"started_at" gorm:"column:started_at"`
}

func (ProjectUpdate) TableName() string { return Project{}.TableName() }
