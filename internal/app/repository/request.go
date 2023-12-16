package repository

import (
	"fmt"
	"time"

	"github.com/Vanv1k/web-course/internal/app/ds"
)

func (r *Repository) GetRequestByID(id int) (*ds.Request, error) {
	request := &ds.Request{}

	err := r.db.First(request, "id = ?", id).Error
	if err != nil {
		return nil, err
	}

	return request, nil
}

func (r *Repository) DeleteRequest(id int, userID uint) error {
	return r.db.Exec("UPDATE requests SET status = 'deleted' WHERE id=? AND user_id=?", id, userID).Error
}

func (r *Repository) GetAllRequests() ([]ds.Request, error) {
	var requests []ds.Request
	err := r.db.Find(&requests, "status <> 'deleted'").Error
	if err != nil {
		return nil, err
	}

	return requests, nil
}

func (r *Repository) GetAllUserRequests(userID uint) ([]ds.Request, error) {
	var requests []ds.Request
	err := r.db.Where("status <> 'deleted' AND user_id=?", userID).Find(&requests).Error
	if err != nil {
		return nil, err
	}

	return requests, nil
}

func (r *Repository) GetRequestsByStatus(status string) ([]ds.Request, error) {
	var requests []ds.Request
	err := r.db.Where("status = ?", status).Find(&requests).Error
	if err != nil {
		return nil, err
	}

	return requests, nil
}

func (r *Repository) GetRequestsByDate(startDate time.Time, endDate time.Time) ([]ds.Request, error) {
	var requests []ds.Request
	if !endDate.IsZero() {
		err := r.db.Where("formation_date >= ? AND formation_date <= ?", startDate, endDate).Find(&requests).Error
		if err != nil {
			return nil, err
		}
		return requests, nil
	}

	err := r.db.Where("formation_date >= ?", startDate).Find(&requests).Error
	if err != nil {
		return nil, err
	}
	return requests, nil
}

func (r *Repository) UpdateRequest(id int, userID uint, request ds.Request) error {
	// Проверяем, существует ли консультация с указанным ID.
	existingRequest, err := r.GetRequestByID(id)
	if err != nil {
		return err // Возвращаем ошибку, если консультация не найдена.
	}

	// Обновляем поля существующей консультации.
	existingRequest.Consultation_place = request.Consultation_place
	existingRequest.Consultation_time = request.Consultation_time
	existingRequest.Company_name = request.Company_name

	// Сохраняем обновленную консультацию в базу данных.
	if err := r.db.Model(ds.Request{}).Where("id = ? AND user_id=?", id, userID).Updates(existingRequest).Error; err != nil {
		return err
	}
	return nil
}

func (r *Repository) UpdateRequestStatus(userID uint, id int, status string) error {
	// Проверяем, существует ли заявка с указанным ID.
	existingRequest, err := r.GetRequestByID(id)
	if err != nil {
		return err // Возвращаем ошибку, если заявка не найдена.
	}

	if existingRequest.Status != "formed" {
		return fmt.Errorf("Неправильный статус: %s, ожадался 'formed'", existingRequest.Status)
	}

	// Обновляем поля существующей консультации.
	existingRequest.Status = status
	existingRequest.ModeratorID = userID
	// Сохраняем обновленную консультацию в базу данных.
	if err := r.db.Model(ds.Request{}).Where("id = ?", id).Updates(existingRequest).Error; err != nil {
		return err
	}
	return nil
}

func (r *Repository) UpdateRequestStatusToSended(id int, userID uint) error {
	// Проверяем, существует ли заявка с указанным ID.
	existingRequest, err := r.GetRequestByID(id)
	if err != nil {
		return err // Возвращаем ошибку, если заявка не найдена.
	}

	if existingRequest.Status != "active" {
		return fmt.Errorf("Неправильный статус: %s, ожадался 'active'", existingRequest.Status)
	}

	// Обновляем поля существующей консультации.
	existingRequest.Status = "formed"

	// Сохраняем обновленную консультацию в базу данных.
	if err := r.db.Model(ds.Request{}).Where("id = ? AND user_id=?", id, userID).Updates(existingRequest).Error; err != nil {
		return err
	}
	return nil
}
