package request

type InputCustomer struct {
	Nik          string `json:"nik"`
	FullName     string `json:"full_name"`
	PlaceOfBirth string `json:"place_of_birth"`
	DateOfBirth  string `json:"date_of_birth"`
	Salary       int    `json:"salary"`
	KtpImage     string `json:"ktp_image"`
	SelfieImage  string `json:"selfie_image"`
}

type UpdateLimit struct {
	Nik   string `json:"nik"`
	Tenor int    `json:"tenor"`
	Limit int    `json:"limit"`
}

type RequestTransaction struct {
	Nik string `json:"nik"`
	Otp string `json:"otp"`
}
