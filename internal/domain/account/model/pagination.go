package model

// Paginator struct for holding pagination info
type Paginator struct {
	TotalRecord int64 `json:"total_record"`
	TotalPage   int64 `json:"total_page"`
	Offset      int64 `json:"offset"`
	Limit       int64 `json:"limit"`
	Page        int64 `json:"page"`
	PrevPage    int64 `json:"prev_page"`
	NextPage    int64 `json:"next_page"`
}

// PaginationData struct for returning pagination stat
type PaginationData struct {
	Total     int64 `json:"total"`
	Page      int64 `json:"page"`
	PerPage   int64 `json:"perPage"`
	Prev      int64 `json:"prev"`
	Next      int64 `json:"next"`
	TotalPage int64 `json:"totalPage"`
}

func (p *Paginator) PaginationData() *PaginationData {
	data := PaginationData{
		Total:     p.TotalRecord,
		Page:      p.Page,
		PerPage:   p.Limit,
		Prev:      0,
		Next:      0,
		TotalPage: p.TotalPage,
	}
	if p.Page != p.PrevPage && p.TotalRecord > 0 {
		data.Prev = p.PrevPage
	}
	if p.Page != p.NextPage && p.TotalRecord > 0 && p.Page <= p.TotalPage {
		data.Next = p.NextPage
	}

	return &data
}
