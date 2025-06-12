package handlers

import (
	"net/http"

	"github.com/davigomesdev/reconfile/internal/adapters/httpx"
	"github.com/davigomesdev/reconfile/internal/adapters/presenters"
	authDTO "github.com/davigomesdev/reconfile/internal/application/dtos/auth"
	authUC "github.com/davigomesdev/reconfile/internal/application/usecases/auth"
	"github.com/gin-gonic/gin"
)

type AuthHandler struct {
	signInUseCase        *authUC.SignInUseCase
	signUpUseCase        *authUC.SignUpUseCase
	refreshTokensUseCase *authUC.RefreshTokensUseCase
}

func NewAuthHandler(signInUseCase *authUC.SignInUseCase, signUpUseCase *authUC.SignUpUseCase, refreshTokensUseCase *authUC.RefreshTokensUseCase) *AuthHandler {
	return &AuthHandler{
		signInUseCase:        signInUseCase,
		signUpUseCase:        signUpUseCase,
		refreshTokensUseCase: refreshTokensUseCase,
	}
}

func (h *AuthHandler) SignIn(c *gin.Context) {
	input, ok := httpx.BodyValidate[authDTO.SignInDTO](c)
	if !ok {
		return
	}

	tokens, err := h.signInUseCase.Execute(c.Request.Context(), *input)
	if err != nil {
		c.Error(err)
		return
	}

	c.JSON(http.StatusOK, presenters.NewAuthPresenter(tokens))
}

func (h *AuthHandler) SignUp(c *gin.Context) {
	input, ok := httpx.BodyValidate[authDTO.SignUpDTO](c)
	if !ok {
		return
	}

	tokens, err := h.signUpUseCase.Execute(c.Request.Context(), *input)
	if err != nil {
		c.Error(err)
		return
	}

	c.JSON(http.StatusOK, presenters.NewAuthPresenter(tokens))
}

func (h *AuthHandler) RefreshTokens(c *gin.Context) {
	input, ok := httpx.BodyValidate[authDTO.RefreshTokensDTO](c)
	if !ok {
		return
	}

	tokens, err := h.refreshTokensUseCase.Execute(c.Request.Context(), *input)
	if err != nil {
		c.Error(err)
		return
	}

	c.JSON(http.StatusOK, presenters.NewAuthPresenter(tokens))
}
