package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type APIServer struct {
	addr  string
	store Store
}

// CNewAPIServer constructor for the service
func NewAPIServer(addr string, store Store) *APIServer {
	return &APIServer{addr: addr, store: store}
}

// Adding the method that the service is going to have, which is to initialize the router and register all of the services and their dependencies
func (s *APIServer) Serve() {
	router := mux.NewRouter()
	subrouter := router.PathPrefix("/api/v1").Subrouter()

	// registering our services
	tasksService := NewTasksService(s.store)
	tasksService.RegisterRoute(router)

	log.Println("Starting the API server at ", s.addr)
	log.Fatal(http.ListenAndServe(s.addr, subrouter))

}
