package controllers

import (
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"testing"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

var (
	userJSON = `{
		"FirstName":    "Adonis",
		"LastName":     "Villanueva",
		"UserName":     "adonisabril",
		"OrgName":      "Cognosante",
		"Location":     "Miami",
		"Title":        "Solutions Architect",
		"Address1":     "1234 Address 1",
		"Address2":     "",
		"City":     	"Sacramento",
		"State":      	"CA",
		"ZipCode":      "98415",
		"Country":      "USA",
		"Gender":       "Male",
		"Account":      "Personal",
		"DOB":          "04/13/1989",
		"Email":        "adonis@alwayswanderlust.com",
		"Password":     "Password123",
		"LastAccessed": "",
		"CreatedDate":	""
	}`
)

func GetEnv() (appName string) {
	err := godotenv.Load()
	if err != nil {
		log.Println("INFO: unable to load .env file")
	}

	appName, ok := os.LookupEnv("APP_NAME")
	if ok != true || appName == "" {
		log.Println("INFO: APP_NAME not set in environment; defaulting to log")
		appName = "log"
	}

	return
}

func TestCreateUser(t *testing.T) {
	// Setup
	e := echo.New()
	req := httptest.NewRequest(http.MethodPost, "/", strings.NewReader(userJSON))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/users")

	// Assertions
	if assert.NoError(t, CreateUser(c)) {
		assert.Equal(t, http.StatusCreated, rec.Code)
		assert.Equal(t, userJSON, rec.Body.String())
	}
}

func TestGetUser(t *testing.T) {
	// Setup
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/users/:userId")
	c.SetParamNames("userId")
	c.SetParamValues("63bf0b1c973a94f082f1387e")

	// Assertions
	if assert.NoError(t, GetAUser(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
		assert.Equal(t, userJSON, rec.Body.String())
	}
}

func TestDeleteAUser(t *testing.T) {
	// Setup
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/users/:userId")
	c.SetParamNames("userId")
	c.SetParamValues("63bf0b1c973a94f082f1387e")

	// Assertions
	if assert.NoError(t, DeleteAUser(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
		assert.Equal(t, userJSON, rec.Body.String())
	}
}

func TestGetAllUsers(t *testing.T) {
	// Setup
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/users", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	// Assertions
	if assert.NoError(t, GetAllUsers(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
		assert.Equal(t, userJSON, rec.Body.String())
	}
}

func TestEditAUser(t *testing.T) {
	// Setup
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/users/:userId")
	c.SetParamNames("userId")
	c.SetParamValues("63bf0b1c973a94f082f1387e")
	// Assertions
	if assert.NoError(t, EditAUser(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
		assert.Equal(t, userJSON, rec.Body.String())
	}
}
