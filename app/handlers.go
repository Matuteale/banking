package app

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

type Customer struct {
	Name    string `json:"name"`
	City    string `json:"city"`
	ZipCode string `json:"zipcode"`
}

func greet(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello World!")
}

func getAllCustomers(w http.ResponseWriter, r *http.Request) {
	customers := []Customer{
		{
			Name:    "Matias",
			City:    "Bs As",
			ZipCode: "1425",
		},
		{
			Name:    "Agus",
			City:    "Bs As",
			ZipCode: "1425",
		},
	}
	w.Header().Add("Contenct-Type", "application/json")
	json.NewEncoder(w).Encode(customers)
}

func getCustomer(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	fmt.Fprint(w, vars["customer_id"])
}
