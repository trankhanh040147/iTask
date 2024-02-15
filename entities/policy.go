package entities

import "iTask/common"

type Policy struct {
	common.SQLModel
	PlaceId       int    `json:"place_id" gorm:"column:place_id"`
	Name          string `json:"name" gorm:"column:name"`
	GroupPolicyId int    `json:"group_policy_id" gorm:"column:group_policy_id"`
}

func (Policy) TableName() string {
	return "policies"
}
