package controller

import (
	"contacts_service/domain"
	"encoding/json"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

type ContactsController struct {
	service domain.ContactService
}

func NewContactsController(service domain.ContactService) *ContactsController {
	return &ContactsController{service: service}
}

func (controller *ContactsController) Init(router *mux.Router) {
	router.HandleFunc("/update", controller.CreateContact).Methods("POST")
	http.Handle("/", router)
}

func (controller *ContactsController) CreateContact(writer http.ResponseWriter, request *http.Request) {

	var contact domain.Input

	err := json.NewDecoder(request.Body).Decode(&contact)
	if err != nil {
		log.Printf("Error in request body json decoding: %s", err)
		writer.WriteHeader(http.StatusBadRequest)
		return
	}

	err = controller.service.CreateContact(&contact)
	if err != nil {
		log.Printf("Error in creating contact because of: %s", err)
		writer.WriteHeader(http.StatusInternalServerError)
		return
	}

	writer.WriteHeader(http.StatusOK)
	return
}
