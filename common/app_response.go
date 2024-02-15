package common

type successRes struct {
	Data   any `json:"data"`
	Paging any `json:"paging,omitempty"`
	Filter any `json:"filter,omitempty"`
}

func NewSuccessResponse(data, paging, filter any) *successRes {
	return &successRes{Data: data, Paging: paging, Filter: filter}
}

func SimpleSuccessResponse(data any) *successRes {
	return NewSuccessResponse(data, nil, nil)
}
