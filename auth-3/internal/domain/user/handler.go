package user

import (
	"golang-authentication/internal/domain/user/dto"

	"github.com/labstack/echo/v5"
)

type Handler struct {
	service Service
}

func NewHandler(service Service) *Handler {
	return &Handler{
		service: service,
	}
}

func (h *Handler) CreateUser(c *echo.Context) error{
	var user dto.RegisterRequest

	if err := c.Bind(&user); err != nil {
		return err
	}

	// if err := c.Bind(&user); err != nil {
	// 	return err
	// }

	result, err := h.service.CreateUser(&user)

	if err != nil {
		return c.JSON(400, err.Error())
	}
	
	return c.JSON(200, result)
}