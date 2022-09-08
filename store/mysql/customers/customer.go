package customers

import (
	"database/sql"

	"simple_api/models/request"
	"simple_api/models/response"
	"simple_api/pkg/database"
)

type customerClient struct {
	DB *sql.DB
}

type Repository interface {
	InsertCustomers(req request.InsertCustomers) error
	GetCustomers(req request.GetCustomers) (
		*response.GetCustomersResponse, error)
}

func NewClient() *customerClient {
	conn, err := database.DBConnect()

	if err != nil {
		return nil
	}

	return &customerClient{
		DB: conn,
	}
}
