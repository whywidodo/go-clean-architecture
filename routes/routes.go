package routes

import (
	"go-clean-arhitecture/constants"
	"go-clean-arhitecture/services"
	"go-clean-arhitecture/utils"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func RoutesApi(e *echo.Echo, usecaseSvc services.UsecaseService) {
	e.GET("/api/test-connection", TestApiConnection)

	private := e.Group("")
	private.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowCredentials: constants.TRUE_VALUE,
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{http.MethodGet, http.MethodPut, http.MethodPost, http.MethodDelete},
	}))

}

func TestApiConnection(c echo.Context) error {
	result := utils.ResponseJSON(true, utils.ToString(http.StatusOK), "API Success Initialized", nil)

	return c.JSON(http.StatusOK, result)
}
