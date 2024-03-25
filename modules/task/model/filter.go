package model

type Filter struct {
	Keyword string `json:"keyword" form:"keyword"`
	Status  int    `json:"status" form:"status"`
}
