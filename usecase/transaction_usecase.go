package usecase

import (
	"context"

	"github.com/d1nnn/domain"
)

type TransactionUsecase struct {
	transactionRepository domain.TransactionRepository
}

func NewTransactionUsecase(repo domain.TransactionRepository) *TransactionUsecase {
	return &TransactionUsecase{
		transactionRepository: repo,
	}
}

func (tu *TransactionUsecase) GetAllFromUser(c context.Context, userId string) ([]domain.Transaction, error) {
	return tu.transactionRepository.GetAllByUserId(c, userId)
}

func (tu *TransactionUsecase) Create(c context.Context, transaction domain.Transaction) error {
	return tu.transactionRepository.Create(c, transaction)
}

func (tu *TransactionUsecase) GetPendingTransactions(c context.Context, userId string) ([]domain.Transaction, error) {
	return tu.transactionRepository.GetPendings(c, userId)
}

func (tu *TransactionUsecase) CancelTransactions(c context.Context, txIds ...string) error {
	return tu.transactionRepository.DeleteTransactions(c, txIds...)
}
func (tu *TransactionUsecase) ApproveTransactions(c context.Context, txIds ...string) error {
	return tu.transactionRepository.ApproveTransactions(c, txIds...)
}
