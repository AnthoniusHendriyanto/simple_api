package server

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"simple_api/handlers"
	"simple_api/services"
	"simple_api/store/mysql/customers"
)

type Server struct {
	services services.Repository
	handlers handlers.Repository
	sql      customers.Repository
}

func Register() *Server {
	SVR := &Server{}

	SVR.sql = customers.NewClient()

	SVR.services = services.NewService(SVR.sql)

	SVR.handlers = handlers.New(SVR.services)

	return SVR
}

func (s *Server) Start() {
	r := mux.NewRouter()

	r.HandleFunc("/createCustomer", s.handlers.CreateCustomerHandler).Methods(http.MethodPost)
	r.HandleFunc("/getCustomer", s.handlers.GetCustomerHandler).Methods(http.MethodGet)

	err := http.ListenAndServe(":9000", r)
	if err != nil {
		fmt.Println("Error ListenAndServe")
		return
	}
}
