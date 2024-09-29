package usecase

import (
	"kredit-plus/service/model"
	"kredit-plus/service/model/request"
	"kredit-plus/service/repo"
	"math"
	"strconv"
	"time"
)

type customerUsecase struct {
	repoUser repo.CustomerRepository
}

func NewCustomerService(repo repo.CustomerRepository) CustomerUsecase {
	return &customerUsecase{repo}
}

func (usecase *customerUsecase) CreateCustomer(params *request.InputCustomer) error {

	tanggalLahir, err := time.Parse("02-01-2006", params.DateOfBirth)
	if err != nil {
		return err
	}
	request := &model.InputCustomer{
		Nik:          params.Nik,
		FullName:     params.FullName,
		LegalName:    params.FullName,
		PlaceOfBirth: params.PlaceOfBirth,
		DateOfBirth:  tanggalLahir,
		Salary:       params.Salary,
		KtpImage:     params.KtpImage,
		SelfieImage:  params.KtpImage,
	}

	err = usecase.repoUser.InputCustomer(request)
	if err != nil {
		return err
	}

	cekTenor, err := usecase.repoUser.ListTenor()
	if err != nil {
		return err
	}

	for _, tenor := range cekTenor {
		AngkaTenor, _ := strconv.Atoi(tenor.Tenor)
		Limit := float64(params.Salary) * 0.5
		Limit = math.Round(Limit)
		limitTenor := int(Limit) * AngkaTenor

		insertTenor := &model.LimitCustomer{
			Nik:   params.Nik,
			Tenor: AngkaTenor,
			Limit: limitTenor,
		}

		err := usecase.repoUser.InputLimit(insertTenor)
		if err != nil {
			return err
		}

	}

	return nil

}
