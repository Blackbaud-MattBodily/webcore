package main

import (
	"encoding/json"
	"fmt"

	"github.com/blackbaudIT/webcore/data/salesforce"
	"github.com/blackbaudIT/webcore/services"
)

var api = salesforce.NewAPI()
var service = services.AccountService{AccountRepo: api}
var contactService = services.ContactService{ContactRepo: api}

func main() {
	fmt.Println("starting...")
	fmt.Println("")

	getContactsWithAccountExample()
	//getContactCountExample()
	//getContactExample()
	//getAccountExample()
	//insertAccountExample()
	//updateAccountExample()
}

func getContactsWithAccountExample() {
	contacts, err := contactService.GetContactsByAuthID("32FBC72D-C0FE-4B50-B0F4-EDCEFD7B4DEF")

	if err != nil {
		fmt.Println(err)
	}

	data, err := json.Marshal(contacts)

	fmt.Printf("Contacts: %s", data)
}

func getContactCountExample() {
	count, err := service.GetContactCount("001d000001TweFmAAJ")

	if err != nil {
		fmt.Println(err)
		fmt.Println()
	}

	fmt.Println(count)
}

func getContactExample() {
	contact, err := contactService.GetContact("003d0000027LKPQ")

	if err != nil {
		fmt.Println(err)
		fmt.Println("")
	}

	data, err := json.Marshal(contact)

	if err != nil {
		fmt.Println(err)
		fmt.Println("")
	}

	fmt.Printf("Contact: %s", data)
	fmt.Println("")
}

func getAccountExample() {
	account, err := service.GetAccount("46558")

	if err != nil {
		fmt.Println(err)
		fmt.Println("")
	}

	data, err := json.Marshal(account)

	if err != nil {
		fmt.Println(err)
		fmt.Println("")
	}

	fmt.Printf("Account: %s", data)
	fmt.Println("")
}

func insertAccountExample() {
	dto := services.AccountDTO{
		Name:           "Integration Testing Account 2",
		SiteID:         "1111111",
		ShippingStreet: "789 Main St",
	}

	id, siteID, err := service.CreateAccount(dto)

	if err != nil {
		fmt.Println(err)
		fmt.Println("")
	} else {
		fmt.Printf("successfully created account (SFDC ID: %s, SiteID: %d", id, siteID)
		fmt.Println("")
	}
}

func updateAccountExample() {
	dto := services.AccountDTO{
		Name:           "Integration Testing Account",
		SiteID:         "93275",
		ShippingStreet: "456 Main St",
	}

	err := service.UpdateAccount(dto)

	if err != nil {
		fmt.Println(err)
		fmt.Println("")
	} else {
		fmt.Println("successfully updated account")
		fmt.Println("")
	}
}
