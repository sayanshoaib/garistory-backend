package repository

import (
	"context"
	"garistroy-backend/entity"
)

type IServicingCenterRepository interface {
	CreateServicingCenter(
		ctx context.Context, center *entity.ServiceCenter) (*entity.ServiceCenter, error)
	GetServicingCenterByID(ctx context.Context, servicingCenterID string) (*entity.ServiceCenter, error)
	GetAllServicingCenter(ctx context.Context) ([]entity.ServiceCenter, int, error)
	ChangePassword(ctx context.Context, servicingCenterID string, password string) error
	UpdateServicingCenter(ctx context.Context, servicingCenter *entity.ServiceCenter, password bool) error
	DeleteServicingCenterByID(ctx context.Context, servicingCenterID string) error
}
