package presenters

import (
	"time"

	"github.com/davigomesdev/reconfile/internal/adapters/collections"
	"github.com/davigomesdev/reconfile/internal/domain/contracts"
	"github.com/davigomesdev/reconfile/internal/domain/entities"
)

type UserOutput struct {
	ID        string    `json:"id"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

func newUserOutput(u *entities.UserEntity) *UserOutput {
	return &UserOutput{
		ID:        u.ID,
		Name:      u.Name,
		Email:     u.Email,
		CreatedAt: u.CreatedAt,
		UpdatedAt: u.UpdatedAt,
	}
}

func NewUserPresenter(u *entities.UserEntity) *collections.Presenter[*UserOutput] {
	return &collections.Presenter[*UserOutput]{
		Data: newUserOutput(u),
	}

}

func NewUserCollectionPresenter(sr *contracts.SearchResult[*entities.UserEntity]) *collections.CollectionPresenter[*UserOutput] {
	data := make([]*UserOutput, 0, len(sr.Items))

	for _, user := range sr.Items {
		if p := newUserOutput(user); p != nil {
			data = append(data, p)
		}
	}

	return &collections.CollectionPresenter[*UserOutput]{
		Data: data,
		Meta: collections.PaginationPresenter{
			CurrentPage: sr.CurrentPage,
			PerPage:     sr.PerPage,
			LastPage:    sr.LastPage(),
			Total:       sr.Total,
		},
	}
}
