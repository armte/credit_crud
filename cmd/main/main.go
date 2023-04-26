package main

import (
	"fmt"
	"github.com/armte/credit_crud/pkg/routes"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func main() {
	r := mux.NewRouter()
	routes.RegisterCreditCrudRoutes(r)
	http.Handle("/", r)
	fmt.Printf("Starting server at port 9010")
	log.Fatal(http.ListenAndServe(":9010", r))
}
