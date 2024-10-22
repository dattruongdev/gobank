package repository

import (
	"context"
	"errors"
	"log"

	"github.com/d1nnn/domain"
	"github.com/google/uuid"
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
	log.Println("transaction repo: ", transaction)

	currentUser := domain.AppUser{
		ID: transaction.UserID,
	}
	targetUser := domain.AppUser{
		ID: transaction.PayeeID,
	}
	tx := tr.db.Find(&currentUser)
	if tx.Error != nil {
		return tx.Error
	}
	log.Println("current user: ", currentUser)

	tx = tr.db.Find(&targetUser)
	if tx.Error != nil {
		return tx.Error
	}

	log.Println("target user: ", targetUser)
	if currentUser.Balance < transaction.Amount {

		var ErrInsufficientBalance = errors.New("current user doesn't have enough balance")
		return ErrInsufficientBalance
	}

	currentUser.Balance -= transaction.Amount
	targetUser.Balance += transaction.Amount

	tx = tr.db.Save(&currentUser)
	if tx.Error != nil {
		return tx.Error
	}
	tx = tr.db.Save(&targetUser)
	if tx.Error != nil {
		return tx.Error
	}

	transaction.ID = uuid.New().String()

	tx = tr.db.Save(&transaction)

	return tx.Error
}
