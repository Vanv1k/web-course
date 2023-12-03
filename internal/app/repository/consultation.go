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

func (r *Repository) GetConsultationsByRequestID(userID uint) (ds.ConsultationInfo, error) {
	var consultationRequests []ds.ConsultationRequest
	var consultationInfo ds.ConsultationInfo
	var request ds.Request

	err := r.db.Find(&request, "status = ? AND user_id = ?", "active", userID).Error
	if err != nil {
		return consultationInfo, err
	}

	err = r.db.Find(&consultationRequests, "Requestid = ?", request.Id).Error
	if err != nil {
		return consultationInfo, err
	}

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

	err := r.minioClient.RemoveServiceImage(id)
	if err != nil {
		return err
	}

	return r.db.Exec("UPDATE consultations SET status = 'deleted' WHERE id=?", id).Error
}

func (r *Repository) CreateConsultation(consultation ds.Consultation) error {
	return r.db.Create(&consultation).Error
}

func (r *Repository) GetAllConsultations(userID uint) ([]ds.Consultation, uint, error) {
	var consultations []ds.Consultation
	var request ds.Request
	err := r.db.Find(&consultations, "status = 'active'").Error
	if err != nil {
		return nil, 0, err
	}
	// if userID == 0 {
	// 	return consultations, 0, nil
	// }
	err = r.db.Where("status = ? AND user_id = ?", "active", userID).First(&request).Error
	if err != nil {
		return consultations, 0, nil
	}

	return consultations, request.Id, nil
}

func (r *Repository) GetConsultationsByPrice(maxPrice int, userID uint) ([]ds.Consultation, uint, error) {
	var consultations []ds.Consultation
	var request ds.Request
	err := r.db.Where("status = ? AND price <= ?", "active", maxPrice).Find(&consultations).Error
	if err != nil {
		return nil, 0, err
	}
	err = r.db.Where("status = ? AND user_id = ?", "active", userID).First(&request).Error
	if err != nil {
		return consultations, 0, nil
	}

	return consultations, request.Id, nil
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

func (r *Repository) AddConsultationToRequest(consultationID int, userID uint) error {
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

func (r *Repository) AddConsultationImage(id int, imageBytes []byte, contentType string) error {
	// Удаление существующего изображения (если есть)
	err := r.minioClient.RemoveServiceImage(id)
	if err != nil {
		return err
	}

	// Загрузка нового изображения в MinIO
	imageURL, err := r.minioClient.UploadServiceImage(id, imageBytes, contentType)
	if err != nil {
		return err
	}

	// Обновление информации об изображении в БД (например, ссылки на MinIO)
	err = r.db.Model(&ds.Consultation{}).Where("id = ?", id).Update("image", imageURL).Error
	if err != nil {
		return err
	}

	return nil
}
