package main

import (
	"log"
	"os"

	"github.com/joho/godotenv"
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

	//TODO: connect to database
	app := server.Application{
		Config: cfg,
		//TODO: add models
	}
	err = app.Serve()
	if err != nil {
		log.Fatal(err)
	}

}