package controllers

import (
	"garistroy-backend/apimodels"
	"garistroy-backend/service"
	"github.com/labstack/echo/v4"
	"log"
	"net/http"
)

type ServiceCenterController struct {
	ServicingCenterService *service.ServicingCenterService
}

func NewServiceCenterController(service *service.ServicingCenterService) *ServiceCenterController {
	return &ServiceCenterController{
		ServicingCenterService: service,
	}
}

func (c *ServiceCenterController) RegisterServicingCenter(ctx echo.Context) error {
	reqServiceCenter := &apimodels.ReqServicingCenter{}
	if err := ctx.Bind(reqServiceCenter); err != nil {
		log.Print(err)
		return ctx.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": false,
			"error":   err,
		})
	}

	resp, err := c.ServicingCenterService.CreateServicingCenter(ctx.Request().Context(), reqServiceCenter)
	if err != nil {
		log.Print(err)
		return ctx.JSON(http.StatusOK, map[string]interface{}{
			"message": false,
			"error":   err,
		})
	}

	return ctx.JSON(http.StatusOK, map[string]interface{}{
		"message": true,
		"data":    resp,
	})
}

func (c *ServiceCenterController) GetServiceCenterByID(ctx echo.Context) error {
	servicingCenterID := ctx.Param("servicing-center-id")
	resp, err := c.ServicingCenterService.GetServicingCenterByID(ctx.Request().Context(), servicingCenterID)
	if err != nil {
		log.Print(err)
		return ctx.JSON(http.StatusOK, map[string]interface{}{
			"message": false,
			"error":   err,
		})
	}

	return ctx.JSON(http.StatusOK, map[string]interface{}{
		"message": true,
		"data":    resp,
	})
}

func (c *ServiceCenterController) GetAllServiceCenter(ctx echo.Context) error {
	resp, count, err := c.ServicingCenterService.GetAllServicingCenter(ctx.Request().Context())
	if err != nil {
		log.Print(err)
		return ctx.JSON(http.StatusOK, map[string]interface{}{
			"message": false,
			"error":   err,
		})
	}

	return ctx.JSON(http.StatusOK, map[string]interface{}{
		"message": true,
		"count":   count,
		"data":    resp,
	})
}

func (c *ServiceCenterController) UpdateServiceCenterByID(ctx echo.Context) error {
	reqServiceCenter := &apimodels.ReqServicingCenter{}
	if err := ctx.Bind(reqServiceCenter); err != nil {
		log.Print(err)
		return ctx.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": false,
			"error":   err,
		})
	}

	resp, err := c.ServicingCenterService.UpdateServicingCenterByID(ctx.Request().Context(), reqServiceCenter)
	if err != nil {
		log.Print(err)
		return ctx.JSON(http.StatusOK, map[string]interface{}{
			"message": false,
			"error":   err,
		})
	}

	return ctx.JSON(http.StatusOK, map[string]interface{}{
		"message": true,
		"data":    resp,
	})
}

func (c *ServiceCenterController) DeleteServiceCenterByID(ctx echo.Context) error {
	servicingCenterID := ctx.Param("servicing-center-id")
	ok, err := c.ServicingCenterService.DeleteServiceCenterByID(ctx.Request().Context(), servicingCenterID)
	if err != nil {
		log.Print(err)
		return ctx.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": ok,
			"error":   err,
		})
	}

	if !ok {
		return ctx.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": ok,
			"error":   apimodels.ErrNoDataMatched,
		})
	}

	return ctx.NoContent(http.StatusNoContent)
}
