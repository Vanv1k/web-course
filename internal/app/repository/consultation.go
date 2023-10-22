package repository

import (
	"errors"
	"log"
	"time"

	"github.com/Vanv1k/web-course/internal/app/ds"
	"gorm.io/gorm"
)

func (r *Repository) GetConsultationByID(id uint) (*ds.Consultation, error) {
	consultation := &ds.Consultation{}

	err := r.db.First(consultation, "id = ?", id).Error
	if err != nil {
		return nil, err
	}

	return consultation, nil
}

func (r *Repository) GetConsultationsByRequestID(id int) (ds.ConsultationInfo, error) {
	var consultationRequests []ds.ConsultationRequest
	var consultationInfo ds.ConsultationInfo

	err := r.db.Find(&consultationRequests, "Requestid = ?", id).Error
	if err != nil {
		return consultationInfo, err
	}
	log.Println(consultationRequests)

	for _, consultationRequest := range consultationRequests {
		consultation, err := r.GetConsultationByID(uint(consultationRequest.Consultationid))
		if err != nil {
			return consultationInfo, err
		}
		consultationInfo.Names = append(consultationInfo.Names, consultation.Name)
		consultationInfo.Prices = append(consultationInfo.Prices, consultation.Price)
	}

	return consultationInfo, nil
}

func (r *Repository) DeleteConsultation(id int) error {
	return r.db.Exec("UPDATE consultations SET status = 'deleted' WHERE id=?", id).Error
}

func (r *Repository) CreateConsultation(consultation ds.Consultation) error {
	return r.db.Create(&consultation).Error
}

func (r *Repository) GetAllConsultations() ([]ds.Consultation, error) {
	var consultations []ds.Consultation
	err := r.db.Find(&consultations, "status = 'active'").Error
	if err != nil {
		return nil, err
	}

	return consultations, nil
}

func (r *Repository) GetConsultationsByPrice(maxPrice int) ([]ds.Consultation, error) {
	var consultations []ds.Consultation
	err := r.db.Where("status = ? AND price <= ?", "active", maxPrice).Find(&consultations).Error
	if err != nil {
		return nil, err
	}

	return consultations, nil
}

func (r *Repository) UpdateConsultation(id int, consultation ds.Consultation) error {
	// Проверяем, существует ли консультация с указанным ID.
	existingConsultation, err := r.GetConsultationByID(uint(id))
	if err != nil {
		return err // Возвращаем ошибку, если консультация не найдена.
	}

	// Обновляем поля существующей консультации.
	existingConsultation.Name = consultation.Name
	existingConsultation.Description = consultation.Description
	existingConsultation.Image = consultation.Image
	existingConsultation.Price = consultation.Price
	existingConsultation.Status = consultation.Status

	// Сохраняем обновленную консультацию в базу данных.
	if err := r.db.Model(ds.Consultation{}).Where("id = ?", id).Updates(existingConsultation).Error; err != nil {
		return err
	}
	return nil
}

func (r *Repository) AddConsultationToRequest(consultationID int, userID int) error {
	var consultationRequest ds.ConsultationRequest
	var request ds.Request
	err := r.db.Where("status = ? AND user_id = ?", "active", userID).First(&request).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			// Запись не найдена, создаем заявку
			newRequest := ds.Request{
				Status:             "active",
				StartDate:          time.Now(),
				UserID:             uint(userID),
				Consultation_place: "Discord #abfa1213",
				Consultation_time:  time.Now(),
				Company_name:       "IT",
			}
			err = r.db.Create(&newRequest).Error
			if err != nil {
				return err
			}
			err = r.db.Where("status = ? AND user_id = ?", "active", userID).First(&request).Error
			if err != nil {
				return err
			}
		} else {
			return err
		}
	}
	err = r.db.Where("requestid = ? AND consultationid = ?", request.Id, consultationID).First(&consultationRequest).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			consultationRequest.Consultationid = consultationID
			consultationRequest.Requestid = int(request.Id)
			err = r.db.Create(&consultationRequest).Error
			if err != nil {
				return err
			}
			log.Println("1")
			return nil
		}
	} else {
		return gorm.ErrRecordNotFound
	}
	log.Println("123132")
	return gorm.ErrRecordNotFound
}
