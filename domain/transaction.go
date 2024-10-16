package domain

import (
	"context"

	"gorm.io/gorm"
)

type Transaction struct {
	gorm.Model
	ID string `json:"transaction_id"`
	Amount float64 `json:"amount"`
	Payee string `json:"payee"`
	Date string `json:"date"`
	UserID string `json:"user_id"`
}

type TransactionRepository interface {
	GetAllByUserId(c context.Context, userId string) ([]Transaction, error)
	Create(c context.Context, tx Transaction) error
}