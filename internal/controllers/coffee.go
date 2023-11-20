package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/lutefd/coffee-server/internal/helpers"
	"github.com/lutefd/coffee-server/internal/services"
)
type Message struct {
    Message string `json:"message"`
}
var coffee services.Coffee
// GetAllCoffees godoc
// @Summary Get all coffees
// @Description Retrieves a list of all coffee entries
// @Tags coffee
// @Accept json
// @Produce json
// @Success 200 {object} services.CoffeeList
// @Router /coffees [get]
func GetAllCofees(w http.ResponseWriter, r *http.Request) {
	all, err := coffee.GetAllCofees()
	if err != nil {
		helpers.MessageLogs.ErrorLog.Println("Error getting all coffees", err)
		return 
	}
	helpers.WriteJSON(w, http.StatusOK, helpers.Envelop{"coffees": all}, nil)
}

// GetCoffeeById godoc
// @Summary Get coffee by ID
// @Description Retrieves a coffee entry by its ID
// @Tags coffee
// @Accept json
// @Produce json
// @Param id path string true "Coffee ID"
// @Success 200 {object} services.CoffeeList
// @Router /coffees/{id} [get]

func GetCoffeeById(w http.ResponseWriter, r *http.Request){
	id := chi.URLParam(r, "id")
	c, err := coffee.GetCoffeeById(id)
	if err != nil {
		helpers.MessageLogs.ErrorLog.Println("Error getting coffee by id", err)
		return
	}
	helpers.WriteJSON(w, http.StatusOK, helpers.Envelop{"coffee": c}, nil)
}

// CreateCoffee godoc
// @Summary Create a coffee entry
// @Description Adds a new coffee entry to the list
// @Tags coffee
// @Accept json
// @Produce json
// @Param coffeeData body services.Coffee true "Coffee Data"
// @Success 201 {object} services.Coffee
// @Router /coffees [post]
func CreateCoffee(w http.ResponseWriter, r *http.Request) {
	var coffeeData services.Coffee
	err := json.NewDecoder(r.Body).Decode(&coffeeData)
	if err != nil {
		helpers.MessageLogs.ErrorLog.Println("Error decoding coffee data", err)
		return 
	}
	coffeeCreated, err := coffee.CreateCoffee(coffeeData)
	if err != nil {
		helpers.MessageLogs.ErrorLog.Println("Error creating coffee", err)
		return 
	}
	helpers.WriteJSON(w, http.StatusCreated, coffeeCreated, nil)
}

// UpdateCoffee godoc
// @Summary Update a coffee entry
// @Description Updates an existing coffee entry identified by ID
// @Tags coffee
// @Accept json
// @Produce json
// @Param id path string true "Coffee ID"
// @Param coffeeData body services.Coffee true "Coffee Data"
// @Success 200 {object} services.Coffee
// @Router /coffees/{id} [put]
func UpdateCoffee(w http.ResponseWriter, r *http.Request) {
	var coffeeData services.Coffee
	id := chi.URLParam(r, "id")
	err := json.NewDecoder(r.Body).Decode(&coffeeData)
	if err != nil {
		helpers.MessageLogs.ErrorLog.Println("Error decoding coffee data", err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	coffeeUpdated, err := coffee.UpdateCoffee(id, coffeeData)
	if err != nil {
		helpers.MessageLogs.ErrorLog.Println("Error updating coffee", err)
		return
	}
	helpers.WriteJSON(w, http.StatusOK, coffeeUpdated, nil)
}

// DeleteCoffee godoc
// @Summary Delete a coffee entry
// @Description Deletes a coffee entry identified by ID
// @Tags coffee
// @Accept json
// @Produce json
// @Param id path string true "Coffee ID"
// @Success 200 {object} Message
// @Router /coffees/{id} [delete]
func DeleteCoffee(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	err := coffee.DeleteCoffee(id)
	if err != nil {
		helpers.MessageLogs.ErrorLog.Println("Error deleting coffee", err)
		return
	}
	helpers.WriteJSON(w, http.StatusOK, helpers.Envelop{"message": "Coffee deleted"}, nil)
}