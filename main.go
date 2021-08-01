package main

import (
	r "goql/src"
	db "goql/src/config"

	"github.com/joho/godotenv"
)

func main() {
	// dotenv file load
	godotenv.Load("./deploy/env/goql.dotenv")

	// Database Connetion
	db.DBConnection()

	// Gin Server Start
	var s r.Routes
	s.StartGin()
}
