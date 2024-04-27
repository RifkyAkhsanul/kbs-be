package controllers

import (
	"net/http"
	"strconv"

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

func GetInfofromID(c echo.Context) error {
	id := c.FormValue("id")
	idInt, err := strconv.Atoi(id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	result, err := models.GetInfofromID(idInt)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(http.StatusAccepted, result)
}
