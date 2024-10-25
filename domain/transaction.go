package domain

import (
	"context"
	"time"
)

type Transaction struct {
	ID          string    `json:"id" gorm:"primaryKey"`
	Amount      float64   `json:"amount" gorm:"default:0.0"`
	Status      string    `json:"status" gorm:"default:PENDING"`
	Date        time.Time `json:"date" gorm:"default:current_timestamp"`
	UserID      string    `json:"user_id" gorm:"not null"`
	CurrentUser AppUser   `json:"current_user" gorm:"foreignKey:UserID"`
	PayeeID     string    `json:"-" gorm:"not null"`
	Payee       AppUser   `json:"payee" gorm:"foreignKey:PayeeID"`
}

type TransactionRepository interface {
	GetAllByUserId(c context.Context, userId string) ([]Transaction, error)
	Create(c context.Context, tx Transaction) error
	GetPendings(c context.Context, userId string) ([]Transaction, error)
	ApproveTransactions(c context.Context, txIds ...string) error
	DeleteTransactions(c context.Context, txIds ...string) error
}
