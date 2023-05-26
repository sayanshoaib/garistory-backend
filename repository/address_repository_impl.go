package repository

import (
	"context"
	"garistroy-backend/entity"
	"github.com/uptrace/bun"
	"log"
)

type AddressRepository struct {
	Client *bun.DB
}

func NewAddressRepository(client *bun.DB) *AddressRepository {
	return &AddressRepository{
		Client: client,
	}
}

func (repo *AddressRepository) CreateAddress(ctx context.Context, address *entity.Address) (*entity.Address, error) {
	_, err := repo.Client.
		NewInsert().
		Model(address).
		ExcludeColumn("updated_at", "deleted_at").
		Returning("*").
		Exec(ctx)
	if err != nil {
		return nil, err
	}

	return address, nil
}
func (repo *AddressRepository) GetAddressByID(ctx context.Context, addressID string) (*entity.Address, error) {
	address := &entity.Address{}
	err := repo.Client.
		NewSelect().
		Model(address).
		Where("id = ?", addressID).
		Where("deleted_at IS NULL").
		Scan(ctx)
	if err != nil {
		log.Print(err)
		return nil, err
	}

	return address, nil
}

func (repo *AddressRepository) GetAddressByZipcode(ctx context.Context, zipcode string) (*entity.Address, error) {
	address := &entity.Address{}
	err := repo.Client.
		NewSelect().
		Model(address).
		Where("zipcode = ?", zipcode).
		Where("deleted_at IS NULL").
		Scan(ctx)
	if err != nil {
		log.Print(err)
		return nil, err
	}

	return address, nil
}

func (repo *AddressRepository) GetAllAddress(ctx context.Context) ([]entity.Address, int, error) {
	var addresses []entity.Address
	count, err := repo.Client.
		NewSelect().
		Model(&addresses).
		Where("deleted_at IS NULL").
		Limit(20).
		ScanAndCount(ctx)
	if err != nil {
		return nil, 0, err
	}

	return addresses, count, nil
}
func (repo *AddressRepository) UpdateAddressByID(
	ctx context.Context, address *entity.Address) (*entity.Address, error) {
	_, err := repo.Client.
		NewUpdate().
		Model(address).
		WherePK().
		Exec(ctx)
	if err != nil {
		return nil, err
	}
	return address, nil
}

func (repo *AddressRepository) DeleteAddressByID(ctx context.Context, addressID int) (bool, error) {
	address := &entity.Address{}
	_, err := repo.Client.
		NewDelete().
		Model(address).
		Where("id = ?", addressID).
		Exec(ctx)
	if err != nil {
		return false, err
	}
	return true, nil
}
