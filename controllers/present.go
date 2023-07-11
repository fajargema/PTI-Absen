package controllers

import (
	"absen/models"
	"absen/services"
	"net/http"
	"strings"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

type PresentController struct {
	service services.PresentService
}

func InitPresentController() PresentController {
	return PresentController{
		service: services.InitPresentService(),
	}
}

func (pc *PresentController) GetHomeWidget(c echo.Context) error {
	token := c.Request().Header.Get("Authorization")
	if token == "" {
		return c.JSON(http.StatusBadRequest, models.Response[string]{
			Status:  "failed",
			Message: "Missing token in request header",
		})
	}
	token = strings.ReplaceAll(token, "Bearer ", "")

	widget, err := pc.service.GetHomeWidget(token)

	if err != nil {
		return c.JSON(http.StatusNotFound, models.Response[string]{
			Status:  "failed",
			Message: "home widget not found",
		})
	}

	return c.JSON(http.StatusOK, models.Response[models.HomeWidget]{
		Status:  "success",
		Message: "home widget found",
		Data:    widget,
	})
}

func (pc *PresentController) GetAll(c echo.Context) error {
	token := c.Request().Header.Get("Authorization")
	if token == "" {
		return c.JSON(http.StatusBadRequest, models.Response[string]{
			Status:  "failed",
			Message: "Missing token in request header",
		})
	}
	token = strings.ReplaceAll(token, "Bearer ", "")

	period := c.QueryParam("period")

	presents, err := pc.service.GetAll(token, period)

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

func (pc *PresentController) GetByID(c echo.Context) error {
	token := c.Request().Header.Get("Authorization")
	if token == "" {
		return c.JSON(http.StatusBadRequest, models.Response[string]{
			Status:  "failed",
			Message: "Missing token in request header",
		})
	}
	token = strings.ReplaceAll(token, "Bearer ", "")

	var presentID string = c.Param("id")

	present, err := pc.service.GetByID(presentID, token)

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

func (pc *PresentController) Search(c echo.Context) error {
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

	presents, err := pc.service.Search(date, token)

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

func (pc *PresentController) Create(c echo.Context) error {
	token := c.Request().Header.Get("Authorization")
	if token == "" {
		return c.JSON(http.StatusBadRequest, models.Response[string]{
			Status:  "failed",
			Message: "Missing token in request header",
		})
	}
	token = strings.ReplaceAll(token, "Bearer ", "")

	form, err := c.MultipartForm()
	if err != nil {
		return c.JSON(http.StatusBadRequest, models.Response[string]{
			Status:  "failed",
			Message: err.Error(),
		})
	}
	files := form.File["url"]

	var presentInput models.PresentInput
	if err := c.Bind(&presentInput); err != nil {
		return c.JSON(http.StatusBadRequest, models.Response[string]{
			Status:  "failed",
			Message: err.Error(),
		})
	}

	validate := validator.New()
	if err := validate.Struct(presentInput); err != nil {
		return c.JSON(http.StatusBadRequest, models.Response[string]{
			Status:  "failed",
			Message: err.Error(),
		})
	}

	present, err := pc.service.Create(presentInput, token, files)

	if err != nil {
		return c.JSON(http.StatusBadRequest, models.Response[string]{
			Status:  "failed",
			Message: err.Error(),
		})
	}

	return c.JSON(http.StatusCreated, models.Response[models.Present]{
		Status:  "success",
		Message: "present created",
		Data:    present,
	})
}
