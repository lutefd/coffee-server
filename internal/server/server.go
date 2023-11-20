package server

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	"github.com/lutefd/coffee-server/internal/router"
	"github.com/lutefd/coffee-server/internal/services"
)

type Config struct {
	Port string
}

type Application struct {
	Config Config
	Models services.Models
}

func (app *Application) Serve() error {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	port := os.Getenv("PORT")
	fmt.Println("API is running on port ",  port)
	srv := &http.Server{
		Addr: fmt.Sprintf(":%s", port),
		Handler: router.Routes(),
		
	}
	return srv.ListenAndServe()
}
