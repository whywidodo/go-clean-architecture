package routes

import (
	"go-clean-architecture/config"
	"go-clean-architecture/constants"
	"go-clean-architecture/services"
	"go-clean-architecture/utils"
	"log"
	"net/http"

	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func RoutesApi(e *echo.Echo, usecaseSvc services.UsecaseService) {
	e.GET("/api/test-connection", TestApiConnection)

	public := e.Group("")
	public.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowCredentials: constants.TRUE_VALUE,
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{http.MethodGet, http.MethodPut, http.MethodPost, http.MethodDelete},
	}))

	private := e.Group("/private")
	// Add Log on Every Process
	private.Use(middleware.BodyDump(func(c echo.Context, reqBody, resBody []byte) {
		log.Println("[Start]")
		log.Println("EndPoint :", c.Path())
		log.Println("Header :", c.Request().Header)
		log.Println("Body :", string(reqBody))
		log.Println("Response :", string(resBody))
		log.Println("[End]")
	}))
	private.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowCredentials: true,
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{http.MethodGet, http.MethodPut, http.MethodPost, http.MethodDelete},
	}))
	// Use JWT Middleware from Labstack Echo JWT
	private.Use(echojwt.WithConfig(echojwt.Config{
		SigningKey: []byte(config.APPKey),
	}))

}

func TestApiConnection(c echo.Context) error {
	result := utils.ResponseJSON(true, utils.ToString(http.StatusOK), map[string]string{
		"en": "API Connection Successfull",
		"id": "Koneksi API Berhasil",
	}, nil)

	return c.JSON(http.StatusOK, result)
}
