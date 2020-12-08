package handler

import (
	"net/http"

	"github.com/labstack/echo"
)

func (h *Handler) getPatientByName(c echo.Context) error {
	p, err := h.patientStore.GetByName("Test")
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}
	if p == nil {
		return c.JSON(http.StatusNotFound, nil)
	}
	return c.JSON(http.StatusOK, "user")
}
