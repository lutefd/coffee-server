package main

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/lutefd/coffee-server/internal/database"
	"github.com/lutefd/coffee-server/internal/server"
)

func main() {
	err := godotenv.Load()

	if err != nil {
		log.Fatal("Error loading .env file")
	}
	cfg := server.Config{
		Port: os.Getenv("PORT"),
	}

	dsn := os.Getenv("DSN")

	dbConn, err := database.ConnectPostgresDB(dsn)
	if err != nil {
		log.Fatal("Cannot connect do database ", err)
	}
	defer dbConn.DB.Close()
	
	app := server.Application{
		Config: cfg,
		//TODO: add models
	}
	err = app.Serve()
	if err != nil {
		log.Fatal(err)
	}

}