package user

import (
	"context"

	userDTO "github.com/davigomesdev/reconfile/internal/application/dtos/user"
	"github.com/davigomesdev/reconfile/internal/application/providers"
	"github.com/davigomesdev/reconfile/internal/domain/errors"
	"github.com/davigomesdev/reconfile/internal/domain/repositories"
)

type UpdatePasswordUserUseCase struct {
	userRepository repositories.UserRepositoryInterface
	hashProvider   *providers.HashProvider
}

func NewUpdatePasswordUserUseCase(
	userRepository repositories.UserRepositoryInterface,
	hashProvider *providers.HashProvider,
) *UpdatePasswordUserUseCase {
	return &UpdatePasswordUserUseCase{
		userRepository: userRepository,
		hashProvider:   hashProvider,
	}
}

func (uc *UpdatePasswordUserUseCase) Execute(ctx context.Context, input userDTO.UpdatePasswordUserDTO) error {
	user, err := uc.userRepository.Get(ctx, input.ID)
	if err != nil {
		return err
	}

	isValid := uc.hashProvider.CompareHash(input.OldPassword, user.Password)
	if !isValid {
		return errors.NewInvalidCredentialsError()
	}

	hash, err := uc.hashProvider.GenerateHash(input.NewPassword)
	if err != nil {
		return err
	}

	if err = user.UpdatePassword(hash); err != nil {
		return err
	}

	if err = uc.userRepository.Update(ctx, user); err != nil {
		return err
	}

	return nil
}
