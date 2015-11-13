package handlers

import (
	"log"
	"net/http"
	"strconv"

	"github.com/blackbaudIT/webcore/Godeps/_workspace/src/github.com/gorilla/mux"
	"github.com/blackbaudIT/webcore/services"
)

//AccountHandler holds an AccountRepository and uses it to handle standard
//http requests related to Accounts.
type AccountHandler struct {
	accountRepo services.AccountRepository
}

//NewAccountHandler creates a new AccountHandler given an AccountRepository
func NewAccountHandler(repo services.AccountRepository) *AccountHandler {
	return &AccountHandler{accountRepo: repo}
}

//GetContactCount returns the number of contacts currently related to a given account.
func (h *AccountHandler) GetContactCount(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	service := &services.AccountService{AccountRepo: h.accountRepo}
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
