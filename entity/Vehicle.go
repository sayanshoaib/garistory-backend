package entity

import (
	"github.com/uptrace/bun"
	"time"
)

type Vehicle struct {
	bun.BaseModel       `bun:"vehicles"`
	ID                  int       `json:"id" bun:"id,pk,autoincrement"`
	VehicleID           string    `json:"vehicleID" bun:"vehicle_id,unique"`
	Make                string    `json:"make" bun:"make"`
	Model               string    `json:"model" bun:"model"`
	Year                string    `json:"year" bun:"year"`
	BodyType            string    `json:"bodyType" bun:"body_type"`
	InitialRegistration string    `json:"initialRegistration" bun:"initial_registration"`
	RegistrationNumber  string    `json:"registrationNumber" bun:"registration_number"`
	Fuel                int       `json:"fuel" bun:"fuel"`
	Mileage             string    `json:"mileage" bun:"mileage"`
	Drive               int       `json:"drive" bun:"drive"`
	Color               string    `json:"color" bun:"color"`
	Transmission        int       `json:"transmission" bun:"transmission"`
	Price               int       `json:"price" bun:"price"`
	ImageURL            string    `json:"imageURL" bun:"image_url"`
	CreatedAt           time.Time `json:"created_at" bun:",nullzero,notnull,default:current_timestamp"`
	UpdatedAt           time.Time `json:"updated_at" bun:",nullzero,notnull,default:current_timestamp"`
	DeletedAt           time.Time `json:"deleted_at" bun:",soft_delete,nullzero"`
}
