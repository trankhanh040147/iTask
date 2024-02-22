package model

type Filter struct {
	Keyword         string `json:"keyword" form:"keyword"`
	CreatedDayRange int    `json:"created_day_range" form:"created_day_range"`
}
