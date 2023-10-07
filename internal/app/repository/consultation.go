package repository

import (
	"github.com/Vanv1k/web-course/internal/app/ds"
)

func (r *Repository) GetConsultationByID(id int) (*ds.Consultation, error) {
	consultation := &ds.Consultation{}

	err := r.db.First(consultation, "id = ?", id).Error
	if err != nil {
		return nil, err
	}

	return consultation, nil
}

func (r *Repository) DeleteConsultation(id int) error {
	return r.db.Exec("UPDATE consultations SET status = 'deleted' WHERE id=?", id).Error
}

func (r *Repository) CreateConsultation(consultation ds.Consultation) error {
	return r.db.Create(consultation).Error
}

func (r *Repository) GetAllConsultations() ([]ds.Consultation, error) {
	var consultations []ds.Consultation
	err := r.db.Find(&consultations, "status = 'active'").Error
	if err != nil {
		return nil, err
	}

	return consultations, nil
}
