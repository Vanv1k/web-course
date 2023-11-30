package ds

import (
	"time"

	"github.com/Vanv1k/web-course/internal/app/role"
)

type Consultation struct {
	Id          uint `gorm:"primarykey;autoIncrement"`
	Name        string
	Description string
	Image       string
	Price       int
	Status      string
}

type ConsultationRequest struct {
	Consultationid int `gorm:"primarykey"`
	Requestid      int `gorm:"primarykey"`
}

type Request struct {
	Id                 uint   `gorm:"primarykey"`
	Status             string `gorm:"size:30"`
	StartDate          time.Time
	FormationDate      time.Time
	EndDate            time.Time
	UserID             uint
	ModeratorID        *uint
	Consultation_place string
	Consultation_time  time.Time
	Company_name       string
}

type User struct {
	Id          uint      `gorm:"primarykey"`
	Name        string    `json:"name"`
	Login       string    `json:"login"`
	Email       string    `json:"email"`
	PhoneNumber string    `json:"phoneNumber"`
	Role        role.Role `sql:"type:string;"`
	Password    string    `gorm:"size:60"`
}

type ConsultationInfo struct {
	Names  []string
	Prices []int
}

type StatusData struct {
	Status string `json:"status"`
}
