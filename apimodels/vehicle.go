package apimodels

import (
	"time"
)

type RespVehicle struct {
	VehicleID           string `json:"vehicleID"`
	Make                string `json:"make"`
	Model               string `json:"model"`
	Year                string `json:"year"`
	BodyType            string `json:"bodyType"`
	InitialRegistration string `json:"initialRegistration"`
	RegistrationNumber  string `json:"registrationNumber"`
	Fuel                int    `json:"fuel"`
	Mileage             int    `json:"mileage"`
	Drive               int    `json:"drive"`
	Color               string `json:"color"`
	Transmission        int    `json:"transmission"`
	Price               int    `json:"price"`
	ImageURL            string `json:"imageURL"`
}

type RespVehicleCreate struct {
	VehicleID           string    `json:"vehicleID"`
	Make                string    `json:"make"`
	Model               string    `json:"model"`
	Year                string    `json:"year"`
	BodyType            string    `json:"bodyType"`
	InitialRegistration string    `json:"initialRegistration"`
	RegistrationNumber  string    `json:"registrationNumber"`
	Fuel                int       `json:"fuel"`
	Mileage             int       `json:"mileage"`
	Drive               int       `json:"drive"`
	Color               string    `json:"color"`
	Transmission        int       `json:"transmission"`
	Price               int       `json:"price"`
	ImageURL            string    `json:"imageURL"`
	CreatedAt           time.Time `json:"time" `
	UpdatedAt           time.Time `json:"updatedAt"`
	DeletedAt           time.Time `json:"deletedAt"`
}

type RespVehicleUpdate struct {
	VehicleID           string    `json:"vehicleID"`
	Make                string    `json:"make"`
	Model               string    `json:"model"`
	Year                string    `json:"year"`
	BodyType            string    `json:"bodyType"`
	InitialRegistration string    `json:"initialRegistration"`
	RegistrationNumber  string    `json:"registrationNumber"`
	Fuel                int       `json:"fuel"`
	Mileage             int       `json:"mileage"`
	Drive               int       `json:"drive"`
	Color               string    `json:"color"`
	Transmission        int       `json:"transmission"`
	Price               int       `json:"price"`
	ImageURL            string    `json:"imageURL"`
	UpdatedAt           time.Time `json:"updatedAt"`
}
