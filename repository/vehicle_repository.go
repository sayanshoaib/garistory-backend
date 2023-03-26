package repository

import (
	"context"
	"database/sql"
	"garistroy-backend/entity"
)

type IVehicleRepository interface {
	CreateVehicle(ctx context.Context, vehicle *entity.Vehicle) (sql.Result, error)
	GetVehicleByVin(ctx context.Context, vin string) (*entity.Vehicle, error)
	GetAllVehicle(ctx context.Context) ([]entity.Vehicle, error)
	UpdateVehicleByVin(ctx context.Context, vin string, vehicleToUpdate *entity.Vehicle) (sql.Result, error)
	DeleteVehicleByVin(ctx context.Context, vin string) (sql.Result, error)
}
