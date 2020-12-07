package patient

import "model"

type Store interface {
	GetByID(uint) (*model.Patient, error)
	GetByEmail(string) (*model.Patient, error)
	Create(*model.Patient) error
	Update(*model.Patient) error
}
