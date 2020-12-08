package store

import (
	"model"

	"github.com/jinzhu/gorm"
)

type PatientStore struct {
	db *gorm.DB
}

func NewPatientStore(db *gorm.DB) *PatientStore {
	return &PatientStore{
		db: db,
	}
}

func (ps *PatientStore) GetByID(id uint) (*model.Patient, error) {
	var p model.Patient
	if err := ps.db.First(&p, id).Error; err != nil {
		if gorm.IsRecordNotFoundError(err) {
			return nil, nil
		}
		return nil, err
	}
	return &p, nil
}

func (ps *PatientStore) GetByNumber(n string) (*model.Patient, error) {
	var p model.Patient
	if err := ps.db.Where(&model.Patient{Number: n}).First(&p).Error; err != nil {
		if gorm.IsRecordNotFoundError(err) {
			return nil, nil
		}
		return nil, err
	}
	return &p, nil
}

func (ps *PatientStore) GetByName(n string) (*model.Patient, error) {
	var p model.Patient
	if err := ps.db.Where(&model.Patient{LastName: n}).First(&p).Error; err != nil {
		if gorm.IsRecordNotFoundError(err) {
			return nil, nil
		}
		return nil, err
	}
	return &p, nil
}

func (ps *PatientStore) Create(p *model.Patient) (err error) {
	return ps.db.Create(p).Error
}

func (ps *PatientStore) Update(p *model.Patient) (err error) {
	return ps.db.Update(p).Error
}
