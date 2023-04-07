package api

import (
	"github.com/labstack/echo/v4"
)

type Server struct {
	Echo *echo.Echo
}

func NewServer() *Server {
	return &Server{
		Echo: echo.New(),
	}
}

func (s *Server) Start() {
	//s.Echo.Use(middleware.Logger())
	//s.Echo.Use(middleware.Recover())
	//
	//repo := repository.NewVehicleRepository(config.ConnectDb())
	//service := service2.NewVehicleRepository(repo)
	//controller := controllers.NewVehicleService(service)
	//
	//s.Echo.POST("/vehicle", controller.CreateVehicle)
	//s.Echo.GET("/vehicle", controller.GetAllVehicle)
	//s.Echo.GET("/vehicle/:vin", controller.GetVehicleByVin)
	//s.Echo.PUT("/vehicle/:vin", controller.UpdateVehicleByVin)
	//s.Echo.DELETE("/vehicle/:vin", controller.DeleteVehicleByVin)
	//
	//s.Echo.Logger.Fatal(s.Echo.Start(":8080"))
}
