package apimodels

type InsuranceCompanyDTO struct {
	InsuranceCompanyID string  `json:"insurance_company_id" binding:"required"`
	Name               string  `json:"name" binding:"required"`
	Email              string  `json:"email" binding:"required,email"`
	Address            Address `json:"address"`
	Password           string  `json:"password" binding:"required"`
	AutoCoverage       bool    `json:"auto_coverage" binding:"required"`
	LifeCoverage       bool    `json:"life_coverage" binding:"required"`
	LicenseNumber      string  `json:"license_number" binding:"required"`
	LicenseType        string  `json:"license_type" binding:"required"`
	ValidityPeriod     string  `json:"validity_period" binding:"required"`
	LicenseStatus      string  `json:"license_status" binding:"required"`
	IssuingAuthority   string  `json:"issuing_authority" binding:"required"`
}
