package handlers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/blackbaudIT/webcore/services"
	"github.com/gorilla/mux"
)

//GetContact responds to an HTTP request for a contact record. It's reliant on an "id" parameter being present in the request's vars.
func (h *WebcoreHandler) GetContact(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	service := &services.ContactService{ContactRepo: h.API}
	contact, err := service.GetContact(vars["id"])

	if err != nil {
		log.Printf("WebcoreHandler.GetContact failed: %s", err)
	}

	data, err := json.Marshal(contact)

	if err != nil {
		log.Printf("WebcoreHandler.GetContact failed to marshal result: %s", err)
	}

	w.Write(data)
}

//GetContactByEmail responds to an HTTP request for a contact record. It's reliant on an "email" parameter being present in the request's vars.
func (h *WebcoreHandler) GetContactByEmail(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	service := &services.ContactService{ContactRepo: h.API}
	contact, err := service.GetContactByEmail(vars["email"])

	if err != nil {
		log.Printf("WebcoreHandler.GetContactByEmail failed: %s", err)
	}

	data, err := json.Marshal(contact)

	if err != nil {
		log.Printf("WebcoreHandler.GetContactByEmail failed to marshal result: %s", err)
	}

	w.Write(data)
}
