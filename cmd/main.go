package main

import "garistroy-backend/api"

func main() {
	server := api.NewServer()
	server.Start()
}
