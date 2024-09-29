package repo

import (
	"kredit-plus/service/model"
	"kredit-plus/service/model/request"
)

type CustomerRepository interface {
	InputCustomer(params *model.InputCustomer) error
	InputLimit(params *model.LimitCustomer) error
	InsertTenor(params *model.ListLimit) error
	ListTenor() (res []model.ListLimit, err error)
	GetLimitCustomer(nik string) (res []model.LimitCustomer, err error)
	UpdateLimit(params request.UpdateLimit) error
}
