package user

import (
	"authentication-test/internal/domain/users/dto"
	"errors"
	"net/http"
	"strconv"

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

func (h *Handler) Create(c *echo.Context) error {
    var req dto.CreateUserRequest

    if err := c.Bind(&req); err != nil {
        return echo.NewHTTPError(http.StatusBadRequest, "Invalid request body")
    }

    if req.Name == "" || req.Email == "" || len(req.Password) < 8 {
        return echo.NewHTTPError(http.StatusBadRequest, "Invalid user information")
    }

    user, err := h.service.Create(c.Request().Context(), &req)
    if err != nil {
        if errors.Is(err, ErrEmailExists) {
            return echo.NewHTTPError(http.StatusConflict, err.Error())
        }

        return echo.NewHTTPError(http.StatusInternalServerError, "Failed to create user")
    }

    return c.JSON(http.StatusCreated, map[string]interface{}{
        "success": true,
        "message": "User created successfully",
        "data":    user,
    })
}

func (h *Handler) FindAll(c *echo.Context) error {
    users, err := h.service.FindAll(c.Request().Context())
    if err != nil {
        return echo.NewHTTPError(http.StatusInternalServerError, "Failed to fetch users")
    }

    return c.JSON(http.StatusOK, map[string]interface{}{
        "success": true,
        "data":    users,
    })
}

func (h *Handler) FindByID(c *echo.Context) error {
    id, err := strconv.ParseUint(c.Param("id"), 10, 64)
    if err != nil || id == 0 {
        return echo.NewHTTPError(http.StatusBadRequest, "Invalid user ID")
    }

    user, err := h.service.FindByID(
        c.Request().Context(),
        uint(id),
    )

    if err != nil {
        if errors.Is(err, ErrUserNotFound) {
            return echo.NewHTTPError(http.StatusNotFound, err.Error())
        }

        return echo.NewHTTPError(http.StatusInternalServerError, "Failed to fetch user")
    }

    return c.JSON(http.StatusOK, map[string]interface{}{
        "success": true,
        "data":    user,
    })
}

func (h *Handler) Update(c *echo.Context) error {
    id, err := strconv.ParseUint(c.Param("id"), 10, 64)
    if err != nil || id == 0 {
        return echo.NewHTTPError(http.StatusBadRequest, "Invalid user ID")
    }

    var req dto.UpdateUserRequest

    if err := c.Bind(&req); err != nil {
        return echo.NewHTTPError(http.StatusBadRequest, "Invalid request body")
    }

    user, err := h.service.Update(
        c.Request().Context(),
        uint(id),
        &req,
    )

    if err != nil {
        if errors.Is(err, ErrUserNotFound) {
            return echo.NewHTTPError(http.StatusNotFound, err.Error())
        }

        return echo.NewHTTPError(http.StatusInternalServerError, "Failed to update user")
    }

    return c.JSON(http.StatusOK, map[string]interface{}{
        "success": true,
        "message": "User updated successfully",
        "data":    user,
    })
}

func (h *Handler) Delete(c *echo.Context) error {
    id, err := strconv.ParseUint(c.Param("id"), 10, 64)
    if err != nil || id == 0 {
        return echo.NewHTTPError(http.StatusBadRequest, "Invalid user ID")
    }

    err = h.service.Delete(c.Request().Context(), uint(id))
    if err != nil {
        if errors.Is(err, ErrUserNotFound) {
            return echo.NewHTTPError(http.StatusNotFound, err.Error())
        }

        return echo.NewHTTPError(http.StatusInternalServerError, "Failed to delete user")
    }

    return c.JSON(http.StatusOK, map[string]interface{}{
        "success": true,
        "message": "User deleted successfully",
    })
}