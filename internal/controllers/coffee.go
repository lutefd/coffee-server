package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/lutefd/coffee-server/internal/helpers"
	"github.com/lutefd/coffee-server/internal/services"
)

var coffee services.Coffee

func GetAllCofees(w http.ResponseWriter, r *http.Request) {
	all, err := coffee.GetAllCofees()
	if err != nil {
		helpers.MessageLogs.ErrorLog.Println("Error getting all coffees", err)
		return 
	}
	helpers.WriteJSON(w, http.StatusOK, helpers.Envelop{"coffees": all}, nil)
}

func GetCoffeeById(w http.ResponseWriter, r *http.Request){
	id := chi.URLParam(r, "id")
	c, err := coffee.GetCoffeeById(id)
	if err != nil {
		helpers.MessageLogs.ErrorLog.Println("Error getting coffee by id", err)
		return
	}
	helpers.WriteJSON(w, http.StatusOK, helpers.Envelop{"coffee": c}, nil)
}

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

func DeleteCoffee(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	err := coffee.DeleteCoffee(id)
	if err != nil {
		helpers.MessageLogs.ErrorLog.Println("Error deleting coffee", err)
		return
	}
	helpers.WriteJSON(w, http.StatusOK, helpers.Envelop{"message": "Coffee deleted"}, nil)
}