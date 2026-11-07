package user

import (
	"github.com/labstack/echo/v5"
	"gorm.io/gorm"
)



func RegisterRoute(e *echo.Echo, db *gorm.DB) {
	repository := NewRegister(db)
	service := NewService(repository)
	handler := NewHandler(service)

	route := e.Group("/api/v1/users")

	route.POST("",handler.CreateUser)
}