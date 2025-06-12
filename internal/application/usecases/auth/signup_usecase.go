package auth

import (
	"context"

	authDTO "github.com/davigomesdev/reconfile/internal/application/dtos/auth"
	"github.com/davigomesdev/reconfile/internal/application/providers"
	"github.com/davigomesdev/reconfile/internal/domain/entities"
	"github.com/davigomesdev/reconfile/internal/domain/repositories"
	jwt_auth "github.com/davigomesdev/reconfile/internal/infrastructure/jwt-auth"
)

type SignUpUseCase struct {
	userRepository repositories.UserRepositoryInterface
	jwtAuthService *jwt_auth.JWTAuthService
	hashProvider   *providers.HashProvider
}

func NewSignUpUseCase(
	userRepository repositories.UserRepositoryInterface,
	jwtAuthService *jwt_auth.JWTAuthService,
	hashProvider *providers.HashProvider,
) *SignUpUseCase {
	return &SignUpUseCase{
		userRepository: userRepository,
		jwtAuthService: jwtAuthService,
		hashProvider:   hashProvider,
	}
}

func (uc *SignUpUseCase) Execute(ctx context.Context, input authDTO.SignUpDTO) (*jwt_auth.AuthTokens, error) {
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

	tokens, err := uc.jwtAuthService.GenerateTokens(entity.ID, entity.Name)
	if err != nil {
		return nil, err
	}

	entity.UpdateRefreshToken(&tokens.RefreshToken)

	if err = uc.userRepository.Create(ctx, entity); err != nil {
		return nil, err
	}

	return tokens, nil
}
