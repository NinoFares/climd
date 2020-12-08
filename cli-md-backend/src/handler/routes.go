package handler

import "github.com/labstack/echo"

func (h *Handler) Register(v1 *echo.Group) {
	patient := v1.Group("/patient")
	patient.GET("", h.getPatientByName)
}
