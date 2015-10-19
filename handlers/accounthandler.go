package handlers

import (
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
