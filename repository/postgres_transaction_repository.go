package repository

import (
	"context"
	"errors"
	"log"

	"github.com/d1nnn/domain"
	"github.com/google/uuid"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
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

	transaction.ID = uuid.New().String()

	tx := tr.db.Save(&transaction)

	return tx.Error
}

func (tr *PostgresTransactionRepository) GetPendings(c context.Context, userId string) ([]domain.Transaction, error) {
	var transactions []domain.Transaction
	tx := tr.db.Where("status = 'PENDING'").Find(&transactions)

	return transactions, tx.Error
}

func (tr *PostgresTransactionRepository) ApproveTransactions(c context.Context, txIds ...string) error {
	err := tr.db.Transaction(func(tx *gorm.DB) error {
		var transactions []domain.Transaction
		if err := tx.Where("id in (?)", txIds).Find(&transactions).Error; err != nil {
			return err
		}
		for _, transaction := range transactions {
			currentUser := domain.AppUser{
				ID: transaction.UserID,
			}
			targetUser := domain.AppUser{
				ID: transaction.PayeeID,
			}
			err := tx.Find(&currentUser).Error
			if err != nil {
				return tx.Rollback().Error
			}

			err = tx.Find(&targetUser).Error
			if err != nil {
				return tx.Rollback().Error
			}

			if currentUser.Balance < transaction.Amount {

				var ErrInsufficientBalance = errors.New("current user doesn't have enough balance")
				return ErrInsufficientBalance
			}

			currentUser.Balance -= transaction.Amount
			targetUser.Balance += transaction.Amount

			err = tx.Save(&currentUser).Error
			if err != nil {
				return tx.Rollback().Error
			}
			err = tx.Save(&targetUser).Error
			if err != nil {
				return tx.Rollback().Error
			}
		}

		err := tx.Clauses(clause.OnConflict{
			Columns:   []clause.Column{{Name: "id"}},
			DoUpdates: clause.Assignments(map[string]interface{}{"status": "COMPLETED"}),
		}).Create(&transactions).Error
		if err != nil {
			return tx.Rollback().Error
		}

		return nil
	})

	return err
}

func (tr *PostgresTransactionRepository) DeleteTransactions(c context.Context, txIds ...string) error {
	tx := tr.db.Where("id in (?)", txIds).Delete(&domain.Transaction{})

	return tx.Error
}
