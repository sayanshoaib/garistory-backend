package entity

import (
	"github.com/uptrace/bun"
	"time"
)

type User struct {
	bun.BaseModel `bun:"admins"`
	ID            int       `json:"id" bun:"id,pk,autoincrement"`
	UserID        string    `json:"user_id" bun:"admin_id,unique"`
	FirstName     string    `json:"first_name" bun:"first_name"`
	LastName      string    `json:"last_name" bun:"last_name"`
	Email         string    `json:"email" bun:"email,unique"`
	Password      string    `json:"password" bun:"password"`
	CreatedAt     time.Time `json:"created_at" bun:",nullzero,notnull,default:current_timestamp"`
	UpdatedAt     time.Time `json:"updated_at" bun:",nullzero,notnull,default:current_timestamp"`
	DeletedAt     time.Time `json:"deleted_at" bun:",soft_delete,nullzero"`
}
