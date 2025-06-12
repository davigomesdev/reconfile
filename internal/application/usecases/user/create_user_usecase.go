package user

import (
	"context"

	userDTO "github.com/davigomesdev/reconfile/internal/application/dtos/user"
	"github.com/davigomesdev/reconfile/internal/application/providers"
	"github.com/davigomesdev/reconfile/internal/domain/entities"
	"github.com/davigomesdev/reconfile/internal/domain/repositories"
)

type CreateUserUseCase struct {
	userRepository repositories.UserRepositoryInterface
	hashProvider   *providers.HashProvider
}

func NewCreateUserUseCase(
	userRepository repositories.UserRepositoryInterface,
	hashProvider *providers.HashProvider,
) *CreateUserUseCase {
	return &CreateUserUseCase{
		userRepository: userRepository,
		hashProvider:   hashProvider,
	}
}

func (uc *CreateUserUseCase) Execute(ctx context.Context, input userDTO.CreateUserDTO) (*entities.UserEntity, error) {
	if err := uc.userRepository.EmailExists(ctx, input.Email); err != nil {
		return nil, err
	}

	hash, err := uc.hashProvider.GenerateHash(input.Password)
	if err != nil {
		return nil, err
	}

	entity, err := entities.NewUserEntity(entities.UserProps{
		Name:     input.Name,
		Email:    input.Email,
		Password: hash,
	})
	if err != nil {
		return nil, err
	}

	if err = uc.userRepository.Create(ctx, entity); err != nil {
		return nil, err
	}

	return entity, nil
}
