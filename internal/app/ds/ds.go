package ds

import "time"

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
	Id          uint   `gorm:"primarykey"`
	Name        string `gorm:"size:60"`
	Email       string `gorm:"unique;size:60"`
	PhoneNumber string `gorm:"unique;size:11"`
	Role        string `gorm:"size:60"`
}

type ConsultationInfo struct {
	Names  []string
	Prices []int
}

type StatusData struct {
	Status string `json:"status"`
}
