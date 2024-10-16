package domain

import (
	"context"

	"gorm.io/gorm"
)

type AppUser struct {
	gorm.Model
	ID       string `json:"id"`
	Email    string `json:"email"`
	FullName string `json:"email"`
}

type UserRepository interface {
	Create(c context.Context, user AppUser) error
	Update(c context.Context, updateRequest AppUser) error
	GetAll(c context.Context) ([]AppUser, error)
	GetByEmail(c context.Context, email string) (AppUser, error)
	GetByName(c context.Context, name string) ([]AppUser, error)
}