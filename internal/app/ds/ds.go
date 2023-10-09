package ds

import "time"

type Consultation struct {
	Id          int `gorm:"primarykey;autoIncrement"`
	Name        string
	Description string
	Image       string
	Price       uint
	Status      string
}

type ConsultationRequest struct {
	ConsultationID uint `gorm:"primarykey"`
	RequestID      uint `gorm:"primarykey"`
}

type Request struct {
	RequestID          uint   `gorm:"primarykey"`
	Status             string `gorm:"size:30"`
	StartDate          time.Time
	FormationDate      time.Time
	EndDate            time.Time
	UserID             uint
	ModeratorID        uint
	Consultation_place string
	Consultation_time  string
	Company_name       string
}

type User struct {
	UserID      uint   `gorm:"primarykey"`
	Name        string `gorm:"size:60"`
	Email       string `gorm:"unique;size:60"`
	PhoneNumber string `gorm:"unique;size:15"`
	Role        string `gorm:"size:60"`
}
