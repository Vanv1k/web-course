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

func (r *Repository) CreateRequest(request ds.Request) error {
	return r.db.Create(request).Error
}

func (r *Repository) GetAllRequests() ([]ds.Request, error) {
	var requests []ds.Request
	err := r.db.Find(&requests, "status = 'active'").Error
	if err != nil {
		return nil, err
	}

	return requests, nil
}
