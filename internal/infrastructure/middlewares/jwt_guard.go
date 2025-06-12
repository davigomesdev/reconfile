package middlewares

import (
	"strings"

	"github.com/davigomesdev/reconfile/internal/domain/errors"
	jwt_auth "github.com/davigomesdev/reconfile/internal/infrastructure/jwt-auth"
	"github.com/gin-gonic/gin"
)

func JWTGuard(jwtAuthService *jwt_auth.JWTAuthService) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		authHeader := ctx.GetHeader("Authorization")
		if authHeader == "" || !strings.HasPrefix(authHeader, "Bearer ") {
			appErr := errors.NewUnauthorizedExceptionError()
			ctx.AbortWithStatusJSON(appErr.Code, gin.H{
				"error": appErr,
			})
			return
		}

		tokenString := strings.TrimPrefix(authHeader, "Bearer ")

		claims, err := jwtAuthService.VerifyToken(tokenString)
		if err != nil {
			appErr := errors.NewUnauthorizedExceptionError()
			ctx.AbortWithStatusJSON(appErr.Code, gin.H{
				"error": appErr,
			})
			return
		}

		ctx.Set("currentUser", claims)
		ctx.Next()
	}
}
