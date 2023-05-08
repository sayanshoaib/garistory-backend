package controllers

import (
	"fmt"
	"garistroy-backend/apimodels"
	"garistroy-backend/service"
	"github.com/labstack/echo/v4"
	"net/http"
	"strconv"
)

type AddressController struct {
	AddressService *service.AddressService
}

func NewAddressController(addressService *service.AddressService) *AddressController {
	return &AddressController{
		AddressService: addressService,
	}
}

func (c *AddressController) CreateAddress(ctx echo.Context) error {
	reqAddress := &apimodels.Address{}
	if err := ctx.Bind(reqAddress); err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": false,
			"error":   err,
		})
	}

	resp, err := c.AddressService.CreateAddress(ctx.Request().Context(), reqAddress)
	if err != nil {
		return ctx.JSON(http.StatusOK, map[string]interface{}{
			"message": false,
			"error":   err,
		})
	}

	return ctx.JSON(http.StatusCreated, map[string]interface{}{
		"message": true,
		"data":    resp,
	})
}

func (c *AddressController) GetAddressByID(ctx echo.Context) error {
	addressID := ctx.Param("id")
	resp, err := c.AddressService.GetAddressByID(ctx.Request().Context(), addressID)
	if err != nil {
		return ctx.JSON(http.StatusNotFound, map[string]interface{}{
			"message": false,
			"error":   err,
		})
	}

	return ctx.JSON(http.StatusOK, map[string]interface{}{
		"message": true,
		"data":    resp,
	})
}

func (c *AddressController) GetAddressByZipcode(ctx echo.Context) error {
	zipcode := ctx.Param("zipcode")
	resp, err := c.AddressService.GetAddressByZipcode(ctx.Request().Context(), zipcode)
	if err != nil {
		return ctx.JSON(http.StatusNotFound, map[string]interface{}{
			"message": false,
			"error":   err,
		})
	}

	return ctx.JSON(http.StatusOK, map[string]interface{}{
		"message": true,
		"data":    resp,
	})
}

func (c *AddressController) GetAllAddress(ctx echo.Context) error {
	resp, count, err := c.AddressService.GetAllAddress(ctx.Request().Context())
	if err != nil {
		return ctx.JSON(http.StatusNotFound, map[string]interface{}{
			"message": false,
			"count":   count,
			"error":   err,
		})
	}

	return ctx.JSON(http.StatusOK, map[string]interface{}{
		"message": true,
		"count":   count,
		"data":    resp,
	})
}

func (c *AddressController) UpdateAddressByID(ctx echo.Context) error {
	addressID := ctx.Param("id")
	id, err := strconv.Atoi(addressID)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": false,
			"error":   err,
		})
	}

	reqAddress := &apimodels.Address{}
	if err := ctx.Bind(reqAddress); err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": false,
			"error":   err,
		})
	}

	resp, err := c.AddressService.UpdateAddressByID(ctx.Request().Context(), id, reqAddress)
	if err != nil {
		return ctx.JSON(http.StatusNotFound, map[string]interface{}{
			"message": false,
			"error":   err,
		})
	}

	return ctx.JSON(http.StatusOK, map[string]interface{}{
		"message": true,
		"data":    resp,
	})
}

func (c *AddressController) DeleteAddressByID(ctx echo.Context) error {
	addressID := ctx.Param("id")
	id, err := strconv.Atoi(addressID)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": false,
			"error":   err,
		})
	}

	ok, err := c.AddressService.DeleteAddressByID(ctx.Request().Context(), id)
	if err != nil {
		return ctx.JSON(http.StatusNotFound, map[string]interface{}{
			"message": ok,
			"error":   err,
		})
	}

	return ctx.JSON(http.StatusNoContent, map[string]interface{}{
		"message": ok,
		"data":    fmt.Sprintf("Successfully deleted address with id: %d", id),
	})
}
