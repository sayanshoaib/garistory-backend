package repository

import (
	"context"
	"database/sql"
	"garistroy-backend/entity"
	"github.com/uptrace/bun"
	"log"
)

type ServicingCenterRepository struct {
	Client            *bun.DB
	AddressRepository *AddressRepository
}

func NewServicingCenterRepository(client *bun.DB, addressRepository *AddressRepository) *ServicingCenterRepository {
	return &ServicingCenterRepository{
		Client:            client,
		AddressRepository: addressRepository,
	}
}

func (repo *ServicingCenterRepository) RegisterServicingCenter(
	ctx context.Context, center *entity.ServiceCenter) (*entity.ServiceCenter, error) {
	reqAddress := center.Address
	addressEntity, err := repo.AddressRepository.CreateAddress(ctx, &reqAddress)
	if err != nil {
		return nil, err
	}

	addressID := addressEntity.ID
	center.AddressID = addressID

	var query *bun.InsertQuery
	query = repo.Client.NewInsert()

	_, err = query.
		Model(center).
		Value("password", "crypt(?, gen_salt('bf'))", center.Password).
		ExcludeColumn("updated_at", "deleted_at").
		Returning("id").
		Exec(ctx)

	if err != nil {
		return nil, err
	}

	return center, nil
}

func (repo *ServicingCenterRepository) GetServicingCenterByID(
	ctx context.Context, servicingCenterID string) (*entity.ServiceCenter, error) {
	//var servicingCenter *entity.ServiceCenter
	servicingCenter := &entity.ServiceCenter{}

	err := repo.Client.NewSelect().
		Model(servicingCenter).
		Relation("Address").
		Where("service_center_id = ?", servicingCenterID).
		Scan(ctx)
	if err != nil {
		return nil, err
	}

	return servicingCenter, nil
}

func (repo *ServicingCenterRepository) GetAllServicingCenter(ctx context.Context) ([]entity.ServiceCenter, int, error) {
	var servicingCenters []entity.ServiceCenter
	count, err := repo.Client.NewSelect().
		Model(&servicingCenters).
		Relation("Address").
		Limit(20).
		ScanAndCount(ctx)
	if err != nil {
		log.Println("inside repo: ", err)
		return nil, 0, err
	}

	return servicingCenters, count, nil
}

func (repo *ServicingCenterRepository) ChangePassword(ctx context.Context, servicingCenterID string, password string) error {
	_, err := repo.Client.NewUpdate().
		Model((*entity.ServiceCenter)(nil)).
		Set("updated_at = NOW(), password =crypt(?, gen_salt('bf'))", password).
		Where("id = ?", servicingCenterID).
		Exec(ctx)
	return err
}

func (repo *ServicingCenterRepository) UpdateServicingCenter(ctx context.Context, servicingCenter *entity.ServiceCenter, updatePassword bool) error {
	update := repo.Client.NewUpdate().
		Model(servicingCenter).
		Value("updated_at", "NOW()").
		ExcludeColumn("created_at")
	if !updatePassword {
		update.ExcludeColumn("password")
	} else {
		update.Value("password", "crypt(?, gen_salt('bf'))", servicingCenter.Password)
	}
	_, err := update.
		Where("service_center_id = ?", servicingCenter.ServiceCenterID).
		Exec(ctx)
	return err
}

func (repo *ServicingCenterRepository) DeleteServicingCenterByID(ctx context.Context, servicingCenterID string) error {
	_, err := repo.Client.NewDelete().
		Model((*entity.ServiceCenter)(nil)).
		Where("service_center_id = ?", servicingCenterID).
		Exec(ctx)
	return err
}

func (repo *ServicingCenterRepository) BeginTx(ctx context.Context) (bun.Tx, error) {
	return repo.Client.BeginTx(ctx, &sql.TxOptions{})
}
