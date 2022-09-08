package services

import (
	"simple_api/models/request"
	"simple_api/models/response"
	"simple_api/store/mysql/customers"
)

type customerService struct {
	SQL customers.Repository
}

type Repository interface {
	InsertCustomers(req request.InsertCustomers) (*response.InsertCustomersResponse, error)
	GetCustomerDetail(req request.GetCustomers) (*response.GetCustomersResponse, error)
}

func NewService(sqlAgent customers.Repository) *customerService {
	return &customerService{
		SQL: sqlAgent,
	}
}
