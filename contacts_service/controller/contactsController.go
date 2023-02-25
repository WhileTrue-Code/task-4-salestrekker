package controller

import (
	"github.com/gorilla/mux"
	"net/http"
)

type ContactsController struct {
	//TODO items
}

func NewContactsController() *ContactsController {
	return &ContactsController{}
}

func (controller *ContactsController) Init(router *mux.Router) {
	router.HandleFunc("/testRoute", controller.NewMethod).Methods("GET")
	http.Handle("/", router)
}

func (controller *ContactsController) NewMethod(writer http.ResponseWriter, request *http.Request) {
	writer.WriteHeader(http.StatusOK)
	return
}
