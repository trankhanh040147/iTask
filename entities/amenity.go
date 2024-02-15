package entities

import "iTask/common"

type Amenity struct {
	common.SQLModel
	PlaceId         int     `json:"place_id" gorm:"column:place_id"`
	Description     *string `json:"description" gorm:"column:description"`
	ConfigAmenityId int     `json:"config_amenity_id" gorm:"column:config_amenity_id"`
}

func (Amenity) TableName() string {
	return "amenities"
}

type ConfigAmenity struct {
	common.SQLModel
	Icon string `json:"icon" gorm:"column:icon"`
	Name string `json:"name" gorm:"column:name"`
}

func (ConfigAmenity) TableName() string {
	return "config_amenity"
}
