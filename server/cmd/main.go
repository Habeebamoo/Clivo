package main

import (
	"github.com/Habeebamoo/Clivo/server/cmd/server"
	"github.com/Habeebamoo/Clivo/server/internal/config"
)

func main() {
	//load config files
	config.Initialize()

	//connect to database
	server.ConnectDB()

	//setup routes
	router := server.ConfigureApp()

	//run the server
	server.Run(router)
}