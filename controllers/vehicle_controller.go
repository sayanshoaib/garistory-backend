package controllers

import (
	"garistroy-backend/apimodels"
	"garistroy-backend/service"
	"github.com/labstack/echo/v4"
	"log"
	"net/http"
)

type VehicleController struct {
	VehicleService *service.VehicleService
}

func NewVehicleService(vehicleService *service.VehicleService) *VehicleController {
	return &VehicleController{
		VehicleService: vehicleService,
	}
}

func (c *VehicleController) CreateVehicle(ctx echo.Context) error {
	vehicleRequest := &apimodels.ReqVehicleCreate{}
	if err := ctx.Bind(vehicleRequest); err != nil {
		return err
	}

	resp, err := c.VehicleService.CreateVehicle(ctx.Request().Context(), vehicleRequest)
	if err != nil {
		return err
	}

	return ctx.JSON(http.StatusCreated, map[string]interface{}{
		"message": "Successfully created a new vehicle",
		"data":    resp,
	})
}

func (c *VehicleController) GetVehicleByVin(ctx echo.Context) error {
	vin := ctx.Param("vin")
	resp, err := c.VehicleService.GetVehicleByVin(ctx.Request().Context(), vin)
	if err != nil {
		return ctx.JSON(http.StatusNotFound, map[string]interface{}{
			"message": "No vehicle found with vin: " + vin,
			"error":   err,
		})
	}

	return ctx.JSON(http.StatusOK, map[string]interface{}{
		"message": "Found a vehicle with vin: " + vin,
		"data":    resp,
	})
}

func (c *VehicleController) GetAllVehicle(ctx echo.Context) error {
	resp, err := c.VehicleService.GetAllVehicle(ctx.Request().Context())
	if err != nil {
		log.Print(err)
		return ctx.JSON(http.StatusNotFound, map[string]interface{}{
			"message": "No vehicle found",
		})
	}

	return ctx.JSON(http.StatusOK, map[string]interface{}{
		"data": resp,
	})
}

func (c *VehicleController) UpdateVehicleByVin(ctx echo.Context) error {
	vehicleToUpdate := &apimodels.ReqVehicleUpdate{}
	if err := ctx.Bind(vehicleToUpdate); err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "Invalid vehicle update request",
			"error":   err,
		})
	}

	vin := vehicleToUpdate.VehicleID

	resp, err := c.VehicleService.UpdateVehicleByVin(ctx.Request().Context(), vin, vehicleToUpdate)
	if err != nil {
		return ctx.JSON(http.StatusNotFound, map[string]interface{}{
			"message": "Vehicle not found with vin: " + vin,
			"error":   err,
		})
	}

	return ctx.JSON(http.StatusOK, map[string]interface{}{
		"message": "Vehicle successfully updated",
		"data":    resp,
	})
}

func (c *VehicleController) DeleteVehicleByVin(ctx echo.Context) error {
	vin := ctx.Param("vin")
	_, err := c.VehicleService.DeleteVehicleByVin(ctx.Request().Context(), vin)
	if err != nil {
		return ctx.JSON(http.StatusNotFound, map[string]interface{}{
			"message": "Vehicle not found with vin: " + vin,
			"error":   err,
		})
	}

	return ctx.NoContent(http.StatusNoContent)
}
