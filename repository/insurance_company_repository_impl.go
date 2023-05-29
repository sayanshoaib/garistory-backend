package repository

import (
	"context"
	"garistroy-backend/entity"
	"github.com/labstack/gommon/log"
	"github.com/uptrace/bun"
)

type InsuranceCompanyRepository struct {
	Client            *bun.DB
	AddressRepository *AddressRepository
}

func NewInsuranceCompanyRepository(client *bun.DB, addressRepository *AddressRepository) *InsuranceCompanyRepository {
	return &InsuranceCompanyRepository{
		Client:            client,
		AddressRepository: addressRepository,
	}
}

func (repo *InsuranceCompanyRepository) CreateInsuranceCompany(
	ctx context.Context, insuranceCompany *entity.InsuranceCompany) (*entity.InsuranceCompany, error) {
	reqAddress := insuranceCompany.Address
	addressEntity, err := repo.AddressRepository.CreateAddress(ctx, &reqAddress)
	if err != nil {
		log.Info(err)
		return nil, err
	}

	addressID := addressEntity.ID
	insuranceCompany.AddressID = addressID

	var query *bun.InsertQuery
	query = repo.Client.NewInsert()

	_, err = query.
		Model(insuranceCompany).
		Value("password", "crypt(?, gen_salt('bf'))", insuranceCompany.Password).
		ExcludeColumn("updated_at", "deleted_at").
		Returning("id").
		Exec(ctx)

	if err != nil {
		log.Info(err)
		return nil, err
	}

	return insuranceCompany, nil
}
