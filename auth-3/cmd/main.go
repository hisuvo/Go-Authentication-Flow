package main

import (
	"golang-authentication/internal/config"
	"golang-authentication/internal/server"
)

func main() {
	env := config.LoadEnv()

	db := config.ConnectDB(env)

	server.Start(db, env)
}