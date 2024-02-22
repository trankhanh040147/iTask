package model

type Filter struct {
	Keyword          string `json:"keyword" form:"keyword"`
	CreatedDateRange int    `json:"created_date_range" form:"created_date_range"`
}
