package httpx

import (
	"github.com/davigomesdev/reconfile/internal/domain/errors"
	jwt_auth "github.com/davigomesdev/reconfile/internal/infrastructure/jwt-auth"
	"github.com/gin-gonic/gin"
)

func CurrentUser(ctx *gin.Context) (*jwt_auth.JWTClaims, error) {
	if user, exists := ctx.Get("currentUser"); exists {
		if claims, ok := user.(*jwt_auth.JWTClaims); ok {
			return claims, nil
		}
	}
	return nil, errors.NewUnauthorizedExceptionError()
}
