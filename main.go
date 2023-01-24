package main

import (
	"stupix/config"
	"stupix/db"
	"stupix/routes"
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
