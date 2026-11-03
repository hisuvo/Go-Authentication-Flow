package blog

import (
	"authentication-test/internal/domain/blogs/dto"
	"authentication-test/internal/pkg/httpresponse"
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
    var req dto.CreateBlogRequest

    if err := c.Bind(&req); err != nil {
        return httpresponse.BadRequest(c, "Invalid request body", err.Error())
    }

    result, err := h.service.Create(c.Request().Context(), &req)
    if err != nil {
        return httpresponse.BadRequest(c,"Blog created failed", err.Error())
    }

    return httpresponse.Created(c, "Blog created successfully", result)

}

func (h *Handler) FindAll(c *echo.Context) error {
    page, _ := strconv.Atoi(c.QueryParam("page"))
    limit, _ := strconv.Atoi(c.QueryParam("limit"))

    req := &dto.BlogQueryRequest{
        Search: c.QueryParam("search"),
        Page:   page,
        Limit:  limit,
        Sort:   c.QueryParam("sort"),
        Order:  c.QueryParam("order"),
    }

    result, err := h.service.FindAll(c.Request().Context(), req)
    if err != nil {
        return echo.NewHTTPError(http.StatusInternalServerError, "Failed to fetch blogs")
    }

    return c.JSON(http.StatusOK, map[string]interface{}{
        "success": true,
        "data":    result,
    })
}

func (h *Handler) FindByID(c *echo.Context) error {
    id, err := strconv.ParseUint(c.Param("id"), 10, 64)

    if err != nil || id == 0 {
        return echo.NewHTTPError(http.StatusBadRequest, "Invalid blog ID")
    }

    result, err := h.service.FindByID(
        c.Request().Context(),
        uint(id),
    )

    if err != nil {
        if errors.Is(err, ErrBlogNotFound) {
            return echo.NewHTTPError(http.StatusNotFound, err.Error())
        }

        return echo.NewHTTPError(http.StatusInternalServerError, "Failed to fetch blog")
    }

    return c.JSON(http.StatusOK, map[string]interface{}{
        "success": true,
        "data":    result,
    })
}

func (h *Handler) Update(c *echo.Context) error {
    id, err := strconv.ParseUint(c.Param("id"), 10, 64)

    if err != nil || id == 0 {
        return echo.NewHTTPError(http.StatusBadRequest, "Invalid blog ID")
    }

    var req dto.UpdateBlogRequest

    if err := c.Bind(&req); err != nil {
        return echo.NewHTTPError(http.StatusBadRequest, "Invalid request body")
    }

    result, err := h.service.Update(
        c.Request().Context(),
        uint(id),
        &req,
    )

    if err != nil {
        if errors.Is(err, ErrBlogNotFound) {
            return echo.NewHTTPError(http.StatusNotFound, err.Error())
        }

        return echo.NewHTTPError(http.StatusBadRequest, err.Error())
    }

    return c.JSON(http.StatusOK, map[string]interface{}{
        "success": true,
        "message": "Blog updated successfully",
        "data":    result,
    })
}

func (h *Handler) Delete(c *echo.Context) error {
    id, err := strconv.ParseUint(c.Param("id"), 10, 64)

    if err != nil || id == 0 {
        return echo.NewHTTPError(http.StatusBadRequest, "Invalid blog ID")
    }

    if err := h.service.Delete(c.Request().Context(), uint(id)); err != nil {
        if errors.Is(err, ErrBlogNotFound) {
            return echo.NewHTTPError(http.StatusNotFound, err.Error())
        }

        return echo.NewHTTPError(http.StatusInternalServerError, "Failed to delete blog")
    }

    return c.JSON(http.StatusOK, map[string]interface{}{
        "success": true,
        "message": "Blog deleted successfully",
    })
}