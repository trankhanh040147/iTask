package model

type Filter struct {
	Keyword string `json:"keyword" form:"keyword"`
	TagType string `json:"tag_type" form:"tag_type"`
}
