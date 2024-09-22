package middleware

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
)

type ErrorResponse struct {
	Message string `json:"message"`
	Code    int    `json:"code,omitempty"`
}

type ValidationErrorResponse struct {
	Message string            `json:"message"`
	Errors  map[string]string `json:"errors"`
}

func CustomHTTPErrorHandler(err error, c echo.Context) {

	log.Error(err)

	if httpError, ok := err.(*echo.HTTPError); ok {
		handleHTTPError(c, httpError)
		return
	}

	if httpError, ok := err.(*echo.HTTPError); ok && httpError.Code == http.StatusUnauthorized {
		handleJWTError(c)
		return
	}

	c.JSON(http.StatusInternalServerError, ErrorResponse{
		Message: "Internal Server Error",
		Code:    http.StatusInternalServerError,
	})
}

func handleHTTPError(c echo.Context, httpError *echo.HTTPError) {
	var message string
	if httpError.Message != nil {
		message = httpError.Message.(string)
	} else {
		message = http.StatusText(httpError.Code)
	}

	c.JSON(httpError.Code, ErrorResponse{
		Message: message,
		Code:    httpError.Code,
	})
}

func handleJWTError(c echo.Context) {

	c.JSON(http.StatusUnauthorized, ErrorResponse{
		Message: "Invalid or expired token",
		Code:    http.StatusUnauthorized,
	})
}
