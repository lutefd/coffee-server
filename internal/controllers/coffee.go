package controllers

import (
	"encoding/json"
	"net/http"

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

func GetCoffeeById(w http.ResponseWriter, r *http.Request) (*services.Coffee, error){
	id := r.URL.Query().Get("id")
	c, err := coffee.GetCoffeeById(id)
	if err != nil {
		helpers.MessageLogs.ErrorLog.Println("Error getting coffee by id", err)
		return nil, err
	}
	helpers.WriteJSON(w, http.StatusOK, helpers.Envelop{"coffee": c}, nil)
	return c, nil
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
	helpers.WriteJSON(w, http.StatusOK, coffeeCreated, nil)
}