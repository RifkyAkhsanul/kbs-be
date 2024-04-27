package controllers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/RifkyAkhsanul/kbs-be/models"
	"github.com/labstack/echo/v4"
)

type GetInfoRequest struct {
	ID int `json:"id"`
}

func GetInfo(c echo.Context) error {
	result, err := models.GetInfo()

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, result)
}

func GetInfofromID(c echo.Context) error {
	var req GetInfoRequest
	decoder := json.NewDecoder(c.Request().Body)
	err := decoder.Decode(&req)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": "Invalid request body: " + err.Error()})
	}

	id := req.ID
	idInt, err := strconv.Atoi(strconv.Itoa(id)) // Convert to int for consistency
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": "Error converting ID to integer: " + err.Error()})
	}

	result, err := models.GetInfofromID(idInt)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(http.StatusAccepted, result)
}
