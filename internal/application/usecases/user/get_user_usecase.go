package user

import (
	"context"

	userDTO "github.com/davigomesdev/reconfile/internal/application/dtos/user"
	"github.com/davigomesdev/reconfile/internal/domain/entities"
	"github.com/davigomesdev/reconfile/internal/domain/repositories"
)

type GetUserUseCase struct {
	userRepository repositories.UserRepositoryInterface
}

func NewGetUserUseCase(userRepository repositories.UserRepositoryInterface) *GetUserUseCase {
	return &GetUserUseCase{userRepository: userRepository}
}

func (uc *GetUserUseCase) Execute(ctx context.Context, input userDTO.GetUserDTO) (*entities.UserEntity, error) {
	return uc.userRepository.Get(ctx, input.ID)
}
