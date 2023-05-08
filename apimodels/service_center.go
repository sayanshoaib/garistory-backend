package apimodels

import (
	"net/mail"
	"strings"
)

type RespServiceCenter struct {
	ServiceCenterID string   `json:"serviceCenterID"`
	Name            string   `json:"name"`
	Email           string   `json:"email"`
	Address         *Address `json:"address"`
}

type ReqServicingCenter struct {
	ServiceCenterID string   `json:"serviceCenterID"`
	Name            string   `json:"name"`
	Email           string   `json:"email"`
	Address         *Address `json:"address"`
	Password        string   `json:"password"`
}

func (req *ReqServicingCenter) Validate() error {
	req.Name = strings.TrimSpace(req.Name)
	req.Email = strings.TrimSpace(req.Email)

	vErr := &ValidationError{}

	if req.Name == "" {
		vErr.Add("name", "is required")
	}
	if req.Email == "" {
		vErr.Add("email", "is required")
	}
	_, err := mail.ParseAddress(req.Email)
	if err != nil {
		vErr.Add("email", "is invalid")
	}

	if vErr.HasErrors() {
		return vErr
	}
	return nil
}
