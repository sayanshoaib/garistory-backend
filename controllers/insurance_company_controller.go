package controllers

import (
	"garistroy-backend/apimodels"
	"garistroy-backend/service"
	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
	"net/http"
)

type InsuranceCompanyController struct {
	InsuranceCompanyService *service.InsuranceCompanyService
}

func NewInsuranceCompanyService(service *service.InsuranceCompanyService) *InsuranceCompanyController {
	return &InsuranceCompanyController{
		InsuranceCompanyService: service,
	}
}

func (c *InsuranceCompanyController) CreateInsuranceCompany(ctx echo.Context) error {
	insuranceCompanyDTO := &apimodels.InsuranceCompanyDTO{}
	if err := ctx.Bind(insuranceCompanyDTO); err != nil {
		log.Info(err)
		return ctx.JSON(http.StatusBadRequest, map[string]interface{}{
			"success": false,
			"error":   err,
		})
	}

	resp, err := c.InsuranceCompanyService.CreateInsuranceCompany(ctx.Request().Context(), insuranceCompanyDTO)
	if err != nil {
		log.Info(err)
		return ctx.JSON(http.StatusInternalServerError, map[string]interface{}{
			"success": false,
			"error":   err,
		})
	}

	return ctx.JSON(http.StatusCreated, map[string]interface{}{
		"success": true,
		"data":    resp,
	})
}
