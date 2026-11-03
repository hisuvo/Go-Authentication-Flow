package main

import (
	"authentication-test/internal/config"
	"authentication-test/internal/database"
	"authentication-test/internal/server"
)

func main() {
	
	cnfg := config.LoadEnv()

	db, _ := database.ConnectDatabase(&cnfg)
	server.Start(db, &cnfg)
}