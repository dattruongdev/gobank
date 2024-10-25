package controller

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/d1nnn/domain"
	"github.com/d1nnn/usecase"
	"github.com/labstack/echo/v4"
)

type PresetController struct {
	presetUsecase *usecase.PresetUsecase
}

func NewPresetController(pu *usecase.PresetUsecase) *PresetController {
	return &PresetController{
		presetUsecase: pu,
	}
}

func (pc *PresetController) GetAllFromUser(c echo.Context) error {
	id := c.Param("userid")

	presets, err := pc.presetUsecase.GetAll(c.Request().Context(), id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, &Response{
			Status:  500,
			Message: "Can't fetch preset from user",
		})
	}

	return c.JSON(http.StatusOK, presets)
}

func (pc *PresetController) DeletePreset(c echo.Context) error {
	log.Println("HELLOOOOOO")
	jsonBody := make(map[string]interface{})

	err := json.NewDecoder(c.Request().Body).Decode(&jsonBody)
	if err != nil {
		return c.JSON(http.StatusBadRequest, &Response{
			Message: err.Error(),
			Status:  400,
		})
	}

	payerId := jsonBody["payerId"].(string)
	var payeeIds []string
	if ids, ok := jsonBody["payeeIds"].([]interface{}); ok {
		for _, id := range ids {
			if str, ok := id.(string); ok {
				payeeIds = append(payeeIds, str)
			} else {
				// Handle the case where the type is not string, if necessary
			}
		}
	}

	err = pc.presetUsecase.Delete(c.Request().Context(), payerId, payeeIds...)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, &Response{
			Message: err.Error(),
			Status:  500,
		})
	}

	return c.JSON(http.StatusOK, &Response{
		Message: "Deleted successfully",
		Status:  200,
	})
}

func (pc *PresetController) CreatePreset(c echo.Context) error {
	jsonBody := make(map[string]interface{})

	err := json.NewDecoder(c.Request().Body).Decode(&jsonBody)
	if err != nil {
		return c.String(http.StatusBadRequest, "bad request")
	}
	payeeId := jsonBody["PayeeId"].(string)
	payerId := jsonBody["PayerId"].(string)
	// payeeName := jsonBody["PayeeName"].(string)

	preset := domain.Preset{
		PayeeID: payeeId,
		PayerID: payerId,
	}

	err = pc.presetUsecase.Create(c.Request().Context(), preset)
	if err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusCreated, &Response{
		Status:  201,
		Message: "Created preset",
	})
}
