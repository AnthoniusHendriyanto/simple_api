package response

type InsertCustomersResponse struct {
	CustomerID  string `json:"customer_id"`
	PhoneNumber string `json:"phone_number"`
	Status      int    `json:"status"`
}

type GetCustomersResponse struct {
	CustomerID  string `json:"customer_id"`
	PhoneNumber string `json:"phone_number"`
	Status      string `json:"status"`
}
