package entity

type AccountType string

const (
	ADMIN           AccountType = "admin"
	ServicingCenter AccountType = "servicing_center"
	InsuranceComp   AccountType = "insurance_company"
)
