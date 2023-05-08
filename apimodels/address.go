package apimodels

import (
	"errors"
	"time"
)

type RespAddress struct {
	Name      string    `json:"name"`
	Street    string    `json:"street"`
	Zipcode   string    `json:"zipcode"`
	Country   string    `json:"country"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

type Address struct {
	Name    string `json:"name"`
	Street  string `json:"street"`
	Zipcode string `json:"zipcode"`
	Country string `json:"country"`
}

func (address *Address) ValidateAddress() error {
	if address.Street == "" {
		return errors.New("street cannot be empty")
	}

	if address.Zipcode == "" {
		return errors.New("zipcode cannot be empty")
	}

	if address.Country == "" {
		return errors.New("country cannot be empty")
	}

	return nil
}
