package auth

import (
	"context"

	authDTO "github.com/davigomesdev/reconfile/internal/application/dtos/auth"
	"github.com/davigomesdev/reconfile/internal/application/providers"
	"github.com/davigomesdev/reconfile/internal/domain/errors"
	"github.com/davigomesdev/reconfile/internal/domain/repositories"
	jwt_auth "github.com/davigomesdev/reconfile/internal/infrastructure/jwt-auth"
)

type SignInUseCase struct {
	userRepository repositories.UserRepositoryInterface
	jwtAuthService *jwt_auth.JWTAuthService
	hashProvider   *providers.HashProvider
}

func NewSignInUseCase(
	userRepository repositories.UserRepositoryInterface,
	jwtAuthService *jwt_auth.JWTAuthService,
	hashProvider *providers.HashProvider,
) *SignInUseCase {
	return &SignInUseCase{
		userRepository: userRepository,
		jwtAuthService: jwtAuthService,
		hashProvider:   hashProvider,
	}
}

func (uc *SignInUseCase) Execute(ctx context.Context, input authDTO.SignInDTO) (*jwt_auth.AuthTokens, error) {
	user, err := uc.userRepository.GetByEmail(ctx, input.Email)
	if err != nil {
		return nil, err
	}

	isValid := uc.hashProvider.CompareHash(input.Password, user.Password)
	if !isValid {
		return nil, errors.NewInvalidCredentialsError()
	}

	tokens, err := uc.jwtAuthService.GenerateTokens(user.ID, user.Name)
	if err != nil {
		return nil, err
	}

	user.UpdateRefreshToken(&tokens.RefreshToken)

	if err = uc.userRepository.Update(ctx, user); err != nil {
		return nil, err
	}

	return tokens, nil
}
