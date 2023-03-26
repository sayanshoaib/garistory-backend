package entity

import "github.com/uptrace/bun"

type ServicingRecord struct {
	bun.BaseModel             `bun:"servicing_record"`
	ID                        int      `json:"id" bun:"id,pk,autoincrement"`
	ServicingRecordID         string   `json:"recordID" bun:"record_id,unique"`
	Description               string   `json:"description" bun:"description"`
	CostOfService             int      `json:"costOfService" bun:"cost_of_service"`
	OdometerReading           int      `json:"odometerReading" bun:"odometer_reading"`
	RecommendedFutureServices []string `json:"recommendedFutureServices" bun:"recommended_future_services"`
	MechanicSignature         string   `json:"mechanicSignature" bun:"mechanic_signature"`
}
