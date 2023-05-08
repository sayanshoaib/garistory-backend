package entity

import (
	"github.com/uptrace/bun"
	"time"
)

type Address struct {
	bun.BaseModel `bun:"address"`
	ID            int       `json:"id" bun:"id,pk,autoincrement"`
	Name          string    `json:"name" bun:"name"`
	Street        string    `json:"street" bun:"street"`
	Zipcode       string    `json:"zipcode" bun:"zipcode"`
	Country       string    `json:"country" bun:"country"`
	CreatedAt     time.Time `json:"created_at" bun:",nullzero,default:current_timestamp"`
	UpdatedAt     time.Time `json:"updated_at" bun:",nullzero,default:current_timestamp"`
	DeletedAt     time.Time `json:"deleted_at" bun:",soft_delete,nullzero"`
}
