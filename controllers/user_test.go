package controllers

import (
	"absen/config"
	"absen/middleware"
	"absen/models"
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"golang.org/x/crypto/bcrypt"
)

type testCaseUser struct {
	name                   string
	path                   string
	expectedStatus         int
	expectedBodyStartsWith string
}

var controller UserController = InitUserController()

func InitEcho() *echo.Echo {
	config.InitDB()

	e := echo.New()

	return e
}

func TestRegisterUser_Success(t *testing.T) {
	testcase := testCaseUser{
		name:                   "success",
		path:                   "/api/v1/users/register",
		expectedStatus:         http.StatusCreated,
		expectedBodyStartsWith: "{\"status\":",
	}

	e := InitEcho()

	password, _ := bcrypt.GenerateFromPassword([]byte("inisecret"), bcrypt.DefaultCost)

	var userInput models.UserInput = models.UserInput{
		Name:     "test",
		Email:    "test@gmail.com",
		Password: string(password),
	}

	jsonBody, err := json.Marshal(&userInput)

	if err != nil {
		t.Errorf("error: %v\n", err)
	}

	bodyReader := bytes.NewReader(jsonBody)

	request := httptest.NewRequest(http.MethodPost, testcase.path, bodyReader)

	recorder := httptest.NewRecorder()

	request.Header.Add("Content-Type", "application/json")

	ctx := e.NewContext(request, recorder)

	ctx.SetPath(testcase.path)

	if assert.NoError(t, controller.Register(ctx)) {
		assert.Equal(t, http.StatusCreated, testcase.expectedStatus)

		body := recorder.Body.String()

		assert.True(t, strings.HasPrefix(body, testcase.expectedBodyStartsWith))
	}
}

func TestRegisterUser_Failed(t *testing.T) {
	testcase := testCaseUser{
		name:                   "failed",
		path:                   "/api/v1/users/register",
		expectedStatus:         http.StatusBadRequest,
		expectedBodyStartsWith: "{\"status\":",
	}

	e := InitEcho()

	userInput := models.UserInput{}

	jsonBody, _ := json.Marshal(&userInput)
	bodyReader := bytes.NewReader(jsonBody)

	request := httptest.NewRequest(http.MethodPost, "/api/v1/users/register", bodyReader)
	recorder := httptest.NewRecorder()

	request.Header.Add("Content-Type", "application/json")

	ctx := e.NewContext(request, recorder)

	ctx.SetPath(testcase.path)

	if assert.NoError(t, controller.Register(ctx)) {
		assert.Equal(t, http.StatusBadRequest, recorder.Code)
		body := recorder.Body.String()

		assert.True(t, strings.HasPrefix(body, testcase.expectedBodyStartsWith))
	}
}

func TestGetUserByUsername_Success(t *testing.T) {
	testcase := testCaseUser{
		name:                   "success",
		path:                   "/api/v1/users",
		expectedStatus:         http.StatusOK,
		expectedBodyStartsWith: "{\"status\":",
	}

	e := InitEcho()

	user, err := config.SeedUser()

	if err != nil {
		t.Errorf("error: %v\n", err)
	}

	username := user.Username

	request := httptest.NewRequest(http.MethodGet, testcase.path, nil)

	recorder := httptest.NewRecorder()

	ctx := e.NewContext(request, recorder)

	ctx.SetPath(testcase.path)

	ctx.SetParamNames("username")
	ctx.SetParamValues(username)

	if assert.NoError(t, controller.GetByUsername(ctx)) {
		assert.Equal(t, http.StatusOK, testcase.expectedStatus)

		body := recorder.Body.String()

		assert.True(t, strings.HasPrefix(body, testcase.expectedBodyStartsWith))
	}
}

