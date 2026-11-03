package httpresponse

import (
	"net/http"

	"github.com/labstack/echo/v5"
)

// Response represent a standard API response
type Response struct {
	Success bool 	`json:"success"`
	Message string  `json:"message"`
	Data any 		`json:"data,omitempty"`
	Errors any 		`json:"errors,omitempty"`
}

// Success return a successful reponse
func Success(c *echo.Context, status int, message string, data any) error {
	return c.JSON(status, Response{
		Success: true,
		Message: message,
		Data: data,
	})
}

// Create returns http 201 created
func Created(c *echo.Context, message string, data any) error {
	return Success(c, http.StatusCreated, message, data)
}

// Ok returns http 200 ok
func Ok(c *echo.Context, message string, data any) error {
	return Success(c, http.StatusOK, message, data)
}

// Error return a errors reponse
func Error(c *echo.Context, status int, message string, errors any) error {
	return c.JSON(status, Response{
		Success: false,
		Message: message,
		Errors: errors,
	})
}

// BadRequest returns HTTP 400.
func BadRequest(c *echo.Context, message string, errors interface{}) error {
	return Error(c, http.StatusBadRequest, message, errors)
}

// Unauthorized returns HTTP 401.
func Unauthorized(c *echo.Context, message string) error {
	return Error(c, http.StatusUnauthorized, message, nil)
}

// Forbidden returns HTTP 403.
func Forbidden(c *echo.Context, message string) error {
	return Error(c, http.StatusForbidden, message, nil)
}

// NotFound returns HTTP 404.
func NotFound(c *echo.Context, message string) error {
	return Error(c, http.StatusNotFound, message, nil)
}

// Conflict returns HTTP 409.
func Conflict(c *echo.Context, message string, errors interface{}) error {
	return Error(c, http.StatusConflict, message, errors)
}

// InternalServerError returns HTTP 500.
func InternalServerError(c *echo.Context, message string) error {
	return Error(c, http.StatusInternalServerError, message, nil)
}