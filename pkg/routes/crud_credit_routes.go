package routes

import (
	"github.com/gorilla/mux"
	"github.com/armte/credit_crud/pkg/controllers"
)

var RegisterCreditCrudRoutes = func(router *mux.Router) {
	router.HandleFunc("/hello/", controllers.HelloHandler).Methods("GET")
	router.HandleFunc("/customer/", controllers.CreateCustomer).Methods("POST")
	router.HandleFunc("/customer/", controllers.GetCustomer).Methods("GET")
	router.HandleFunc("/customer/{customerId}", controllers.GetCustomerById).Methods("GET")
	router.HandleFunc("/customer/{customerId}", controllers.UpdateCustomer).Methods("PUT")
	router.HandleFunc("/customer/{customerId}", controllers.DeleteCustomer).Methods("DELETE")
	router.HandleFunc("/account/", controllers.CreateAccount).Methods("POST")
	router.HandleFunc("/account/", controllers.GetAccount).Methods("GET")
	router.HandleFunc("/account/{accountId}", controllers.GetAccountById).Methods("GET")
	router.HandleFunc("/account/{accountId}", controllers.UpdateAccount).Methods("PUT")
	router.HandleFunc("/account/{accountId}", controllers.DeleteAccount).Methods("DELETE")
	router.HandleFunc("/card/", controllers.CreateCard).Methods("POST")
	router.HandleFunc("/card/", controllers.GetCard).Methods("GET")
	router.HandleFunc("/card/{cardId}", controllers.GetCardById).Methods("GET")
	router.HandleFunc("/card/{cardId}", controllers.UpdateCard).Methods("PUT")
	router.HandleFunc("/card/{cardId}", controllers.DeleteCard).Methods("DELETE")
}