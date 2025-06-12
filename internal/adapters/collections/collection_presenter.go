package collections

type Presenter[T any] struct {
	Data T `json:"data"`
}

type CollectionPresenter[T any] struct {
	Data []T                 `json:"data"`
	Meta PaginationPresenter `json:"meta"`
}
