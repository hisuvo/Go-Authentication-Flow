package user

import (
	"github.com/labstack/echo/v5"
	"gorm.io/gorm"
)

func RegisterRoutes(
    api *echo.Group,
    db *gorm.DB,
) {
    repository := NewRepository(db)
    service := NewService(repository)
    handler := NewHandler(service)

    users := api.Group("/users")
    
    users.POST("", handler.Create)
    users.GET("", handler.FindAll)
    users.GET("/:id", handler.FindByID)
    users.PUT("/:id", handler.Update)
    users.DELETE("/:id", handler.Delete)
}