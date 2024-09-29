package repo

import (
	"kredit-plus/lib/db"
	"kredit-plus/service/model"
)

type transactionRepository struct{}

func NewTransactionRepository() TransactionRepository {
	return &transactionRepository{}
}

func (repo *transactionRepository) CreateTransaction(params *model.Transaction) error {
	query := `insert into
					transaction
						(
							nomor_kontrak,
							nik,
							otr,
							admin_fee,
							jumlah_cicilan,
							jumlah_bunga,
							nama_asset,
							created_at)
					values 
						( ? , ? , ? , ? , ? , ? , ? , now() )`

	if _, err := db.EngineSQL.Exec(query,
		params.NomorKontrak,
		params.Nik,
		params.OTR,
		params.AdminFee,
		params.JumlahCicilan,
		params.JumlahBunga,
		params.NamaAset); err != nil {
		return err
	}
	return nil
}

func (repo *transactionRepository) GetTransaction() (TransactionID int, err error) {
	query := `select count(*) from transaction`
	if _, err := db.EngineSQL.SQL(query).Get(&TransactionID); err != nil {
		return 0, err
	}

	return TransactionID, nil

}
