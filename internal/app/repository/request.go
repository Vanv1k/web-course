package repository

import (
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

func (r *Repository) DeleteRequest(id int) error {
	return r.db.Exec("UPDATE requests SET status = 'deleted' WHERE id=?", id).Error
}

func (r *Repository) GetAllRequests() ([]ds.Request, error) {
	var requests []ds.Request
	err := r.db.Find(&requests, "status = 'active'").Error
	if err != nil {
		return nil, err
	}

	return requests, nil
}

func (r *Repository) UpdateRequest(id int, request ds.Request) error {
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
	if err := r.db.Model(ds.Request{}).Where("id = ?", id).Updates(existingRequest).Error; err != nil {
		return err
	}
	return nil
}

func (r *Repository) UpdateRequestStatus(id int, status string) error {
	// Проверяем, существует ли консультация с указанным ID.
	existingRequest, err := r.GetRequestByID(id)
	if err != nil {
		return err // Возвращаем ошибку, если консультация не найдена.
	}

	// Обновляем поля существующей консультации.
	existingRequest.Status = status

	// Сохраняем обновленную консультацию в базу данных.
	if err := r.db.Model(ds.Request{}).Where("id = ?", id).Updates(existingRequest).Error; err != nil {
		return err
	}
	return nil
}
