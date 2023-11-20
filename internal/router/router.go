package router

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/cors"
	_ "github.com/lutefd/coffee-server/docs" // Import your generated docs
	"github.com/lutefd/coffee-server/internal/controllers"
	httpSwagger "github.com/swaggo/http-swagger"
)

// @title Coffee API
// @version 1
// @description This is a sample server for a coffee store.
// @BasePath /api/v1
func Routes() http.Handler{
    router := chi.NewRouter()
    router.Use(middleware.Recoverer)
    router.Use(middleware.Logger)
    router.Use(cors.Handler(cors.Options{
        AllowedOrigins: []string{"*"},
        AllowedMethods: []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"},
        AllowedHeaders: []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
        ExposedHeaders: []string{"Link"},
        AllowCredentials: true,
        MaxAge: 300,
    }))
    router.Get("/", func(w http.ResponseWriter, r *http.Request) {
        w.Write([]byte("Welcome to Coffee API"))
    })
    router.Get("/api/v1/coffees", controllers.GetAllCofees)
    router.Get("/api/v1/coffees/{id}", controllers.GetCoffeeById)
    router.Post("/api/v1/coffees/create", controllers.CreateCoffee)
    router.Put("/api/v1/coffees/{id}", controllers.UpdateCoffee)
    router.Delete("/api/v1/coffees/{id}", controllers.DeleteCoffee)

    // Add the route that serves the Swagger UI
    router.Get("/swagger/*", httpSwagger.WrapHandler)

    return router
}