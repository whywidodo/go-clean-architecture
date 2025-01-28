package main

import (
	"context"
	"errors"
	"fmt"
	"go-clean-architecture/app"
	"go-clean-architecture/config"
	"go-clean-architecture/constants"
	"go-clean-architecture/repositories"
	"go-clean-architecture/routes"
	"go-clean-architecture/utils"
	"net/http"

	"github.com/go-playground/locales/id"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	id_translations "github.com/go-playground/validator/v10/translations/id"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

// Custom Validator
type CustomValidator struct {
	validator  *validator.Validate
	translator ut.Translator
}

// Passing Variable
var (
	uni         *ut.UniversalTranslator
	echoHandler *echo.Echo
)

var ctx = context.Background()

// Custom Validator and Translation
func (cv *CustomValidator) Validate(i interface{}) error {
	err := cv.validator.Struct(i)
	if err != nil {
		errs := err.(validator.ValidationErrors)
		for _, row := range errs {
			return errors.New(row.Translate(cv.translator))
		}
	}

	return cv.validator.Struct(i)
}

func main() {
	// Open Connection
	if err := config.OpenConnection(); err != nil {
		panic(fmt.Sprintf("Open connection failed: %s", err.Error()))
	}
	defer config.CloseConnection()

	// Connection Database
	DB := config.DBConnection()
	MongoDB := config.ConnectMongo(ctx)

	// Configuration Repository
	repo := repositories.NewRepository(DB, ctx, MongoDB)

	// Configuration Services
	services := app.SetupApp(DB, repo)

	// Routing API
	routes.RoutesApi(echoHandler, services)
	echoHandler.Use(middleware.Logger())

	echoHandler.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowCredentials: constants.TRUE_VALUE,
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{http.MethodGet, http.MethodPut, http.MethodPost, http.MethodDelete},
	}))
	port := fmt.Sprintf(":%s", config.APPPort)
	echoHandler.Logger.Fatal(echoHandler.Start(port))
}

func init() {
	e := echo.New()
	echoHandler = e
	validateCustom := validator.New()

	id := id.New()
	uni = ut.New(id, id)
	trans, _ := uni.GetTranslator("id")

	id_translations.RegisterDefaultTranslations(validateCustom, trans)

	e.Validator = &CustomValidator{
		validator:  validateCustom,
		translator: trans,
	}

	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowCredentials: constants.TRUE_VALUE,
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{http.MethodGet, http.MethodPut, http.MethodPost, http.MethodDelete},
	}))

	e.HTTPErrorHandler = func(err error, c echo.Context) {
		report, ok := err.(*echo.HTTPError)
		if !ok {
			report = echo.NewHTTPError(http.StatusInternalServerError, err.Error())
		}

		// Handle Duplicate Error Response
		if c.Response().Committed {
			return
		}

		result := utils.ResponseJSON(constants.FALSE_VALUE, utils.ToString(report.Code), map[string]string{
			"en": report.Message.(string),
			"id": report.Message.(string),
		}, nil)

		c.Logger().Error(report)
		c.JSON(report.Code, result)
	}
}
