package domain

import (
	"context"
)

type AppUser struct {
	ID           string        `json:"id" gorm:"primaryKey"`
	Email        string        `json:"email"`
	FullName     string        `json:"fullname"`
	Balance      float64       `json:"balance" gorm:"default:0"`
	Transactions []Transaction `gorm:"foreignKey:UserID" json:"transactions"`
}

type UserRepository interface {
	Create(c context.Context, user AppUser) error
	Update(c context.Context, updateRequest AppUser) error
	GetAll(c context.Context, userId string) ([]AppUser, error)
	GetByEmail(c context.Context, email string) (AppUser, error)
	GetByName(c context.Context, name string) ([]AppUser, error)
	GetById(c context.Context, userId string) (AppUser, error)
}
