package user

import (
	"context"

	userDTO "github.com/davigomesdev/reconfile/internal/application/dtos/user"
	"github.com/davigomesdev/reconfile/internal/domain/repositories"
)

type DeleteUserUseCase struct {
	userRepository repositories.UserRepositoryInterface
}

func NewDeleteUserUseCase(userRepository repositories.UserRepositoryInterface) *DeleteUserUseCase {
	return &DeleteUserUseCase{userRepository: userRepository}
}

func (uc *DeleteUserUseCase) Execute(ctx context.Context, input userDTO.DeleteUserDTO) error {
	_, err := uc.userRepository.Get(ctx, input.ID)
	if err != nil {
		return err
	}

	if err := uc.userRepository.Delete(ctx, input.ID); err != nil {
		return err
	}

	return nil
}
