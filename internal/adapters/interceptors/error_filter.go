package interceptors

import (
	"github.com/davigomesdev/reconfile/internal/domain/errors"
	"github.com/gin-gonic/gin"
)

func ErrorFilter() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.Next()

		if len(ctx.Errors) == 0 {
			return
		}

		err := ctx.Errors.Last().Err

		if appErr, ok := err.(*errors.AppError); ok {
			ctx.AbortWithStatusJSON(appErr.Code, gin.H{
				"error": appErr,
			})
			return
		}

		appErr := errors.NewInternalError()
		ctx.AbortWithStatusJSON(appErr.Code, gin.H{
			"error": appErr,
		})
	}
}
