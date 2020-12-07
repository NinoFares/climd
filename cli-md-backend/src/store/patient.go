package store

import (
	"modele"

	"github.com/jinzhu/gorm"
)

type PatientStore struct {
	db *gorm.DB
}

func NewUserStore(db *gorm.DB) *PatientStore {
	return &PatientStore{
		db: db,
	}
}

func (ps *PatientStore) GetByID(id uint) (*modele.Patient, error) {
	var p modele.Patient
	if err := ps.db.First(&p, id).Error; err != nil {
		if gorm.IsRecordNotFoundError(err) {
			return nil, nil
		}
		return nil, err
	}
	return &p, nil
}

func (ps *PatientStore) GetByNumber(n string) (*modele.Patient, error) {
	var p modele.Patient
	if err := ps.db.Where(&modele.Patient{Number: n}).First(&p).Error; err != nil {
		if gorm.IsRecordNotFoundError(err) {
			return nil, nil
		}
		return nil, err
	}
	return &p, nil
}

func (ps *PatientStore) Create(p *modele.Patient) (err error) {
	return ps.db.Create(p).Error
}

func (ps *PatientStore) Update(p *modele.Patient) (err error) {
	return ps.db.Update(p).Error
}
