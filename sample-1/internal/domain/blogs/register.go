package blog

import (
	"github.com/labstack/echo/v5"
	"gorm.io/gorm"
)

func RegisterRoutes(api *echo.Group, db *gorm.DB) {
    repository := NewRepository(db)
    service := NewService(repository)
    handler := NewHandler(service)

    blogs := api.Group("/blogs")

    blogs.POST("", handler.Create)
    blogs.GET("", handler.FindAll)
    blogs.GET("/:id", handler.FindByID)
    blogs.PUT("/:id", handler.Update)
    blogs.DELETE("/:id", handler.Delete)
}