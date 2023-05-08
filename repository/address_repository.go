package repository

import (
	"context"
	"garistroy-backend/entity"
)

type IAddressRepository interface {
	CreateAddress(ctx context.Context, address *entity.Address) (*entity.Address, error)
	GetAddressByID(ctx context.Context, addressID string) (*entity.Address, error)
	GetAddressByZipcode(ctx context.Context, zipcode string) (*entity.Address, error)
	GetAllAddress(ctx context.Context) ([]entity.Address, int, error)
	UpdateAddressByID(ctx context.Context, address *entity.Address) (*entity.Address, error)
	DeleteAddressByID(ctx context.Context, addressID int) (bool, error)
}
