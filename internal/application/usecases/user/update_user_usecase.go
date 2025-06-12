package user

import (
	"context"

	userDTO "github.com/davigomesdev/reconfile/internal/application/dtos/user"
	"github.com/davigomesdev/reconfile/internal/domain/entities"
	"github.com/davigomesdev/reconfile/internal/domain/repositories"
)

type UpdateUserUseCase struct {
	userRepository repositories.UserRepositoryInterface
}

func NewUpdateUserUseCase(userRepository repositories.UserRepositoryInterface) *UpdateUserUseCase {
	return &UpdateUserUseCase{userRepository: userRepository}
}

func (uc *UpdateUserUseCase) Execute(ctx context.Context, input userDTO.UpdateUserDTO) (*entities.UserEntity, error) {
	user, err := uc.userRepository.Get(ctx, input.ID)
	if err != nil {
		return nil, err
	}

	if user.Email != input.Email {
		if err = uc.userRepository.EmailExists(ctx, input.Email); err != nil {
			return nil, err
		}
	}

	if err = user.Update(entities.UserProps{
		Name:  input.Name,
		Email: input.Email,
	}); err != nil {
		return nil, err
	}

	if err = uc.userRepository.Update(ctx, user); err != nil {
		return nil, err
	}

	return user, nil
}
