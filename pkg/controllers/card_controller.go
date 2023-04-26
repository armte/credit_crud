package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"net/http"
	"github.com/armte/credit_crud/pkg/utils"
	"github.com/armte/credit_crud/pkg/models"
)

var NewCard models.Card

func GetCard(w http.ResponseWriter, r *http.Request) {
	newCards := models.GetAllCards()
	res, _ := json.Marshal(newCards)
	w.Header().Set("Content-Type","application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func GetCardById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	cardId := vars["cardId"]
	Id, err := uuid.Parse(cardId)
	if err != nil {
		fmt.Println("error while parsing")
		return
	}

	cardDetails, _, errGet := models.GetCardById(Id)
	if errGet != nil {
		http.Error(w, "404 Not Found: " + errGet.Error(), http.StatusNotFound)
		return
	}

	res, _ := json.Marshal(cardDetails)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func CreateCard(w http.ResponseWriter, r *http.Request) {
	CreateCard := &models.Card{}
	utils.ParseBody(r, CreateCard)
	c, err := CreateCard.CreateCard()
	if err != nil {
		http.Error(w, "400 Bad Request: " + err.Error(), http.StatusBadRequest)
	} else {
		res, _ := json.Marshal(c)
		w.WriteHeader(http.StatusCreated)
		w.Write(res)
	}
}

func DeleteCard(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	cardId := vars["cardId"]
	Id, err := uuid.Parse(cardId)
	if err != nil {
		fmt.Println("error while parsing")
	}
	card, errGet := models.DeleteCard(Id)
	if errGet != nil {
		http.Error(w, "404 Not Found: " + errGet.Error(), http.StatusNotFound)
		return
	}
	res, _ := json.Marshal(card)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func UpdateCard(w http.ResponseWriter, r *http.Request) {
	updateCard := &models.Card{}
	utils.ParseBody(r, updateCard)
	vars := mux.Vars(r)
	cardId := vars["cardId"]
	Id, err := uuid.Parse(cardId)
	if err != nil {
		fmt.Println("error while parsing")
	}

	cardDetails, db, errGet := models.GetCardById(Id)

	if errGet != nil {
		http.Error(w, "404 Not Found: " + errGet.Error(), http.StatusNotFound)
		return
	}

	if updateCard.Credit_Limit != 0 {
		cardDetails.Credit_Limit = updateCard.Credit_Limit
	}
	if updateCard.Account_Number != 0 {
		cardDetails.Account_Number = updateCard.Account_Number
	}
	
	db.Save(&cardDetails)

	res, _ := json.Marshal(cardDetails)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}