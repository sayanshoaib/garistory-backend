package repository

import (
	"context"
	"garistroy-backend/entity"
)

type IInsuranceCompanyRepository interface {
	CreateInsuranceCompany(
		ctx context.Context, center *entity.InsuranceCompany) (*entity.InsuranceCompany, error)
	//GetInsuranceCompanyByID(ctx context.Context, InsuranceCompanyID string) (*entity.InsuranceCompany, error)
	//GetAllInsuranceCompany(ctx context.Context) ([]entity.InsuranceCompany, int, error)
	//ChangePassword(ctx context.Context, insuranceCompanyID string, password string) error
	//UpdateInsuranceCompany(
	//	ctx context.Context, InsuranceCompany *entity.InsuranceCompany) (*entity.InsuranceCompany, error)
	//DeleteInsuranceCompanyByID(ctx context.Context, InsuranceCompanyID string) error
}
