package common

import (
	"database/sql/driver"
	"encoding/json"
	"time"
)

type SQLModel struct {
	Id int `json:"id" gorm:"column:id"`
	// FakeId    string    `json:"id" gorm:"-"`
	CreatedAt *time.Time `json:"created_at" gorm:"column:created_at"`
	UpdatedAt *time.Time `json:"updated_at" gorm:"column:updated_at"`
}

type Image struct {
	// Width  float32 `json:"width" gorm:"column:width"`
	// Height float32 `json:"height" gorm:"column:height"`
	Url string `json:"url" gorm:"column:url"`
}

func (j *Image) Value() (driver.Value, error) {
	if j == nil {
		return nil, nil
	}
	return json.Marshal(j)
}
