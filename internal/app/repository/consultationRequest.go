package repository

import (
	"github.com/Vanv1k/web-course/internal/app/ds"
)

func (r *Repository) DeleteConsultationRequest(idC int, idR int) error {
	var consultationRequest ds.ConsultationRequest
	err := r.db.Where("requestid = ? AND consultationid = ?", idR, idC).First(&consultationRequest).Error
	if err != nil {
		return err
	}
	return r.db.Delete(&consultationRequest).Error
}

func (r *Repository) ClearShoppingCart(requestID uint) error {
	var consultationRequests []ds.ConsultationRequest

	err := r.db.Find(&consultationRequests, "Requestid = ?", requestID).Error
	if err != nil {
		return err
	}

	err = r.db.Delete(&consultationRequests).Error
	if err != nil {
		return err
	}

	return nil
}
