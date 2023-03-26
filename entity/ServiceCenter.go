package entity

import (
	"github.com/uptrace/bun"
	"time"
)

type ServiceCenter struct {
	bun.BaseModel   `bun:"service_center"`
	ID              int       `json:"id" bun:"id,pk,autoincrement"`
	ServiceCenterID string    `json:"serviceCenterID" bun:"service_center,unique"`
	Name            string    `json:"name" bun:"name"`
	Email           string    `json:"email" bun:"email,unique"`
	Address         *Address  `json:"address" bun:"address,rel:has-one,join:id=id"`
	Password        string    `json:"password" bun:"password"`
	CreatedAt       time.Time `json:"created_at" bun:",nullzero,notnull,default:current_timestamp"`
	UpdatedAt       time.Time `json:"updated_at" bun:",nullzero,notnull,default:current_timestamp"`
	DeletedAt       time.Time `json:"deleted_at" bun:",soft_delete,nullzero"`
}
