package server

import (
	"golang-authentication/internal/config"
	"golang-authentication/internal/domain/user"

	"github.com/labstack/echo/v5"
	"gorm.io/gorm"
)

func Start(db *gorm.DB, env *config.Config) {
	e := echo.New()

	e.GET("/", func(c *echo.Context) error {
		return c.JSON(200, map[string]any{
			"success":true,
			"message": "Hello, World!",
		})
	})

	user.RegisterRoute(e, db)

	e.Start(":"+env.Port)
}