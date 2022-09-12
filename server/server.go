package server

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"simple_api/handlers"
)

type Server struct {
	// services services.Repository
	// handlers handlers.Repository
	// sql customers.Repository
}

func Start() {
	r := mux.NewRouter()

	r.HandleFunc("/createCustomer", handlers.CreateCustomerHandler).Methods(http.MethodPost)
	r.HandleFunc("/getCustomer", handlers.GetCustomerHandler).Methods(http.MethodGet)

	err := http.ListenAndServe(":9000", r)
	if err != nil {
		fmt.Println("Error ListenAndServe")
		return
	}
}
