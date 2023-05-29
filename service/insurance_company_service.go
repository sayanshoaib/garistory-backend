package service

import (
	"context"
	"garistroy-backend/apimodels"
	"garistroy-backend/entity"
	"garistroy-backend/repository"
	"time"
)

type InsuranceCompanyService struct {
	InsuranceCompanyRepository *repository.InsuranceCompanyRepository
}

func NewInsuranceCompanyService(
	insuranceCompanyRepository *repository.InsuranceCompanyRepository) *InsuranceCompanyService {
	return &InsuranceCompanyService{
		InsuranceCompanyRepository: insuranceCompanyRepository,
	}
}

func (service *InsuranceCompanyService) CreateInsuranceCompany(
	ctx context.Context, insuranceCompanyDTO *apimodels.InsuranceCompanyDTO) (*entity.InsuranceCompany, error) {
	insuranceCompanyEntity := &entity.InsuranceCompany{
		InsuranceCompanyID: insuranceCompanyDTO.InsuranceCompanyID,
		Name:               insuranceCompanyDTO.Name,
		Email:              insuranceCompanyDTO.Email,
		Address: entity.Address{
			Name:      insuranceCompanyDTO.Address.Name,
			Street:    insuranceCompanyDTO.Address.Street,
			Zipcode:   insuranceCompanyDTO.Address.Zipcode,
			Country:   insuranceCompanyDTO.Address.Country,
			CreatedAt: time.Now(),
		},
		Password:         insuranceCompanyDTO.Password,
		AutoCoverage:     insuranceCompanyDTO.AutoCoverage,
		LifeCoverage:     insuranceCompanyDTO.LifeCoverage,
		LicenseNumber:    insuranceCompanyDTO.LicenseNumber,
		LicenseType:      insuranceCompanyDTO.LicenseType,
		ValidityPeriod:   insuranceCompanyDTO.ValidityPeriod,
		LicenseStatus:    insuranceCompanyDTO.LicenseStatus,
		IssuingAuthority: insuranceCompanyDTO.IssuingAuthority,
		CreatedAt:        time.Now(),
	}

	data, err := service.InsuranceCompanyRepository.CreateInsuranceCompany(ctx, insuranceCompanyEntity)
	if err != nil {
		return nil, err
	}

	return data, nil
}
