package controller

import (
	"net/http"

	"github.com/d1nnn/domain"
	"github.com/d1nnn/usecase"
	"github.com/labstack/echo/v4"
)

type TransactionController struct {
	transactionUsecase *usecase.TransactionUsecase
}

func NewTransactionController(uc *usecase.TransactionUsecase) *TransactionController {
	return &TransactionController {
		transactionUsecase: uc,
	}
}

func(tc *TransactionController) CreateTransaction(c echo.Context) error {
	var transaction domain.Transaction
	err := c.Bind(&transaction); if err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}

	err = tc.transactionUsecase.Create(c.Request().Context(), transaction)
	if err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}

	return err
}

func (tc *TransactionController) GetByUserId(c echo.Context) error {
	id := c.Param("userid")
	transactions, err := tc.transactionUsecase.GetAllFromUser(c.Request().Context(), id)
	if err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, transactions)
}