func TestLoginUser_Success(t *testing.T) {
	testcase := testCaseUser{
		name:                   "success",
		path:                   "/api/v1/users/login",
		expectedStatus:         http.StatusOK,
		expectedBodyStartsWith: "{\"status\":",
	}

	e := InitEcho()

	password, _ := bcrypt.GenerateFromPassword([]byte("inisecret"), bcrypt.DefaultCost)

	var userAuth models.UserAuth = models.UserAuth{
		Username: "usernametest",
		Password: string(password),
	}

	jsonBody, err := json.Marshal(&userAuth)

	if err != nil {
		t.Errorf("error: %v\n", err)
	}

	bodyReader := bytes.NewReader(jsonBody)

	request := httptest.NewRequest(http.MethodPost, testcase.path, bodyReader)

	recorder := httptest.NewRecorder()

	request.Header.Add("Content-Type", "application/json")

	ctx := e.NewContext(request, recorder)

	ctx.SetPath(testcase.path)

	if assert.NoError(t, controller.Login(ctx)) {
		assert.Equal(t, http.StatusOK, testcase.expectedStatus)

		body := recorder.Body.String()

		assert.True(t, strings.HasPrefix(body, testcase.expectedBodyStartsWith))
	}
}

func TestLoginUser_Failed(t *testing.T) {
	testcase := testCaseUser{
		name:                   "failed",
		path:                   "/api/v1/users/login",
		expectedStatus:         http.StatusBadRequest,
		expectedBodyStartsWith: "{\"status\":",
	}

	e := InitEcho()

	userAuth := models.UserAuth{}

	jsonBody, _ := json.Marshal(&userAuth)
	bodyReader := bytes.NewReader(jsonBody)

	request := httptest.NewRequest(http.MethodPost, "/api/v1/users/login", bodyReader)
	recorder := httptest.NewRecorder()

	request.Header.Add("Content-Type", "application/json")

	ctx := e.NewContext(request, recorder)

	ctx.SetPath(testcase.path)

	if assert.NoError(t, controller.Login(ctx)) {
		assert.Equal(t, http.StatusBadRequest, recorder.Code)
		body := recorder.Body.String()

		assert.True(t, strings.HasPrefix(body, testcase.expectedBodyStartsWith))
	}
}

func TestUpdateUser_Success(t *testing.T) {
	testcase := testCaseUser{
		name:                   "success",
		path:                   "/api/v1/users",
		expectedStatus:         http.StatusOK,
		expectedBodyStartsWith: "{\"status\":",
	}

	e := InitEcho()

	user, _ := config.SeedUser()
	token, _ := middleware.CreateToken(user.ID, user.Name)
	tokenString := fmt.Sprintf("Bearer %s", token)

	password, _ := bcrypt.GenerateFromPassword([]byte("passupdate"), bcrypt.DefaultCost)

	userInput := models.UserInput{
		Name:     "updated",
		Username: "updatedusername",
		Email:    "updated@gmail.com",
		Password: string(password),
	}

	jsonBody, _ := json.Marshal(&userInput)
	bodyReader := bytes.NewReader(jsonBody)

	req := httptest.NewRequest(http.MethodPut, "/api/v1/users", bodyReader)
	rec := httptest.NewRecorder()

	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", tokenString)

	c := e.NewContext(req, rec)

	c.SetPath(testcase.path)

	if assert.NoError(t, controller.Update(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
		body := rec.Body.String()

		assert.True(t, strings.HasPrefix(body, testcase.expectedBodyStartsWith))
	}
}

func TestUpdateUser_Failed(t *testing.T) {
	testcase := testCaseUser{
		name:                   "failed",
		path:                   "/api/v1/users",
		expectedStatus:         http.StatusBadRequest,
		expectedBodyStartsWith: "{\"status\":",
	}

	e := InitEcho()

	password, _ := bcrypt.GenerateFromPassword([]byte("passupdate"), bcrypt.DefaultCost)

	userInput := models.UserInput{
		Name:     "updated",
		Username: "updatedusername",
		Email:    "updated@gmail.com",
		Password: string(password),
	}

	jsonBody, _ := json.Marshal(&userInput)
	bodyReader := bytes.NewReader(jsonBody)

	req := httptest.NewRequest(http.MethodPut, "/api/v1/users", bodyReader)
	rec := httptest.NewRecorder()

	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", "")

	c := e.NewContext(req, rec)

	c.SetPath(testcase.path)

	if assert.NoError(t, controller.Update(c)) {
		assert.Equal(t, http.StatusBadRequest, rec.Code)
		body := rec.Body.String()

		assert.True(t, strings.HasPrefix(body, testcase.expectedBodyStartsWith))
	}
}
