package entity

import "time"

type InsuranceCompany struct {
	ID                 int       `json:"id" bun:",pk,autoincrement"`
	InsuranceCompanyID string    `json:"service_center_id" bun:",notnull"`
	Name               string    `json:"name" bun:",notnull"`
	Email              string    `json:"email" bun:",notnull"`
	AddressID          int       `bun:"bun:address_id"`
	Address            Address   `json:"address" bun:"address,rel:belongs-to,join:address_id=id"`
	Password           string    `json:"password" bun:",notnull"`
	AutoCoverage       bool      `json:"auto_coverage" bun:",notnull"`
	LifeCoverage       bool      `json:"life_coverage" bun:",notnull"`
	LicenseNumber      string    `json:"license_number" bun:",notnull"`
	LicenseType        string    `json:"license_type" bun:",notnull"`
	ValidityPeriod     string    `json:"validity_period" bun:",notnull"`
	LicenseStatus      string    `json:"license_status" bun:",notnull"`
	IssuingAuthority   string    `json:"issuing_authority" bun:",notnull"`
	CreatedAt          time.Time `json:"created_at" bun:",nullzero,default:current_timestamp"`
	UpdatedAt          time.Time `json:"updated_at" bun:",nullzero,default:current_timestamp"`
	DeletedAt          time.Time `json:"deleted_at" bun:",soft_delete,nullzero"`
}
