package service

import (
	"context"
	"garistroy-backend/apimodels"
	"garistroy-backend/entity"
	"garistroy-backend/repository"
	"log"
	"time"
)

type ServicingCenterService struct {
	ServicingCenterRepository *repository.ServicingCenterRepository
}

func NewServicingCenterService(repo *repository.ServicingCenterRepository) *ServicingCenterService {
	return &ServicingCenterService{
		ServicingCenterRepository: repo,
	}
}

func (service *ServicingCenterService) RegisterServicingCenter(
	ctx context.Context, center *apimodels.ReqServicingCenter) (*apimodels.RespServiceCenter, error) {
	err := center.Validate()
	if err != nil {
		return nil, err
	}

	servicingCenterEntity := &entity.ServiceCenter{
		ServiceCenterID: center.ServiceCenterID,
		Name:            center.Name,
		Email:           center.Email,
		Address: entity.Address{
			Name:      center.Address.Name,
			Street:    center.Address.Street,
			Zipcode:   center.Address.Zipcode,
			Country:   center.Address.Country,
			CreatedAt: time.Now(),
		},
		Password:  center.Password,
		CreatedAt: time.Now(),
	}

	respEntity, err := service.ServicingCenterRepository.RegisterServicingCenter(ctx, servicingCenterEntity)
	if err != nil {
		return nil, err
	}

	respServicingCenter := &apimodels.RespServiceCenter{
		ServiceCenterID: respEntity.ServiceCenterID,
		Name:            respEntity.Name,
		Email:           respEntity.Email,
		Address: &apimodels.Address{
			Name:    respEntity.Address.Name,
			Street:  respEntity.Address.Street,
			Zipcode: respEntity.Address.Zipcode,
			Country: respEntity.Address.Country,
		},
	}

	return respServicingCenter, nil
}

func (service *ServicingCenterService) GetServicingCenterByID(
	ctx context.Context, servicingCenterID string) (*apimodels.RespServiceCenter, error) {
	if servicingCenterID == "" {
		return nil, apimodels.ValidationError{
			Errors: map[string][]string{
				"error": []string{
					"provide valid service center ID",
				},
			},
		}
	}

	respEntity, err := service.ServicingCenterRepository.GetServicingCenterByID(ctx, servicingCenterID)
	if err != nil {
		return nil, apimodels.ErrNoDataMatched
	}

	resp := &apimodels.RespServiceCenter{
		ServiceCenterID: respEntity.ServiceCenterID,
		Name:            respEntity.Name,
		Email:           respEntity.Email,
		Address: &apimodels.Address{
			Name:    respEntity.Address.Name,
			Street:  respEntity.Address.Street,
			Zipcode: respEntity.Address.Zipcode,
			Country: respEntity.Address.Country,
		},
	}

	return resp, nil
}

func (service *ServicingCenterService) GetAllServicingCenter(
	ctx context.Context) ([]*apimodels.RespServiceCenter, int, error) {
	respEntities, count, err := service.ServicingCenterRepository.GetAllServicingCenter(ctx)
	if err != nil {
		log.Print(err)
		return nil, 0, apimodels.ErrNoDataMatched
	}

	var respServicingCenters []*apimodels.RespServiceCenter
	for _, respEntity := range respEntities {
		resp := &apimodels.RespServiceCenter{
			ServiceCenterID: respEntity.ServiceCenterID,
			Name:            respEntity.Name,
			Email:           respEntity.Email,
			Address: &apimodels.Address{
				Name:    respEntity.Address.Name,
				Street:  respEntity.Address.Street,
				Zipcode: respEntity.Address.Zipcode,
				Country: respEntity.Address.Country,
			},
		}

		respServicingCenters = append(respServicingCenters, resp)
	}

	return respServicingCenters, count, nil
}

func (service *ServicingCenterService) UpdateServicingCenterByID(
	ctx context.Context,
	center *apimodels.ReqServicingCenter) (*apimodels.RespServiceCenter, error) {

	err := center.Validate()
	if err != nil {
		return nil, err
	}

	if center.ServiceCenterID == "" {
		return nil, apimodels.ValidationError{
			Errors: map[string][]string{
				"error": []string{
					"provide valid service center ID",
				},
			},
		}
	}

	servicingCenterEntity := &entity.ServiceCenter{
		ServiceCenterID: center.ServiceCenterID,
		Name:            center.Name,
		Email:           center.Email,
		Address: entity.Address{
			Name:      center.Address.Name,
			Street:    center.Address.Street,
			Zipcode:   center.Address.Zipcode,
			Country:   center.Address.Country,
			UpdatedAt: time.Now(),
		},
		Password:  center.Password,
		UpdatedAt: time.Now(),
	}

	err = service.ServicingCenterRepository.UpdateServicingCenter(ctx, servicingCenterEntity, false)
	if err != nil {
		return nil, apimodels.ErrNoDataMatched
	}

	respServicingCenter := &apimodels.RespServiceCenter{
		ServiceCenterID: servicingCenterEntity.ServiceCenterID,
		Name:            servicingCenterEntity.Name,
		Email:           servicingCenterEntity.Email,
		Address: &apimodels.Address{
			Name:    servicingCenterEntity.Address.Name,
			Street:  servicingCenterEntity.Address.Street,
			Zipcode: servicingCenterEntity.Address.Zipcode,
			Country: servicingCenterEntity.Address.Country,
		},
	}

	return respServicingCenter, nil
}

func (service *ServicingCenterService) DeleteServiceCenterByID(ctx context.Context, servicingCenterID string) (bool, error) {
	if servicingCenterID == "" {
		return false, apimodels.ValidationError{
			Errors: map[string][]string{
				"error": {
					"invalid service-center ID",
				},
			},
		}
	}

	err := service.ServicingCenterRepository.DeleteServicingCenterByID(ctx, servicingCenterID)
	if err != nil {
		return false, apimodels.ErrNoDataMatched
	}

	return true, nil
}
