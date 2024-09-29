package repo

import (
	"kredit-plus/lib/db"
	"kredit-plus/service/model"
	"kredit-plus/service/model/request"
	"strconv"
)

type customerRepository struct {
}

func NewUserRepository() CustomerRepository {
	return &customerRepository{}
}

func (repo *customerRepository) InputCustomer(params *model.InputCustomer) error {
	query := `insert into 
					Customer 
						(
							nik ,
							full_name ,
							legal_name ,
							place_of_birth,
							date_of_birth, 
							salary, 
							ktp_image,
							selfie_image,
							created_at)
					values (
						? , ? , ? , ? , ? , ? , ? , ? , now() 
					)`
	if _, err := db.EngineSQL.Exec(query,
		params.Nik,
		params.FullName,
		params.FullName,
		params.LegalName,
		params.DateOfBirth,
		params.Salary,
		params.KtpImage,
		params.SelfieImage,
	); err != nil {
		return err
	}

	return nil
}

func (repo *customerRepository) InputLimit(params *model.LimitCustomer) error {
	query := `insert into 
					customer_limit 
						(
							nik ,
							tenor ,
							limits ,
							created_at)
					values (
						? , ? , ? , now() 
					)`
	if _, err := db.EngineSQL.Exec(query,
		params.Nik,
		params.Tenor,
		params.Limit,
	); err != nil {
		return err
	}

	return nil
}

func (repo *customerRepository) InsertTenor(params *model.ListLimit) error {
	query := `insert into 
					tenor_limit 
						(
							tenor ,
							created_at)
					values (
						 ? , now() 
					)`
	if _, err := db.EngineSQL.Exec(query,
		params.Tenor,
	); err != nil {
		return err
	}

	return nil
}

func (repo *customerRepository) ListTenor() (res []model.ListLimit, err error) {
	query := `select tenor from tenor_limit`
	resLimit, err := db.EngineSQL.Query(query)
	if err != nil {
		return nil, err
	}

	for _, tenor := range resLimit {
		dataTenor := model.ListLimit{
			Tenor: string(tenor["tenor"]),
		}
		res = append(res, dataTenor)
	}

	return res, nil
}

func (repo *customerRepository) GetLimitCustomer(nik string) (res []model.LimitCustomer, err error) {
	query := `	select 
					tenor,
					limits 
				from
					customer_limit
				where
					nik = ?`
	resLimit, err := db.EngineSQL.Query(query,
		nik,
	)
	if err != nil {
		return nil, err
	}

	for _, Limit := range resLimit {
		tenor, _ := strconv.Atoi(string(Limit["tenor"]))
		limit, _ := strconv.Atoi(string(Limit["limits"]))
		req := model.LimitCustomer{
			Nik:   nik,
			Tenor: tenor,
			Limit: limit,
		}
		res = append(res, req)
	}

	return res, nil
}

func (repo *customerRepository) UpdateLimit(params request.UpdateLimit) error {
	query := `update 
					customer_limit 
				set	
					limits = ?,
					updated_at = now()
				where
					tenor = ? 
				and
					nik = ?`
	if _, err := db.EngineSQL.Exec(query, params.Limit, params.Tenor, params.Nik); err != nil {
		return err
	}
	return nil
}
