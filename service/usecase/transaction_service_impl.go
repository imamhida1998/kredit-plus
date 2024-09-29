package usecase

import (
	"encoding/json"
	"errors"
	"fmt"
	"kredit-plus/service/model"
	"kredit-plus/service/model/request"
	"kredit-plus/service/repo"
	"kredit-plus/service/utils"
	"strconv"
	"sync"
	"time"
)

type transactionService struct {
	repoUser        repo.CustomerRepository
	repoTransaction repo.TransactionRepository
	redis           repo.RedisRepository
	Mutex           sync.Mutex
}

func NewTransaction(repoUser repo.CustomerRepository, repoTransaction repo.TransactionRepository, redis repo.RedisRepository) TransactionService {
	return &transactionService{repoUser: repoUser, repoTransaction: repoTransaction, redis: redis}
}

func (usecase *transactionService) Transaction(params *request.RequestTransaction) error {
	usecase.Mutex.Lock()
	defer usecase.Mutex.Unlock()

	timeNow := time.Now().Format("01/02/2006")
	KeyOTP := params.Nik + "#OTP"
	cekData := usecase.redis.GetValue(KeyOTP)
	if cekData.Value != "" {
		count, _ := strconv.Atoi(cekData.Value)
		if count > 1 {
			Gettl := usecase.redis.GetTtl(KeyOTP)
			ttl := utils.ShortDur(Gettl)
			message := fmt.Sprintf("Request ditolak, Mohon kembali pada pukul %s", ttl)
			return errors.New(message)
		} else {
			var dataTransaction request.TransactionRequest
			KeyTransaction := params.Nik + "#Transaction#" + params.Otp

			Data := usecase.redis.GetValue(KeyTransaction)

			err := json.Unmarshal([]byte(Data.Value), &dataTransaction)
			if err != nil {
				return err
			}

			Id, err := usecase.repoTransaction.GetTransaction()
			if err != nil {
				return err
			}

			suratKontrak := fmt.Sprintf("%d/%s", Id, timeNow)
			GetLimit, err := usecase.repoUser.GetLimitCustomer(dataTransaction.NIK)
			if err != nil {
				return err
			}

			limitMapping := make(map[int]int)
			for _, limit := range GetLimit {
				limitMapping[limit.Tenor] = limit.Limit
			}

			limit := limitMapping[dataTransaction.Tenor]
			if dataTransaction.HargaAset > limit {
				return errors.New("harga melebihi limit")
			}
			countLimit := limit - dataTransaction.HargaAset
			UpdateLimit := request.UpdateLimit{
				Nik:   dataTransaction.NIK,
				Tenor: dataTransaction.Tenor,
				Limit: countLimit,
			}

			err = usecase.repoUser.UpdateLimit(UpdateLimit)
			if err != nil {
				return err
			}

			requestTransaction := &model.Transaction{
				NomorKontrak:  suratKontrak,
				Nik:           dataTransaction.NIK,
				OTR:           dataTransaction.HargaAset,
				AdminFee:      2000,
				JumlahCicilan: dataTransaction.Tenor,
				JumlahBunga:   2000,
				NamaAset:      dataTransaction.NamaAset,
			}

			err = usecase.repoTransaction.CreateTransaction(requestTransaction)
			if err != nil {
				return err
			}
		}
	} else {
		return errors.New("anda tidak mengajukan kredit")
	}
	return nil

}

// create transaction sebagai preorder sblm transaction
func (usecase *transactionService) CreateTransaction(params *request.TransactionRequest) (Otp string, err error) {

	otp := utils.RandomNumber(6)
	NoOtp := strconv.Itoa(otp)
	KeyOTP := params.NIK + "#OTP"
	getData := usecase.redis.GetValue(KeyOTP)
	count, _ := strconv.Atoi(getData.Value)

	if count != 0 {
		count += 1
		req := &model.RedisStoreRequest{
			KeyValue: KeyOTP,
			Value:    strconv.Itoa(count),
			Lifetime: 300,
		}
		err := usecase.redis.StoreValue(req)
		if err != nil {
			return "", err
		}
	} else {
		KeyTransaction := params.NIK + "#Transaction#" + NoOtp
		req := &model.RedisStoreRequest{
			KeyValue: KeyOTP,
			Value:    "1",
			Lifetime: 300,
		}
		data, err := json.Marshal(params)
		if err != nil {
			return "", err
		}
		reqTx := &model.RedisStoreRequest{
			KeyValue: KeyTransaction,
			Value:    string(data),
			Lifetime: 300,
		}

		err = usecase.redis.StoreValue(req)
		if err != nil {
			return "", err
		}
		err = usecase.redis.StoreValue(reqTx)
		if err != nil {
			return "", err
		}

	}
	return NoOtp, nil

}
