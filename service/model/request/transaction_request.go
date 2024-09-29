package request

type TransactionRequest struct {
	NIK       string `json:"nik"`
	HargaAset int    `json:"harga_aset"`
	NamaAset  string `json:"nama_aset"`
	Tenor     int    `json:"tenor"`
}
