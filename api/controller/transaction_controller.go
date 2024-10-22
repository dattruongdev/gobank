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
			Message: "Invalid body content",
			Status:  400,
		})
	}

	amount, err := strconv.ParseFloat(jsonBody["Amount"], 64)
	if err != nil {
		return c.JSON(http.StatusBadRequest, &Response{
			Message: "Amount is not a number",
			Status:  400,
		})
	}
	payeeId := jsonBody["PayeeId"]
	payerId := jsonBody["PayerId"]

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
	log.Println(transaction)

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
