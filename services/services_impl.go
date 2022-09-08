package services

import (
	"fmt"

	"simple_api/models/request"
	"simple_api/models/response"
)

func (s customerService) InsertCustomers(req request.InsertCustomers) (*response.InsertCustomersResponse, error) {
	err := s.SQL.InsertCustomers(req)
	if err != nil {
		fmt.Printf("Error inserting customers : %v", err)
		return nil, err
	}
	resp := &response.InsertCustomersResponse{
		CustomerID:  req.CustomerID,
		PhoneNumber: req.PhoneNumber,
		Status:      req.Status,
	}

	return resp, nil
}

func (s customerService) GetCustomerDetail(req request.GetCustomers) (*response.GetCustomersResponse, error) {
	resp, err := s.SQL.GetCustomers(req)
	if err != nil {
		fmt.Printf("Error Get customers : %v", err)
		return nil, err
	}

	return resp, nil
}
