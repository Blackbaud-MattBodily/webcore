package main

import (
	"encoding/json"
	"fmt"

	"github.com/blackbaudIT/webcore/data/salesforce"
	"github.com/blackbaudIT/webcore/services"
)

var api = salesforce.NewAPI()
var service = services.AccountService{AccountRepo: api}

func main() {
	fmt.Println("starting...")
	fmt.Println("")
	getAccountExample()
	//insertAccountExample()
	//updateAccountExample()
}

func getAccountExample() {
	account, err := service.GetAccount("5740")

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
