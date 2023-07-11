package admin

import (
	m "absen/middleware"
	"absen/models"
	services "absen/services/admin"
	"net/http"
	"strings"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

type AdminUserController struct {
	service services.AdminUserService
}

func InitAdminUserController() AdminUserController {
	return AdminUserController{
		service: services.InitAdminUserService(),
	}
}

func (auc *AdminUserController) GetAll(c echo.Context) error {
	token := c.Request().Header.Get("Authorization")
	if token == "" {
		return c.JSON(http.StatusBadRequest, models.Response[string]{
			Status:  "failed",
			Message: "Missing token in request header",
		})
	}
	token = strings.ReplaceAll(token, "Bearer ", "")

	allowedRoles := []string{"isAdmin"}
	isAllowed, err := m.IsAllowedRole(token, allowedRoles)
	if err != nil {
		return c.JSON(http.StatusUnauthorized, models.Response[string]{
			Status:  "failed",
			Message: "Unauthorized",
		})
	}

	if !isAllowed {
		return c.JSON(http.StatusForbidden, models.Response[string]{
			Status:  "failed",
			Message: "Forbidden",
		})
	}

	users, err := auc.service.GetAll(token)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, models.Response[string]{
			Status:  "failed",
			Message: "failed to fetch users data",
		})
	}

	return c.JSON(http.StatusOK, models.Response[[]models.User]{
		Status:  "success",
		Message: "all users",
		Data:    users,
	})
}

func (auc *AdminUserController) GetByID(c echo.Context) error {
	token := c.Request().Header.Get("Authorization")
	if token == "" {
		return c.JSON(http.StatusBadRequest, models.Response[string]{
			Status:  "failed",
			Message: "Missing token in request header",
		})
	}
	token = strings.ReplaceAll(token, "Bearer ", "")

	allowedRoles := []string{"isAdmin"}
	isAllowed, err := m.IsAllowedRole(token, allowedRoles)
	if err != nil {
		return c.JSON(http.StatusUnauthorized, models.Response[string]{
			Status:  "failed",
			Message: "Unauthorized",
		})
	}

	if !isAllowed {
		return c.JSON(http.StatusForbidden, models.Response[string]{
			Status:  "failed",
			Message: "Forbidden",
		})
	}

	var userID string = c.Param("id")

	user, err := auc.service.GetByID(userID, token)

	if err != nil {
		return c.JSON(http.StatusNotFound, models.Response[string]{
			Status:  "failed",
			Message: "user not found",
		})
	}

	return c.JSON(http.StatusOK, models.Response[models.User]{
		Status:  "success",
		Message: "user found",
		Data:    user,
	})
}

func (auc *AdminUserController) Create(c echo.Context) error {
	token := c.Request().Header.Get("Authorization")
	if token == "" {
		return c.JSON(http.StatusBadRequest, models.Response[string]{
			Status:  "failed",
			Message: "Missing token in request header",
		})
	}
	token = strings.ReplaceAll(token, "Bearer ", "")

	allowedRoles := []string{"isAdmin"}
	isAllowed, err := m.IsAllowedRole(token, allowedRoles)
	if err != nil {
		return c.JSON(http.StatusUnauthorized, models.Response[string]{
			Status:  "failed",
			Message: "Unauthorized",
		})
	}

	if !isAllowed {
		return c.JSON(http.StatusForbidden, models.Response[string]{
			Status:  "failed",
			Message: "Forbidden",
		})
	}

	var userInput models.UserInput

	if err := c.Bind(&userInput); err != nil {
		return c.JSON(http.StatusBadRequest, models.Response[string]{
			Status:  "failed",
			Message: "invalid request",
		})
	}

	validate := validator.New()
	if err := validate.Struct(userInput); err != nil {
		return c.JSON(http.StatusBadRequest, models.Response[string]{
			Status:  "failed",
			Message: err.Error(),
		})
	}

	user, err := auc.service.Create(userInput, token)

	if err != nil {
		return c.JSON(http.StatusBadRequest, models.Response[string]{
			Status:  "failed",
			Message: err.Error(),
		})
	}

	return c.JSON(http.StatusCreated, models.Response[models.User]{
		Status:  "success",
		Message: "user created",
		Data:    user,
	})
}

func (auc *AdminUserController) Update(c echo.Context) error {
	token := c.Request().Header.Get("Authorization")
	if token == "" {
		return c.JSON(http.StatusBadRequest, models.Response[string]{
			Status:  "failed",
			Message: "Missing token in request header",
		})
	}
	token = strings.ReplaceAll(token, "Bearer ", "")

	allowedRoles := []string{"isAdmin"}
	isAllowed, err := m.IsAllowedRole(token, allowedRoles)
	if err != nil {
		return c.JSON(http.StatusUnauthorized, models.Response[string]{
			Status:  "failed",
			Message: "Unauthorized",
		})
	}

	if !isAllowed {
		return c.JSON(http.StatusForbidden, models.Response[string]{
			Status:  "failed",
			Message: "Forbidden",
		})
	}

	var userInput models.UserInput
	if err := c.Bind(&userInput); err != nil {
		return c.JSON(http.StatusBadRequest, models.Response[string]{
			Status:  "failed",
			Message: "invalid request",
		})
	}

	validate := validator.New()
	if err := validate.Struct(userInput); err != nil {
		return c.JSON(http.StatusBadRequest, models.Response[string]{
			Status:  "failed",
			Message: err.Error(),
		})
	}

	var userID string = c.Param("id")

	user, err := auc.service.Update(userInput, userID, token)
	if err != nil {
		return c.JSON(http.StatusBadRequest, models.Response[string]{
			Status:  "failed",
			Message: err.Error(),
		})
	}

	return c.JSON(http.StatusOK, models.Response[models.User]{
		Status:  "success",
		Message: "user updated",
		Data:    user,
	})
}

func (auc *AdminUserController) Delete(c echo.Context) error {
	token := c.Request().Header.Get("Authorization")
	if token == "" {
		return c.JSON(http.StatusBadRequest, models.Response[string]{
			Status:  "failed",
			Message: "Missing token in request header",
		})
	}
	token = strings.ReplaceAll(token, "Bearer ", "")

	allowedRoles := []string{"isAdmin"}
	isAllowed, err := m.IsAllowedRole(token, allowedRoles)
	if err != nil {
		return c.JSON(http.StatusUnauthorized, models.Response[string]{
			Status:  "failed",
			Message: "Unauthorized",
		})
	}

	if !isAllowed {
		return c.JSON(http.StatusForbidden, models.Response[string]{
			Status:  "failed",
			Message: "Forbidden",
		})
	}

	var userID string = c.Param("id")

	err = auc.service.Delete(userID, token)
	if err != nil {
		return c.JSON(http.StatusNotFound, models.Response[string]{
			Status:  "failed",
			Message: "Not Found",
		})
	}

	return c.JSON(http.StatusOK, models.Response[string]{
		Status:  "success",
		Message: "user deleted",
	})
}
