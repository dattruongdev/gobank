package usecase

import (
	"context"

	"github.com/d1nnn/domain"
)

type SignupUsecase struct {
	userRepository domain.UserRepository
}

type SignupRequest struct {
	UserID string
	Email string
	FullName string
}

func NewSignupUsecase(repo domain.UserRepository) *SignupUsecase {

	return &SignupUsecase {
		userRepository: repo,
	}
}

func (su *SignupUsecase) CreateUser(c context.Context, request SignupRequest) error {
	user := domain.AppUser {
		Email: request.Email,
		FullName: request.FullName,
		ID: request.UserID,
	}

	err := su.userRepository.Create(c,user)

	return err
}
