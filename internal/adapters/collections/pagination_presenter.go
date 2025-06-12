package collections

type PaginationPresenter struct {
	CurrentPage int `json:"currentPage"`
	PerPage     int `json:"perPage"`
	LastPage    int `json:"lastPage"`
	Total       int `json:"total"`
}
