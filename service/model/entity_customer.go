package model

import "time"

type InputCustomer struct {
	Nik          string    `json:"nik"`
	FullName     string    `json:"full_name"`
	LegalName    string    `json:"legal_name"`
	PlaceOfBirth string    `json:"place_of_birth"`
	DateOfBirth  time.Time `json:"date_of_birth"`
	Salary       int       `json:"salary"`
	KtpImage     string    `json:"ktp_image"`
	SelfieImage  string    `json:"selfie_image"`
	CreatedAt    time.Time `json:"created_at"`
}

type LimitCustomer struct {
	Nik   string `json:"nik"`
	Tenor int    `json:"tenor"`
	Limit int    `json:"limit"`
}

type ListLimit struct {
	Tenor string `json:"tenor"`
}
