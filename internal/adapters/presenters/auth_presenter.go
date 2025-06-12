package presenters

import (
	"github.com/davigomesdev/reconfile/internal/adapters/collections"
	jwt_auth "github.com/davigomesdev/reconfile/internal/infrastructure/jwt-auth"
)

func NewAuthPresenter(a *jwt_auth.AuthTokens) *collections.Presenter[*jwt_auth.AuthTokens] {
	return &collections.Presenter[*jwt_auth.AuthTokens]{
		Data: a,
	}

}
