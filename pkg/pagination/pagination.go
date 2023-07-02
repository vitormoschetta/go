package pagination

type Pagination struct {
	PageSize    int `json:"page_size"`
	CurrentPage int `json:"current_page"`
	Total       int `json:"total"`
	LastPage    int `json:"last_page"`
}

func (p *Pagination) GetOffset() int {
	return (p.CurrentPage - 1) * p.PageSize
}

func (p *Pagination) BuildLastPage() {
	if p.Total == 0 {
		p.LastPage = 0
	}
	p.LastPage = (p.Total + p.PageSize - 1) / p.PageSize
}
