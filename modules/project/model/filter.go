package model

type Filter struct {
	Keyword       *string `json:"keyword"`
	DateRangeFrom int     `json:"date_range_from"`
}
