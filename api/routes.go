package api

import (
	"garistroy-backend/config"
	"garistroy-backend/controllers"
	"garistroy-backend/repository"
	service2 "garistroy-backend/service"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type Server struct {
	echo *echo.Echo
}

func NewServer() *Server {
	return &Server{
		echo: echo.New(),
	}
}

func (s *Server) Start() {
	s.echo.Use(middleware.Logger())
	s.echo.Use(middleware.Recover())
	repo := repository.NewVehicleRepository(config.DB.Db)
	service := service2.NewVehicleRepository(repo)
	controller := controllers.NewVehicleService(service)

	s.echo.POST("/vehicle", controller.CreateVehicle)
	s.echo.GET("/vehicle", controller.GetAllVehicle)
	s.echo.GET("/vehicle/:vin", controller.GetVehicleByVin)
	s.echo.PUT("/vehicle/:vin", controller.UpdateVehicleByVin)
	s.echo.DELETE("/vehicle/:vin", controller.DeleteVehicleByVin)

	s.echo.Logger.Fatal(s.echo.Start(":8080"))
}
