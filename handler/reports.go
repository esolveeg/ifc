package handler

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
)

func (h *Handler) GetGuestsByRoom(c echo.Context) error {
	room := c.Param("room")
	fmt.Println(room)
	_, err := h.db.Raw("SELECT ROOM ,  GUEST_NAEM_ID , GINDER , RESERVATION_NAME_ID , BIRTH_DATE , COUNTRY , GUEST_NAME , ARRIVAL , DEPARTURE  FROM POLICE_IFC WHERE ROOM = 1;").Rows()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	return c.JSON(http.StatusOK, "ok")
}
