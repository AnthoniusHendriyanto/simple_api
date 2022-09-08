package customers

import (
	"database/sql"
	"fmt"

	"simple_api/models/request"
	"simple_api/models/response"
)

func (o *customerClient) InsertCustomers(req request.InsertCustomers) error {
	query, err := o.DB.Prepare(`INSERT INTO customers.customers SET customerId=?, phoneNUmber=?, status=?;`)
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

func (o *customerClient) GetCustomers(req request.GetCustomers) (
	*response.GetCustomersResponse, error) {

	res := &response.GetCustomersResponse{}

	err := o.DB.QueryRow(`SELECT c.customerId, c.phoneNumber, cs.description 
	FROM customers.customers c INNER JOIN
    customers.customers_status cs ON c.status=cs.status
	WHERE c.customerId =?;`, req.CustomerID).Scan(&res.CustomerID, &res.PhoneNumber, &res.Status)

	if err != nil {
		fmt.Printf("Error getting customers: %v", err)
		return nil, err
	}

	return res, nil
}
