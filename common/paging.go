package common

type Paging struct {
	Page       int    `json:"page" form:"page"`
	Limit      int    `json:"limit" form:"limit"`
	Total      int64  `json:"total" form:"-"`
	FakeCursor string `json:"cursor,omitempty" form:"cursor"`
	NextCursor string `json:"next_cursor,omitempty"`
}

func (p *Paging) Process() {
	if p.Page < 1 {
		p.Page = 1
	}

	if p.Limit <= 0 {
		p.Limit = 10
	}

	if p.Limit >= 100 {
		p.Limit = 100
	}
}
