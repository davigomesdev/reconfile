package entities

import (
	"time"

	"github.com/davigomesdev/reconfile/internal/domain/contracts"
	"github.com/davigomesdev/reconfile/internal/domain/validators"
)

type UserProps struct {
	ID           *string
	Name         string
	Email        string
	Password     string
	RefreshToken *string
	CreatedAt    *time.Time
	UpdatedAt    *time.Time
	DeletedAt    *time.Time
}

type UserEntity struct {
	contracts.Entity
	Name         string  `json:"name" validate:"required,min=3,max=255"`
	Email        string  `json:"email" validate:"required,email"`
	Password     string  `json:"password" validate:"required,min=6,max=255"`
	RefreshToken *string `json:"refreshToken" validate:"omitempty,max=255"`
}

func NewUserEntity(props UserProps) (*UserEntity, error) {
	entity := &UserEntity{
		Name:         props.Name,
		Email:        props.Email,
		Password:     props.Password,
		RefreshToken: props.RefreshToken,
	}

	entity.Init(props.ID, props.CreatedAt, props.UpdatedAt, props.DeletedAt)

	if err := validators.ValidatorFields(entity); err != nil {
		return nil, err
	}

	return entity, nil
}

func (u *UserEntity) Update(props UserProps) error {
	u.Name = props.Name
	u.Email = props.Email
	u.UpdatedAt = time.Now()

	return validators.ValidatorFields(u)

}

func (u *UserEntity) UpdatePassword(password string) error {
	u.Password = password
	u.UpdatedAt = time.Now()

	return validators.ValidatorFields(u)
}

func (u *UserEntity) UpdateRefreshToken(token *string) error {
	u.RefreshToken = token
	u.UpdatedAt = time.Now()

	return validators.ValidatorFields(u)
}
