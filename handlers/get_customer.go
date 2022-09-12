package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"simple_api/models/request"
	"simple_api/services"
)

const (
	contentType         = "Content-Type"
	MIMEApplicationJSON = "application/json"
	customerID          = "customerId"
)

// func (s customerHandler) GetCustomerHandler(w http.ResponseWriter, r *http.Request) {
// 	w.Header().Set(contentType, MIMEApplicationJSON)
//
// 	customerId := r.URL.Query().Get(customerID)
//
// 	getCustomer := request.GetCustomers{
// 		CustomerID: customerId,
// 	}
//
// 	fmt.Printf("Request: %v\n", getCustomer)
//
// 	resp, err := s.services.GetCustomerDetail(getCustomer)
// 	if err != nil {
// 		w.WriteHeader(http.StatusServiceUnavailable)
// 		fmt.Println("Error inserting customers")
// 		return
// 	}
//
// 	fmt.Printf("Response: %v\n", resp)
//
// 	w.WriteHeader(http.StatusCreated)
// 	err = json.NewEncoder(w).Encode(resp)
//
// 	if err != nil {
// 		return
// 	}
// }

func GetCustomerHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set(contentType, MIMEApplicationJSON)

	customerId := r.URL.Query().Get(customerID)

	getCustomer := request.GetCustomers{
		CustomerID: customerId,
	}

	fmt.Printf("Request: %v\n", getCustomer)

	resp, err := services.GetCustomerDetail(getCustomer)
	if err != nil {
		w.WriteHeader(http.StatusServiceUnavailable)
		fmt.Println("Error inserting customers")
		return
	}

	fmt.Printf("Response: %v\n", resp)

	w.WriteHeader(http.StatusCreated)
	err = json.NewEncoder(w).Encode(resp)

	if err != nil {
		return
	}
}
