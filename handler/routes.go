package handler

import (
	"github.com/labstack/echo/v4"
)

func (h *Handler) Register(v1 *echo.Group) {

	v1.GET("/room/:no", h.GetGuestsByRoom)

}
