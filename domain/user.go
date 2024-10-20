package domain

import (
	"context"
)

type AppUser struct {
	ID       string `json:"id"`
	Email    string `json:"email"`
	FullName string `json:"fullname"`
	Transactions []Transaction `gorm:"foreignKey:UserID" json:"transactions"`
}

type UserRepository interface {
	Create(c context.Context, user AppUser) error
	Update(c context.Context, updateRequest AppUser) error
	GetAll(c context.Context, userId string) ([]AppUser, error)
	GetByEmail(c context.Context, email string) (AppUser, error)
	GetByName(c context.Context, name string) ([]AppUser, error)
}