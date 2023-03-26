package repository

import (
	"context"
	"database/sql"
	"garistroy-backend/entity"
	"github.com/uptrace/bun"
)

type VehicleRepository struct {
	Client *bun.DB
}

func NewVehicleRepository(client *bun.DB) *VehicleRepository {
	return &VehicleRepository{
		Client: client,
	}
}

func (repo *VehicleRepository) CreateVehicle(ctx context.Context, vehicle *entity.Vehicle) (sql.Result, error) {
	resp, err := repo.Client.NewInsert().Model(vehicle).
		ExcludeColumn("updated_at", "deleted_at").
		Exec(ctx)

	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (repo *VehicleRepository) GetVehicleByVin(ctx context.Context, vin string) (*entity.Vehicle, error) {
	vehicle := &entity.Vehicle{}
	err := repo.Client.NewSelect().
		Model(vehicle).
		Where("vehicle_id = ?", vin).
		Scan(ctx)

	if err != nil {
		return nil, err
	}

	return vehicle, nil
}

func (repo *VehicleRepository) GetAllVehicle(ctx context.Context) ([]entity.Vehicle, error) {
	var vehicles []entity.Vehicle
	_, err := repo.Client.NewSelect().Model(&vehicles).Limit(20).ScanAndCount(ctx)

	if err != nil {
		return nil, err
	}

	return vehicles, nil
}

func (repo *VehicleRepository) UpdateVehicleByVin(
	ctx context.Context, vin string, vehicleToUpdate *entity.Vehicle) (sql.Result, error) {
	res, err := repo.Client.NewUpdate().
		Model(vehicleToUpdate).
		Column("vehicle_id").
		Where("vehicle_id = ?", vin).
		Exec(ctx)

	if err != nil {
		return nil, err
	}

	return res, err
}

func (repo *VehicleRepository) DeleteVehicleByVin(ctx context.Context, vin string) (sql.Result, error) {
	user := new(entity.Vehicle)
	res, err := repo.Client.NewDelete().Model(user).Where("vehicle_id = ?", vin).Exec(ctx)
	if err != nil {
		return nil, err
	}

	return res, nil
}
