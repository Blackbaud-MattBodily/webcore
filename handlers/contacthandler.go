package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/blackbaudIT/webcore/Godeps/_workspace/src/github.com/gorilla/mux"
	"github.com/blackbaudIT/webcore/services"
)

//ContactHandler holds a ContactRepository and uses it to respond to standard
//http requests related to Contacts.
type ContactHandler struct {
	contactRepo services.ContactRepository
}

//NewContactHandler creates a new ContactHandler using a given ContactRepository.
func NewContactHandler(repo services.ContactRepository) *ContactHandler {
	return &ContactHandler{contactRepo: repo}
}

//GetContact responds to an HTTP request for a contact record. It's reliant on an "id" parameter being present in the request's vars.
func (h *ContactHandler) GetContact(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	service := &services.ContactService{ContactRepo: h.contactRepo}
	contact, err := service.GetContact(vars["id"])

	if err != nil {
		log.Printf("ContactHandler.GetContact failed: %s", err)
	}

	data, err := json.Marshal(contact)

	if err != nil {
		log.Printf("ContactHandler.GetContact failed to marshal result: %s", err)
	}

	w.Write(data)
}

//GetContactsByEmail responds to an HTTP request for a contact record. It's reliant on an "email" parameter being present in the request's vars.
func (h *ContactHandler) GetContactsByEmail(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	service := &services.ContactService{ContactRepo: h.contactRepo}
	contact, err := service.GetContactsByEmail(vars["email"])

	if err != nil {
		log.Printf("ContactHandler.GetContactsByEmail failed: %s", err)
	}

	data, err := json.Marshal(contact)

	if err != nil {
		log.Printf("ContactHandler.GetContactsByEmail failed to marshal result: %s", err)
	}

	w.Write(data)
}

//GetContactsByAuthID responds to an HTTP request for all contact records associated with a given BBAuthID.
func (h *ContactHandler) GetContactsByAuthID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	service := &services.ContactService{ContactRepo: h.contactRepo}
	contacts, err := service.GetContactsByAuthID(vars["authID"])

	if err != nil {
		log.Printf("ContactHandler.GetContactsByAuthID failed: %s", err)
	}

	data, err := json.Marshal(contacts)

	if err != nil {
		log.Printf("ContactHandler.GetContactsByAuthID failed to marshal result: %s", err)
	}

	w.Write(data)
}

//GetContactCount returns the number of contacts currently related to a given account.
func (h *ContactHandler) GetContactCount(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	service := &services.ContactService{ContactRepo: h.contactRepo}
	w.Header().Set("Content-type", "application/json")
	count, err := service.GetContactCount(vars["accountId"])

	if err != nil {
		log.Printf("ContactHandler.GetContactCount failed to retrieve count: %s", err)
		w.Write([]byte("{\"count\":0}"))
		return
	}

	response := "{\"count\":" + strconv.Itoa(count) + "}"
	w.Write([]byte(response))
}
