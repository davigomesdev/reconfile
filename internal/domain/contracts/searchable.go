package contracts

type SortDirection string

const (
	SortAsc  SortDirection = "asc"
	SortDesc SortDirection = "desc"
)

type SearchInput struct {
	Page    *int
	PerPage *int
	Sort    *string
	SortDir *SortDirection
	Filter  *string
}

type SearchResult[E any] struct {
	Items       []E
	Total       int
	CurrentPage int
	PerPage     int
	Sort        *string
	SortDir     *SortDirection
	Filter      *string
}

func (sr SearchResult[E]) LastPage() int {
	if sr.PerPage == 0 {
		return 1
	}
	pages := sr.Total / sr.PerPage
	if sr.Total%sr.PerPage != 0 {
		pages++
	}
	if pages == 0 {
		pages = 1
	}
	return pages
}
