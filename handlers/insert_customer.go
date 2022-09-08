package handlers

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"simple_api/models/request"
)

func (s customerHandler) CreateCustomerHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set(contentType, MIMEApplicationJSON)

	body, err := io.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Println("Errors Read All")
		return
	}
	createCustomer := request.InsertCustomers{}

	err = json.Unmarshal(body, &createCustomer)
	if err != nil {
		w.WriteHeader(http.StatusPreconditionFailed)
		fmt.Println("Errors Unmarshal Create Customer")
		return
	}

	fmt.Printf("Request: %v\n", createCustomer)

	resp, err := s.services.InsertCustomers(createCustomer)
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
