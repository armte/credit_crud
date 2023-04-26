package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
	"github.com/armte/credit_crud/pkg/utils"
	"github.com/armte/credit_crud/pkg/models"
)

var NewAccount models.Account

func GetAccount(w http.ResponseWriter, r *http.Request) {
	newAccounts := models.GetAllAccounts()
	res, _ := json.Marshal(newAccounts)
	w.Header().Set("Content-Type","application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func GetAccountById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	accountId := vars["accountId"]
	Id, err := strconv.ParseInt(accountId,0,0)
	if err != nil {
		fmt.Println("error while parsing")
		return
	}

	accountDetails, _, errGet := models.GetAccountById(Id)
	if errGet != nil {
		http.Error(w, "404 Not Found: " + errGet.Error(), http.StatusNotFound)
		return
	}

	res, _ := json.Marshal(accountDetails)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func CreateAccount(w http.ResponseWriter, r *http.Request) {
	CreateAccount := &models.Account{}
	utils.ParseBody(r, CreateAccount)
	c, err := CreateAccount.CreateAccount()
	if err != nil {
		http.Error(w, "400 Bad Request: " + err.Error(), http.StatusBadRequest)
	} else {
		res, _ := json.Marshal(c)
		w.WriteHeader(http.StatusCreated)
		w.Write(res)
	}
}

func DeleteAccount(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	accountId := vars["accountId"]
	Id, err := strconv.ParseInt(accountId, 0, 0)
	if err != nil {
		fmt.Println("error while parsing")
	}
	account, errGet := models.DeleteAccount(Id)
	if errGet != nil {
		http.Error(w, "404 Not Found: " + errGet.Error(), http.StatusNotFound)
		return
	}
	res, _ := json.Marshal(account)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func UpdateAccount(w http.ResponseWriter, r *http.Request) {
	updateAccount := &models.Account{}
	utils.ParseBody(r, updateAccount)
	vars := mux.Vars(r)
	accountId := vars["accountId"]
	Id, err := strconv.ParseInt(accountId, 0, 0)
	if err != nil {
		fmt.Println("error while parsing")
	}

	accountDetails, db, errGet := models.GetAccountById(Id)

	if errGet != nil {
		http.Error(w, "404 Not Found: " + errGet.Error(), http.StatusNotFound)
		return
	}

	if updateAccount.Account_Type != "" {
		accountDetails.Account_Type = updateAccount.Account_Type
	}
	if updateAccount.Account_State != "" {
		accountDetails.Account_State = updateAccount.Account_State
	}
	if updateAccount.Account_Postal_Code != "" {
		accountDetails.Account_Postal_Code = updateAccount.Account_Postal_Code
	}
	if updateAccount.Account_Country_Code != "" {
		accountDetails.Account_Country_Code = updateAccount.Account_Country_Code
	}
	if updateAccount.Customer_Number != 0 {
		accountDetails.Customer_Number = updateAccount.Customer_Number
	}
	

	db.Save(&accountDetails)

	res, _ := json.Marshal(accountDetails)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}