package handler

import (
	"github.com/labstack/echo/v4"
)

func (h *Handler) Register(v1 *echo.Group) {
	images := v1.Group("images")
	images.POST("/resize", h.resize)
}
