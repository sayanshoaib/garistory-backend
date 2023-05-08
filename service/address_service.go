package service

import (
	"context"
	"garistroy-backend/apimodels"
	"garistroy-backend/entity"
	"garistroy-backend/repository"
	"time"
)

type AddressService struct {
	AddressRepository *repository.AddressRepository
}

func NewAddressService(addressRepo *repository.AddressRepository) *AddressService {
	return &AddressService{
		AddressRepository: addressRepo,
	}
}

func (service *AddressService) CreateAddress(ctx context.Context, address *apimodels.Address) (*apimodels.RespAddress, error) {
	addressEntity := &entity.Address{
		Name:      address.Name,
		Street:    address.Street,
		Zipcode:   address.Zipcode,
		Country:   address.Country,
		CreatedAt: time.Now(),
	}
	entity, err := service.AddressRepository.CreateAddress(ctx, addressEntity)
	if err != nil {
		return nil, err
	}

	resp := &apimodels.RespAddress{
		Name:      entity.Name,
		Street:    entity.Street,
		Zipcode:   entity.Zipcode,
		Country:   entity.Country,
		CreatedAt: entity.CreatedAt,
		UpdatedAt: entity.UpdatedAt,
	}

	return resp, nil
}

func (service *AddressService) GetAddressByID(ctx context.Context, id string) (*apimodels.RespAddress, error) {
	if id == "" {
		return nil, apimodels.ErrNoRows
	}

	entity, err := service.AddressRepository.GetAddressByID(ctx, id)
	if err != nil {
		return nil, apimodels.ErrNoDataMatched
	}

	resp := &apimodels.RespAddress{
		Name:      entity.Name,
		Street:    entity.Street,
		Zipcode:   entity.Zipcode,
		Country:   entity.Country,
		CreatedAt: entity.CreatedAt,
		UpdatedAt: entity.UpdatedAt,
	}
	return resp, nil
}

func (service *AddressService) GetAddressByZipcode(ctx context.Context, zipcode string) (*apimodels.RespAddress, error) {
	if zipcode == "" {
		return nil, apimodels.ErrNoRows
	}

	entity, err := service.AddressRepository.GetAddressByZipcode(ctx, zipcode)
	if err != nil {
		return nil, apimodels.ErrNoDataMatched
	}

	resp := &apimodels.RespAddress{
		Name:      entity.Name,
		Street:    entity.Street,
		Zipcode:   entity.Zipcode,
		Country:   entity.Country,
		CreatedAt: entity.CreatedAt,
		UpdatedAt: entity.UpdatedAt,
	}
	return resp, nil
}

func (service *AddressService) GetAllAddress(ctx context.Context) ([]*apimodels.RespAddress, int, error) {
	var addresses []*apimodels.RespAddress
	entities, count, err := service.AddressRepository.GetAllAddress(ctx)
	if err != nil {
		return nil, 0, apimodels.ErrNoRows
	}

	for _, entity := range entities {
		resp := &apimodels.RespAddress{
			Name:      entity.Name,
			Street:    entity.Street,
			Zipcode:   entity.Zipcode,
			Country:   entity.Country,
			CreatedAt: entity.CreatedAt,
			UpdatedAt: entity.UpdatedAt,
		}
		addresses = append(addresses, resp)
	}

	return addresses, count, nil
}

func (service *AddressService) UpdateAddressByID(ctx context.Context, id int, newAddress *apimodels.Address) (*apimodels.RespAddress, error) {
	if id == 0 {
		return nil, apimodels.ErrNoRows
	}

	addressEntity := &entity.Address{
		ID:        id,
		Name:      newAddress.Name,
		Street:    newAddress.Street,
		Zipcode:   newAddress.Zipcode,
		Country:   newAddress.Country,
		UpdatedAt: time.Now(),
	}
	entity, err := service.AddressRepository.UpdateAddressByID(ctx, addressEntity)
	if err != nil {
		return nil, err
	}

	resp := &apimodels.RespAddress{
		Name:      entity.Name,
		Street:    entity.Street,
		Zipcode:   entity.Zipcode,
		Country:   entity.Country,
		CreatedAt: entity.CreatedAt,
		UpdatedAt: entity.UpdatedAt,
	}

	return resp, nil
}

func (service *AddressService) DeleteAddressByID(ctx context.Context, id int) (bool, error) {
	if id == 0 {
		return false, apimodels.ErrNoRows
	}

	ok, err := service.AddressRepository.DeleteAddressByID(ctx, id)
	if err != nil {
		return false, apimodels.ErrNoDataMatched
	}

	return ok, nil
}
