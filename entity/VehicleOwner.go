package entity

import (
	"github.com/uptrace/bun"
	"time"
)

type VehicleOwner struct {
	bun.BaseModel  `bun:"vehicle_owners"`
	ID             int        `json:"id" bun:"id,pk,autoincrement"`
	VehicleOwnerID string     `json:"user_id" bun:"vehicle_owner_id,unique"`
	FirstName      string     `json:"first_name" bun:"first_name"`
	LastName       string     `json:"last_name" bun:"last_name"`
	Email          string     `json:"email" bun:"email,unique"`
	Address        *Address   `json:"address" bun:"address,rel:has-one,join:id=id"`
	Vehicles       []*Vehicle `json:"vehicles" bun:"vehicles,rel:has-many,join:id=vehicle_id"`
	Password       string     `json:"password" bun:"password"`
	CreatedAt      time.Time  `json:"created_at" bun:",nullzero,notnull,default:current_timestamp"`
	UpdatedAt      time.Time  `json:"updated_at" bun:",nullzero,notnull,default:current_timestamp"`
	DeletedAt      time.Time  `json:"deleted_at" bun:",soft_delete,nullzero"`
}
