package customers

import (
	"database/sql"
	"fmt"

	"simple_api/models/request"
	"simple_api/models/response"
	"simple_api/pkg/database"
)

func InsertCustomers(req request.InsertCustomers) error {
	conn, err := database.DBConnect()

	if err != nil {
		return nil
	}

	query, err := conn.Prepare(`INSERT INTO customers.customers SET customerId=?, phoneNUmber=?, status=?;`)
	if err != nil {
		return err
	}

	defer func(query *sql.Stmt) {
		_ = query.Close()
	}(query)

	_, errExec := query.Exec(req.CustomerID, req.PhoneNumber, req.Status)
	if errExec != nil {
		return errExec
	}

	return nil
}

func GetCustomers(req request.GetCustomers) (
	*response.GetCustomersResponse, error) {

	conn, err := database.DBConnect()

	if err != nil {
		return nil, err
	}

	res := &response.GetCustomersResponse{}

	err = conn.QueryRow(`SELECT c.customerId, c.phoneNumber, cs.description 
	FROM customers.customers c INNER JOIN
    customers.customers_status cs ON c.status=cs.status
	WHERE c.customerId =?;`, req.CustomerID).Scan(&res.CustomerID, &res.PhoneNumber, &res.Status)

	if err != nil {
		fmt.Printf("Error getting customers: %v", err)
		return nil, err
	}

	return res, nil
}
