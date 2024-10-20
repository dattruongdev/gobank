package repository

import (
	"context"

	"github.com/d1nnn/domain"
	"gorm.io/gorm"
)

type PostgresTransactionRepository struct {
	db *gorm.DB
}

func NewPostgresTransactionRepository(db *gorm.DB) domain.TransactionRepository {
	return &PostgresTransactionRepository{
		db: db,
	}
}

func (tr *PostgresTransactionRepository) GetAllByUserId(c context.Context, userId string) ([]domain.Transaction, error) {
	var transactions []domain.Transaction
	tx := tr.db.Preload("Payee").Where("user_id = ?", userId).Find(&transactions)

	return transactions, tx.Error
}
func (tr *PostgresTransactionRepository) Create(c context.Context, transaction domain.Transaction) error {
	tx := tr.db.Save(&transaction)

	return tx.Error
}
