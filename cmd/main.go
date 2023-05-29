package main

import (
	"garistroy-backend/config"
	"garistroy-backend/controllers"
	"garistroy-backend/repository"
	service2 "garistroy-backend/service"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	//server := api.NewServer()
	//server.Start()
	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept},
	}))

	repo := repository.NewVehicleRepository(config.ConnectDb())
	service := service2.NewVehicleRepository(repo)
	controller := controllers.NewVehicleService(service)

	addressRepo := repository.NewAddressRepository(config.ConnectDb())
	addressService := service2.NewAddressService(addressRepo)
	addressController := controllers.NewAddressController(addressService)

	servicingCenterRepo := repository.NewServicingCenterRepository(config.ConnectDb(), addressRepo)
	servicingCenterService := service2.NewServicingCenterService(servicingCenterRepo)
	servicingCenterController := controllers.NewServiceCenterController(servicingCenterService)

	insuranceCompanyRepo := repository.NewInsuranceCompanyRepository(config.ConnectDb(), addressRepo)
	insuranceCompanyService := service2.NewInsuranceCompanyService(insuranceCompanyRepo)
	insuranceCompanyController := controllers.NewInsuranceCompanyService(insuranceCompanyService)

	// Vehicles endpoints
	e.POST("/vehicles", controller.CreateVehicle)
	e.GET("/vehicles", controller.GetAllVehicle)
	e.GET("/vehicles/:vin", controller.GetVehicleByVin)
	e.PUT("/vehicles/:vin", controller.UpdateVehicleByVin)
	e.DELETE("/vehicles/:vin", controller.DeleteVehicleByVin)

	// Servicing center endpoints
	e.POST("/service-center", servicingCenterController.RegisterServicingCenter)
	e.GET("/service-center", servicingCenterController.GetAllServiceCenter)
	e.GET("/service-center/:servicing-center-id", servicingCenterController.GetServiceCenterByID)
	e.PUT("/service-center/:servicing-center-id", servicingCenterController.UpdateServiceCenterByID)
	e.DELETE("/service-center/:servicing-center-id", servicingCenterController.DeleteServiceCenterByID)

	// Address endpoints
	e.POST("/service-center/address", addressController.CreateAddress)
	e.GET("/address", addressController.GetAllAddress)
	e.GET("/service-center/address/:id", addressController.GetAddressByID)
	e.GET("/service-center/address/:zipcode", addressController.GetAddressByZipcode)
	e.PUT("/service-center/address/:id", addressController.UpdateAddressByID)
	e.DELETE("/service-center/address/:id", addressController.DeleteAddressByID)

	// Insurance Company endpoints
	e.POST("/insurance-companies", insuranceCompanyController.CreateInsuranceCompany)
	//e.GET("/insurance-companies")
	//e.GET("/insurance-companies/:insurance-company-id")
	//e.PUT("/insurance-companies/:insurance-company-id")
	//e.DELETE("/insurance-companies/:insurance-company-id")

	// User
	//e.POST("/users")
	//e.GET("/users")
	//e.GET("/users/:id")
	e.Logger.Fatal(e.Start(":8080"))
}
