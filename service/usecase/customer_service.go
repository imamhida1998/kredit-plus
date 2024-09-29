package usecase

import "kredit-plus/service/model/request"

type CustomerUsecase interface {
	CreateCustomer(params *request.InputCustomer) error
}
