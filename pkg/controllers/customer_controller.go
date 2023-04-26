package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/biter777/countries"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
	"github.com/armte/credit_crud/pkg/utils"
	"github.com/armte/credit_crud/pkg/models"
)

var NewCustomer models.Customer

func GetCustomer(w http.ResponseWriter, r *http.Request) {
	newCustomers := models.GetAllCustomers()
	res, _ := json.Marshal(newCustomers)
	w.Header().Set("Content-Type","application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func GetCustomerById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	customerId := vars["customerId"]
	Id, err := strconv.ParseInt(customerId,0,0)
	if err != nil {
		fmt.Println("error while parsing")
		return
	}

	customerDetails, _, errGet := models.GetCustomerById(Id)
	if errGet != nil {
		http.Error(w, "404 Not Found: " + errGet.Error(), http.StatusNotFound)
		return
	}

	res, _ := json.Marshal(customerDetails)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func CreateCustomer(w http.ResponseWriter, r *http.Request) {
	CreateCustomer := &models.Customer{}
	utils.ParseBody(r, CreateCustomer)
	if countries.ByName(CreateCustomer.Customer_Country_Code) == 0 {
		http.Error(w, "400 Bad Request: Country Code Invalid or Missing", http.StatusBadRequest)
		return
	}
	c, err := CreateCustomer.CreateCustomer()
	if err != nil {
		http.Error(w, "400 Bad Request: " + err.Error(), http.StatusBadRequest)
		return
	}
	res, _ := json.Marshal(c)
	w.WriteHeader(http.StatusCreated)
	w.Write(res)
}

func DeleteCustomer(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	customerId := vars["customerId"]
	Id, err := strconv.ParseInt(customerId, 0, 0)
	if err != nil {
		fmt.Println("error while parsing")
	}
	customer, errGet := models.DeleteCustomer(Id)
	if errGet != nil {
		http.Error(w, "404 Not Found: " + errGet.Error(), http.StatusNotFound)
		return
	}
	res, _ := json.Marshal(customer)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func UpdateCustomer(w http.ResponseWriter, r *http.Request) {
	updateCustomer := &models.Customer{}
	utils.ParseBody(r, updateCustomer)
	vars := mux.Vars(r)
	customerId := vars["customerId"]
	Id, err := strconv.ParseInt(customerId, 0, 0)
	if err != nil {
		fmt.Println("error while parsing")
	}

	customerDetails, db, errGet := models.GetCustomerById(Id)

	if errGet != nil {
		http.Error(w, "404 Not Found: " + errGet.Error(), http.StatusNotFound)
		return
	}

	if updateCustomer.Customer_Name != "" {
		customerDetails.Customer_Name = updateCustomer.Customer_Name
	}
	if updateCustomer.Customer_Type != "" {
		customerDetails.Customer_Type = updateCustomer.Customer_Type
	}
	if updateCustomer.Customer_Name != "" {
		customerDetails.Customer_Name = updateCustomer.Customer_Name
	}
	if updateCustomer.Customer_Country_Code != "" && countries.ByName(updateCustomer.Customer_Country_Code) != 0 {
		customerDetails.Customer_Country_Code = updateCustomer.Customer_Country_Code
	}
	if updateCustomer.Customer_State != "" {
		customerDetails.Customer_State = updateCustomer.Customer_State
	}
	if updateCustomer.Customer_Postal_Code != "" {
		customerDetails.Customer_Postal_Code = updateCustomer.Customer_Postal_Code
	}
	if updateCustomer.Customer_City != "" {
		customerDetails.Customer_City = updateCustomer.Customer_City
	}
	if updateCustomer.Customer_Address != "" {
		customerDetails.Customer_Address = updateCustomer.Customer_Address
	}
	if updateCustomer.Customer_Phone != "" {
		customerDetails.Customer_Phone = updateCustomer.Customer_Phone
	}

	db.Save(&customerDetails)

	res, _ := json.Marshal(customerDetails)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}