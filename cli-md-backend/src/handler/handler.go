package handler

import "patient"

type Handler struct {
	patientStore patient.Store
}

func NewHandler(ps patient.Store) *Handler {
	return &Handler{
		patientStore: ps,
	}
}
