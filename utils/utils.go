package utils

import (
	"encoding/json"
	"go-clean-arhitecture/models"
	"strconv"
	"strings"
	"time"

	"github.com/labstack/echo/v4"
)

func ReplaceSQL(old, searchPattern string) string {
	tmpCount := strings.Count(old, searchPattern)
	for m := 1; m <= tmpCount; m++ {
		old = strings.Replace(old, searchPattern, "$"+strconv.Itoa(m), 1)
	}

	return old
}

func BindValidateStruct(ctx echo.Context, i interface{}) error {
	if err := ctx.Bind(i); err != nil {
		return err
	}

	if err := ctx.Validate(i); err != nil {
		return err
	}

	return nil
}

func ResponseJSON(success bool, code string, msg string, result interface{}) models.Response {
	tm := time.Now()
	response := models.Response{
		Success:          success,
		StatusCode:       code,
		Result:           result,
		Message:          msg,
		ResponseDatetime: tm,
	}

	return response
}

func ToString(data interface{}) string {
	bytes, _ := json.Marshal(data)
	return string(bytes)
}

func TimestampNow() string {
	return time.Now().Format("2006-01-02 15:04:05")
}
