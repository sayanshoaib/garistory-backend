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

	repo := repository.NewVehicleRepository(config.ConnectDb())
	service := service2.NewVehicleRepository(repo)
	controller := controllers.NewVehicleService(service)

	e.POST("/vehicles", controller.CreateVehicle)
	e.GET("/vehicles", controller.GetAllVehicle)
	e.GET("/vehicles/:vin", controller.GetVehicleByVin)
	e.PUT("/vehicles/:vin", controller.UpdateVehicleByVin)
	e.DELETE("/vehicles/:vin", controller.DeleteVehicleByVin)

	e.Logger.Fatal(e.Start(":8080"))
}
