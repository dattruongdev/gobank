package controller

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/d1nnn/domain"
	"github.com/d1nnn/usecase"
	"github.com/labstack/echo/v4"
)

type TransactionController struct {
	transactionUsecase *usecase.TransactionUsecase
}

func NewTransactionController(uc *usecase.TransactionUsecase) *TransactionController {
	return &TransactionController{
		transactionUsecase: uc,
	}
}

func (tc *TransactionController) CreateTransaction(c echo.Context) error {

	jsonBody := make(map[string]string)
	log.Println("hello")

	err := json.NewDecoder(c.Request().Body).Decode(&jsonBody)
	log.Println("hi")

	if err != nil {
		return c.JSON(http.StatusBadRequest, &Response{
			Message: err.Error(),
			Status:  400,
		})
	}

	amount, err := strconv.ParseFloat(jsonBody["amount"], 64)
	if err != nil {
		return c.JSON(http.StatusBadRequest, &Response{
			Message: "Amount is not a number",
			Status:  400,
		})
	}
	payeeId := jsonBody["payeeId"]
	payerId := jsonBody["payerId"]
	// status := jsonBody["status"]

	if payeeId == payerId {
		return c.JSON(http.StatusBadRequest, &Response{
			Message: "Can't send money to your self",
			Status:  400,
		})

	}

	transaction := domain.Transaction{
		Amount:  amount,
		PayeeID: payeeId,
		UserID:  payerId,
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

func (tc *TransactionController) GetPendingTransactions(c echo.Context) error {
	userId := c.Param("userId")

	transactions, err := tc.transactionUsecase.GetPendingTransactions(c.Request().Context(), userId)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, &Response{
			Message: err.Error(),
			Status:  500,
		})
	}

	return c.JSON(http.StatusOK, transactions)
}

func (tc *TransactionController) CancelTransactions(c echo.Context) error {
	jsonBody := make(map[string]interface{})

	err := json.NewDecoder(c.Request().Body).Decode(&jsonBody)

	if err != nil {
		return c.JSON(http.StatusBadRequest, &Response{
			Message: err.Error(),
			Status:  400,
		})
	}

	var txIds []string

	if ids, ok := jsonBody["txIds"].([]interface{}); ok {
		for _, id := range ids {
			if txId, ok := id.(string); ok {
				txIds = append(txIds, txId)
			} else {
				log.Println("Invalid txId:", id) // log the invalid ID
			}
		}
	}

	err = tc.transactionUsecase.CancelTransactions(c.Request().Context(), txIds...)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, &Response{
			Message: err.Error(),
			Status:  500,
		})
	}

	return c.JSON(http.StatusOK, &Response{
		Message: "Deleted transaction(s)",
		Status:  200,
	})
}

func (tc *TransactionController) ApproveTransactions(c echo.Context) error {
	jsonBody := make(map[string]interface{})

	err := json.NewDecoder(c.Request().Body).Decode(&jsonBody)

	if err != nil {
		return c.JSON(http.StatusBadRequest, &Response{
			Message: err.Error(),
			Status:  400,
		})
	}

	var txIds []string

	if ids, ok := jsonBody["txIds"].([]interface{}); ok {
		for _, id := range ids {
			if txId, ok := id.(string); ok {
				txIds = append(txIds, txId)
			} else {
				log.Println("Invalid txId:", id) // log the invalid ID
			}
		}
	}

	err = tc.transactionUsecase.ApproveTransactions(c.Request().Context(), txIds...)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, &Response{
			Message: err.Error(),
			Status:  500,
		})
	}

	return c.JSON(http.StatusOK, &Response{
		Message: "Deleted transaction(s)",
		Status:  200,
	})
}
