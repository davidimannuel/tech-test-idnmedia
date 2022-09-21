package usecases

import "time"

const (
	defaultPerPage = 10
)

type BaseEntity struct {
	CreatedAt time.Time
	CreatedBy string
	UpdatedAt time.Time
	UpdatedBy string
}

type Pagination struct {
	CurrentPage int
	PerPage     int
	TotalData   int
}

func (p *Pagination) Validate() (res int) {
	if p.CurrentPage <= 0 {
		p.CurrentPage = 1
	}

	if p.PerPage <= 0 {
		p.PerPage = defaultPerPage
	}
	return
}

func (p Pagination) GetLastPage() (res int) {

	res = p.TotalData / p.PerPage
	if (p.TotalData % p.PerPage) != 0 {
		res += 1
	}
	return
}

func (p Pagination) GetOffset() int {

	return (p.CurrentPage - 1) * p.PerPage
}
