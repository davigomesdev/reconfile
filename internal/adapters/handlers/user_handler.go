package handlers

import (
	"net/http"

	"github.com/davigomesdev/reconfile/internal/adapters/httpx"
	"github.com/davigomesdev/reconfile/internal/adapters/presenters"
	userDTO "github.com/davigomesdev/reconfile/internal/application/dtos/user"
	userUC "github.com/davigomesdev/reconfile/internal/application/usecases/user"
	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	getUserUseCase            *userUC.GetUserUseCase
	searchUserUseCase         *userUC.SearchUserUseCase
	createUserUseCase         *userUC.CreateUserUseCase
	updateUserUseCase         *userUC.UpdateUserUseCase
	updatePasswordUserUseCase *userUC.UpdatePasswordUserUseCase
	deleteUserUseCase         *userUC.DeleteUserUseCase
}

func NewUserHandler(
	getUserUseCase *userUC.GetUserUseCase,
	searchUserUseCase *userUC.SearchUserUseCase,
	createUserUseCase *userUC.CreateUserUseCase,
	updateUserUseCase *userUC.UpdateUserUseCase,
	updatePasswordUserUseCase *userUC.UpdatePasswordUserUseCase,
	deleteUserUseCase *userUC.DeleteUserUseCase,
) *UserHandler {
	return &UserHandler{
		getUserUseCase:            getUserUseCase,
		searchUserUseCase:         searchUserUseCase,
		updateUserUseCase:         updateUserUseCase,
		updatePasswordUserUseCase: updatePasswordUserUseCase,
		deleteUserUseCase:         deleteUserUseCase,
	}
}

func (h *UserHandler) Current(c *gin.Context) {
	currentUser, err := httpx.CurrentUser(c)
	if err != nil {
		c.Error(err)
		return
	}

	user, err := h.getUserUseCase.Execute(
		c.Request.Context(),
		userDTO.GetUserDTO{ID: currentUser.ID},
	)
	if err != nil {
		c.Error(err)
		return
	}

	c.JSON(http.StatusOK, presenters.NewUserPresenter(user))
}

func (h *UserHandler) Get(c *gin.Context) {
	input, ok := httpx.ParamValidate[userDTO.GetUserDTO](c)
	if !ok {
		return
	}

	user, err := h.getUserUseCase.Execute(c.Request.Context(), *input)
	if err != nil {
		c.Error(err)
		return
	}

	c.JSON(http.StatusOK, presenters.NewUserPresenter(user))
}

func (h *UserHandler) Search(c *gin.Context) {
	input, ok := httpx.QueryValidate[userDTO.SearchUserDTO](c)
	if !ok {
		return
	}

	searchResult, err := h.searchUserUseCase.Execute(c.Request.Context(), *input)
	if err != nil {
		c.Error(err)
		return
	}

	c.JSON(http.StatusOK, presenters.NewUserCollectionPresenter(searchResult))
}

func (h *UserHandler) Create(c *gin.Context) {
	input, ok := httpx.BodyValidate[userDTO.CreateUserDTO](c)
	if !ok {
		return
	}

	user, err := h.createUserUseCase.Execute(c.Request.Context(), *input)
	if err != nil {
		c.Error(err)
		return
	}

	c.JSON(http.StatusOK, presenters.NewUserPresenter(user))
}

func (h *UserHandler) Update(c *gin.Context) {
	idDTO, ok := httpx.ParamValidate[userDTO.GetUserDTO](c)
	if !ok {
		return
	}

	input, ok := httpx.BodyValidate[userDTO.UpdateUserDTO](c)
	if !ok {
		return
	}

	input.ID = idDTO.ID

	user, err := h.updateUserUseCase.Execute(c.Request.Context(), *input)
	if err != nil {
		c.Error(err)
		return
	}

	c.JSON(http.StatusOK, presenters.NewUserPresenter(user))
}

func (h *UserHandler) UpdateCurrent(c *gin.Context) {
	currentUser, err := httpx.CurrentUser(c)
	if err != nil {
		c.Error(err)
		return
	}

	input, ok := httpx.BodyValidate[userDTO.UpdateUserDTO](c)
	if !ok {
		return
	}

	input.ID = currentUser.ID

	user, err := h.updateUserUseCase.Execute(c.Request.Context(), *input)
	if err != nil {
		c.Error(err)
		return
	}

	c.JSON(http.StatusOK, presenters.NewUserPresenter(user))
}

func (h *UserHandler) UpdatePassword(c *gin.Context) {
	currentUser, err := httpx.CurrentUser(c)
	if err != nil {
		c.Error(err)
		return
	}

	input, ok := httpx.BodyValidate[userDTO.UpdatePasswordUserDTO](c)
	if !ok {
		return
	}

	input.ID = currentUser.ID

	if err := h.updatePasswordUserUseCase.Execute(c.Request.Context(), *input); err != nil {
		c.Error(err)
		return
	}

	c.JSON(http.StatusNoContent, nil)
}

func (h *UserHandler) Delete(c *gin.Context) {
	input, ok := httpx.ParamValidate[userDTO.DeleteUserDTO](c)
	if !ok {
		return
	}

	if err := h.deleteUserUseCase.Execute(c.Request.Context(), *input); err != nil {
		c.Error(err)
		return
	}

	c.JSON(http.StatusNoContent, nil)
}
