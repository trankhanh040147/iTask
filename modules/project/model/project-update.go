package model

type ProjectUpdate struct {
	// use pointer(*) to allow update data to "", 0, false... except nil
	Title       *string `json:"title" gorm:"column:title;"`
	Description *string `json:"description" gorm:"column:;description"`
	Status      *string `json:"status" gorm:"column:status;"`
}

func (ProjectUpdate) TableName() string { return Project{}.TableName() }
