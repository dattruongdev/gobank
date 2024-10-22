package usecase

import (
	"context"

	"github.com/d1nnn/domain"
)

type UserUsecase struct {
	userRepository domain.UserRepository
}

func NewUserUsecase(repo domain.UserRepository) *UserUsecase {
	return &UserUsecase{
		userRepository: repo,
	}
}

func (uu *UserUsecase) GetUsers(c context.Context, userId string) ([]domain.AppUser, error) {
	return uu.userRepository.GetAll(c, userId)
}

func (uu *UserUsecase) GetByEmail(c context.Context, email string) (domain.AppUser, error) {
	return uu.userRepository.GetByEmail(c, email)
}
func (uu *UserUsecase) GetById(c context.Context, userId string) (domain.AppUser, error) {
	return uu.userRepository.GetById(c, userId)
}
