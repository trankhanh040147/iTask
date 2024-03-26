package model

type SimpleProject struct {
	// use pointer(*) to allow update data to "", 0, false... except nil
	ID   *int    `json:"id" gorm:"column:id"`
	Name *string `json:"name" gorm:"column:name"`
	//Description *string    `json:"description" gorm:"column:description"`
	//Thumbnail   *string    `json:"thumbnail_url" gorm:"column:thumbnail_url"`
}

func (SimpleProject) TableName() string { return Project{}.TableName() }
