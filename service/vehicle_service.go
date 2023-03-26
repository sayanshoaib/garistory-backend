package service

import (
	"context"
	"garistroy-backend/apimodels"
	"garistroy-backend/entity"
	"garistroy-backend/repository"
	"time"
)

type VehicleService struct {
	VehicleRepository repository.IVehicleRepository
}

func NewVehicleRepository(vehicleRepository *repository.VehicleRepository) *VehicleService {
	return &VehicleService{
		VehicleRepository: vehicleRepository,
	}
}

func (service *VehicleService) CreateVehicle(
	ctx context.Context, reqVehicle *apimodels.ReqVehicleCreate) (*apimodels.RespVehicleCreate, error) {
	vehicle := &entity.Vehicle{
		VehicleID:           reqVehicle.VehicleID,
		Make:                reqVehicle.Make,
		Model:               reqVehicle.Model,
		Year:                reqVehicle.Year,
		BodyType:            reqVehicle.BodyType,
		InitialRegistration: reqVehicle.InitialRegistration,
		RegistrationNumber:  reqVehicle.RegistrationNumber,
		Fuel:                reqVehicle.Fuel,
		Mileage:             reqVehicle.Mileage,
		Drive:               reqVehicle.Drive,
		Color:               reqVehicle.Color,
		Transmission:        reqVehicle.Transmission,
		Price:               reqVehicle.Price,
		ImageURL:            reqVehicle.ImageURL,
		CreatedAt:           time.Now(),
	}
	_, err := service.VehicleRepository.CreateVehicle(ctx, vehicle)
	if err != nil {
		return nil, err
	}

	resp := &apimodels.RespVehicleCreate{
		VehicleID:           reqVehicle.VehicleID,
		Make:                reqVehicle.Make,
		Model:               reqVehicle.Model,
		Year:                reqVehicle.Year,
		BodyType:            reqVehicle.BodyType,
		InitialRegistration: reqVehicle.InitialRegistration,
		RegistrationNumber:  reqVehicle.RegistrationNumber,
		Fuel:                reqVehicle.Fuel,
		Mileage:             reqVehicle.Mileage,
		Drive:               reqVehicle.Drive,
		Color:               reqVehicle.Color,
		Transmission:        reqVehicle.Transmission,
		Price:               reqVehicle.Price,
		ImageURL:            reqVehicle.ImageURL,
		CreatedAt:           time.Now(),
	}

	return resp, nil
}

func (service *VehicleService) GetVehicleByVin(ctx context.Context, vin string) (*apimodels.RespVehicle, error) {
	vehicle, err := service.VehicleRepository.GetVehicleByVin(ctx, vin)
	if err != nil {
		return nil, err
	}

	return &apimodels.RespVehicle{
		VehicleID:           vehicle.VehicleID,
		Make:                vehicle.Make,
		Model:               vehicle.Model,
		Year:                vehicle.Year,
		BodyType:            vehicle.BodyType,
		InitialRegistration: vehicle.InitialRegistration,
		RegistrationNumber:  vehicle.RegistrationNumber,
		Fuel:                vehicle.Fuel,
		Mileage:             vehicle.Mileage,
		Drive:               vehicle.Drive,
		Color:               vehicle.Color,
		Transmission:        vehicle.Transmission,
		Price:               vehicle.Price,
		ImageURL:            vehicle.ImageURL,
	}, nil
}

func (service *VehicleService) GetAllVehicle(ctx context.Context) ([]*apimodels.RespVehicle, error) {
	vehicles, err := service.VehicleRepository.GetAllVehicle(ctx)
	if err != nil {
		return nil, err
	}

	var respVehicles []*apimodels.RespVehicle

	for _, vehicle := range vehicles {
		respVehicle := &apimodels.RespVehicle{
			VehicleID:           vehicle.VehicleID,
			Make:                vehicle.Make,
			Model:               vehicle.Model,
			Year:                vehicle.Year,
			BodyType:            vehicle.BodyType,
			InitialRegistration: vehicle.InitialRegistration,
			RegistrationNumber:  vehicle.RegistrationNumber,
			Fuel:                vehicle.Fuel,
			Mileage:             vehicle.Mileage,
			Drive:               vehicle.Drive,
			Color:               vehicle.Color,
			Transmission:        vehicle.Transmission,
			Price:               vehicle.Price,
			ImageURL:            vehicle.ImageURL,
		}

		respVehicles = append(respVehicles, respVehicle)
	}

	return respVehicles, nil
}

func (service *VehicleService) UpdateVehicleByVin(
	ctx context.Context, vin string, vehicleUpdate *apimodels.ReqVehicleUpdate) (*apimodels.RespVehicleUpdate, error) {
	vehicleToUpdate := &entity.Vehicle{
		VehicleID:           vehicleUpdate.VehicleID,
		Make:                vehicleUpdate.Make,
		Model:               vehicleUpdate.Model,
		Year:                vehicleUpdate.Year,
		BodyType:            vehicleUpdate.BodyType,
		InitialRegistration: vehicleUpdate.InitialRegistration,
		RegistrationNumber:  vehicleUpdate.RegistrationNumber,
		Fuel:                vehicleUpdate.Fuel,
		Mileage:             vehicleUpdate.Mileage,
		Drive:               vehicleUpdate.Drive,
		Color:               vehicleUpdate.Color,
		Transmission:        vehicleUpdate.Transmission,
		Price:               vehicleUpdate.Price,
		ImageURL:            vehicleUpdate.ImageURL,
		UpdatedAt:           time.Now(),
	}

	_, err := service.VehicleRepository.UpdateVehicleByVin(ctx, vin, vehicleToUpdate)
	if err != nil {
		return nil, err
	}

	resp := &apimodels.RespVehicleUpdate{
		VehicleID:           vehicleUpdate.VehicleID,
		Make:                vehicleUpdate.Make,
		Model:               vehicleUpdate.Model,
		Year:                vehicleUpdate.Year,
		BodyType:            vehicleUpdate.BodyType,
		InitialRegistration: vehicleUpdate.InitialRegistration,
		RegistrationNumber:  vehicleUpdate.RegistrationNumber,
		Fuel:                vehicleUpdate.Fuel,
		Mileage:             vehicleUpdate.Mileage,
		Drive:               vehicleToUpdate.Drive,
		Color:               vehicleToUpdate.Color,
		Transmission:        vehicleToUpdate.Transmission,
		Price:               vehicleToUpdate.Price,
		ImageURL:            vehicleToUpdate.ImageURL,
		UpdatedAt:           vehicleToUpdate.UpdatedAt,
	}

	return resp, nil
}

func (service *VehicleService) DeleteVehicleByVin(ctx context.Context, vin string) (bool, error) {
	_, err := service.VehicleRepository.DeleteVehicleByVin(ctx, vin)
	if err != nil {
		return false, err
	}

	return true, nil
}
