package admin

import (
	"absen/models"
	services "absen/services/admin"
	"net/http"
	"strings"
	"time"

	"github.com/labstack/echo/v4"
)

type AdminPresentController struct {
	service services.AdminPresentService
}

func InitAdminPresentController() AdminPresentController {
	return AdminPresentController{
		service: services.InitAdminPresentService(),
	}
}

func (apc *AdminPresentController) GetAll(c echo.Context) error {
	token := c.Request().Header.Get("Authorization")
	if token == "" {
		return c.JSON(http.StatusBadRequest, models.Response[string]{
			Status:  "failed",
			Message: "Missing token in request header",
		})
	}
	token = strings.ReplaceAll(token, "Bearer ", "")

	period := c.QueryParam("period")

	presents, err := apc.service.GetAll(token, period)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, models.Response[string]{
			Status:  "failed",
			Message: "failed to fetch presents data",
		})
	}

	return c.JSON(http.StatusOK, models.Response[[]models.Present]{
		Status:  "success",
		Message: "all presents",
		Data:    presents,
	})
}

func (apc *AdminPresentController) GetByID(c echo.Context) error {
	token := c.Request().Header.Get("Authorization")
	if token == "" {
		return c.JSON(http.StatusBadRequest, models.Response[string]{
			Status:  "failed",
			Message: "Missing token in request header",
		})
	}
	token = strings.ReplaceAll(token, "Bearer ", "")

	var presentID string = c.Param("id")

	present, err := apc.service.GetByID(presentID, token)

	if err != nil {
		return c.JSON(http.StatusNotFound, models.Response[string]{
			Status:  "failed",
			Message: "present not found",
		})
	}

	return c.JSON(http.StatusOK, models.Response[models.Present]{
		Status:  "success",
		Message: "present found",
		Data:    present,
	})
}

func (apc *AdminPresentController) Search(c echo.Context) error {
	token := c.Request().Header.Get("Authorization")
	if token == "" {
		return c.JSON(http.StatusBadRequest, models.Response[string]{
			Status:  "failed",
			Message: "Missing token in request header",
		})
	}
	token = strings.ReplaceAll(token, "Bearer ", "")

	date, err := time.Parse(time.DateOnly, c.QueryParam("date"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, models.Response[string]{
			Status:  "failed",
			Message: "invalid date",
		})
	}

	presents, err := apc.service.Search(date, token)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, models.Response[string]{
			Status:  "failed",
			Message: "failed to fetch presents data",
		})
	}

	return c.JSON(http.StatusOK, models.Response[[]models.Present]{
		Status:  "success",
		Message: "all presents",
		Data:    presents,
	})
}
