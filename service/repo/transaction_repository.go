package repo

import "kredit-plus/service/model"

type TransactionRepository interface {
	CreateTransaction(params *model.Transaction) error
	GetTransaction() (TransactionID int, err error)
}
