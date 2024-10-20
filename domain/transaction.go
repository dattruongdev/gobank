package domain

import (
	"context"
	"time"
)

type Transaction struct {
	ID         string    `json:"transaction_id"`
	Amount     float64   `json:"amount"`
	Date       time.Time `json:"date"`
	UserID     string    `json:"user_id"`
	PayeeID    string    `json:"-"`
	CategoryID string    `json:"-"`
	Payee      AppUser   `json:"payee" gorm:"foreignKey:PayeeID"`
	Category   Category  `json:"category" gorm:"foreignKey:CategoryID"`
}

type TransactionRepository interface {
	GetAllByUserId(c context.Context, userId string) ([]Transaction, error)
	Create(c context.Context, tx Transaction) error
}
