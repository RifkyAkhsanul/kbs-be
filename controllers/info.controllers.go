package controllers

import (
	"net/http"

	"github.com/RifkyAkhsanul/kbs-be/models"
	"github.com/labstack/echo/v4"
)

func GetInfo(c echo.Context) error {
	result, err := models.GetInfo()

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, result)
}
