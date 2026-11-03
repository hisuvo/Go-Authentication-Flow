package routes

import (
	"authentication-test/internal/config"
	blog "authentication-test/internal/domain/blogs"
	user "authentication-test/internal/domain/users"

	"github.com/labstack/echo/v5"
	"gorm.io/gorm"
)

func RegistersRoutes(e *echo.Echo, db *gorm.DB, cnfg *config.Config) {
	api := e.Group("api/v1")

	user.RegisterRoutes(api, db)
	blog.RegisterRoutes(api, db)
}