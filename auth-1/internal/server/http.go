package server

import (
	"authentication-test/internal/config"
	"authentication-test/internal/routes"
	"context"

	"github.com/labstack/echo/v5"
	"gorm.io/gorm"
)

func Start(db *gorm.DB, cnfg *config.Config) {
	e := echo.New()

	e.GET("/", func(c *echo.Context) error {
		return c.JSON(200, "Hello world")
	})
	
	routes.RegistersRoutes(e, db, cnfg)

	sc := echo.StartConfig{Address: ":" + cnfg.Port}

	if err := sc.Start(context.Background(), e); err != nil {
		e.Logger.Error("failed to start server", "error", err)
	}

}