package repository

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"github.com/Vanv1k/web-course/internal/app/ds"
)

type Repository struct {
	db *gorm.DB
}

func New(dsn string) (*Repository, error) {
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	return &Repository{
		db: db,
	}, nil
}

// func (r *Repository) GetConsultationByID(id int) (*ds.Consultation, error) {
// 	consultation := &ds.Consultation{}

// 	err := r.db.First(consultation, "id = ?", id).Error
// 	if err != nil {
// 		return nil, err
// 	}

// 	return consultation, nil
// }

func (r *Repository) CreateConsultation(consultation ds.Consultation) error {
	return r.db.Create(consultation).Error
}

func (r *Repository) GetAllConsultations() ([]ds.Consultation, error) {
	var consultations []ds.Consultation

	err := r.db.Find(&consultations).Error
	if err != nil {
		return nil, err
	}

	return consultations, nil
}
