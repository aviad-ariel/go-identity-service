package main

import (
	"go-identity-service/config"
	"go-identity-service/db"
	"go-identity-service/routes"
)

func main() {
	config.LoadConfig(".")
	db.MongoClient = db.Connect()
	router := routes.SetupServer()
	err := router.Run(config.Env.Port)
	if err != nil {
		panic("Failed to start server")
	}
}
