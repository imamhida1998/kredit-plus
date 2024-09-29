package usecase

import "kredit-plus/service/model/request"

type TransactionService interface {
	Transaction(params *request.RequestTransaction) error
	CreateTransaction(params *request.TransactionRequest) (Otp string, err error)
}
