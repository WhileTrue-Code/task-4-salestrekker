package controller

import (
	"contacts_service/domain"
	"contacts_service/errors"
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
	router.HandleFunc("/get/{id}", controller.GetContactByID).Methods("GET")
	router.HandleFunc("/list", controller.GetAllContacts).Methods("GET")
	router.HandleFunc("/delete/{id}", controller.DeleteContactByID).Methods("DELETE")
	http.Handle("/", router)
}

func (controller *ContactsController) CreateContact(writer http.ResponseWriter, request *http.Request) {

	var contact domain.Input
	err := json.NewDecoder(request.Body).Decode(&contact)
	if err != nil {
		log.Printf("Error in request body json decoding: %s", err)
		http.Error(writer, errors.BadRequestMsg, http.StatusBadRequest)
		return
	}

	err = controller.service.CreateContact(&contact)
	if err != nil {
		if err.Error() == errors.EmptyFieldError {
			http.Error(writer, err.Error(), http.StatusBadRequest)
		} else if err.Error() == errors.ContactAlreadyExist {
			http.Error(writer, err.Error(), http.StatusNotAcceptable)
		} else {
			http.Error(writer, errors.ServerInternalErrorMsg, http.StatusInternalServerError)
		}
		return
	}

	writer.WriteHeader(http.StatusOK)
}

func (controller *ContactsController) GetContactByID(writer http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)

	id := vars["id"]
	contact, err := controller.service.GetOneContactByID(id)
	if err != nil {
		if err.Error() == errors.WrongIdFormatError {
			http.Error(writer, err.Error(), http.StatusNotAcceptable)
		} else {
			http.Error(writer, err.Error(), http.StatusNotFound)
		}
		return
	}

	jsonResponse(contact, writer)
}

func (controller *ContactsController) GetAllContacts(writer http.ResponseWriter, request *http.Request) {
	contacts, err := controller.service.GetAllContacts()
	if err != nil {
		http.Error(writer, errors.ServerInternalErrorMsg, http.StatusInternalServerError)
		return
	}

	if contacts.Results == nil {
		contacts.Results = make([]domain.Input, 0)
	}

	jsonResponse(contacts, writer)
}

func (controller *ContactsController) DeleteContactByID(writer http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)

	id := vars["id"]
	err := controller.service.DeleteOneContactByID(id)
	if err != nil {
		if err.Error() == errors.ContactNotFoundError {
			http.Error(writer, err.Error(), http.StatusNotFound)
		} else if err.Error() == errors.WrongIdFormatError {
			http.Error(writer, err.Error(), http.StatusNotAcceptable)
		} else {
			http.Error(writer, errors.ServerInternalErrorMsg, http.StatusInternalServerError)
		}
		return
	}

	writer.WriteHeader(http.StatusOK)
}
