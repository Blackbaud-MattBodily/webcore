package salesforce

import (
	"errors"
	"fmt"
	"regexp"

	"github.com/blackbaudIT/webcore/services"
)

//SFDCContact wraps the ContactDTO so that SFDC fields can be mapped onto it.
type SFDCContact struct {
	services.ContactDTO
}

//SFDCContactQueryResponse wraps the base SFDCQueryResponse and attaches a slice of SFDCContact pointers which will be written into.
type SFDCContactQueryResponse struct {
	SFDCQueryResponse

	Records []*services.ContactDTO `json:"Records" force:"records"`
}

//ApiName is the SFDC ApiName of the Contact object.
func (s SFDCContact) ApiName() string {
	return "Contact"
}

//ExternalIdApiName is the SFDC external id for the Contact object.
func (s SFDCContact) ExternalIdApiName() string {
	return "eBus_Contact_ID__c"
}

//GetContact returns a Salesforce contact given an SFDC ID or a BBAuthID.
func (a API) GetContact(id string) (*services.ContactDTO, error) {
	var err error
	err = nil

	contact := &SFDCContact{}

	if id == "" {
		return nil, errors.New("id cannot be an empty string")
	}

	if len(id) == 15 || len(id) == 18 {
		err = a.client.GetSFDCObject(id, contact)
	} else {
		return nil, errors.New("id must be a valid 15 or 18 character SFDC id")
	}

	if err != nil {
		return &contact.ContactDTO, fmt.Errorf("Error querying SFDC: %s \n", err)
	}
	return &contact.ContactDTO, nil
}

//QueryContacts returns a slice of SFDCContacts that represents the results of the SOQL query given.
func (a API) QueryContacts(query string) ([]*services.ContactDTO, error) {
	queryResponse := &SFDCContactQueryResponse{}

	err := a.client.QuerySFDCObject(query, queryResponse)

	//This is kind of ugly, but I don't know of a better way to get this information
	//on the AccountDTO.
	for _, contact := range queryResponse.Records {
		//TODO: Come back and add error handling.
		count, _ := a.GetContactCount(contact.Account.SalesForceID)
		contact.Account.ContactCount = count
	}

	return queryResponse.Records, err
}

//GetByAuthID returns a contact query string that selects contacts with the given
//BBAuthID.
func (a API) GetByAuthID(id string) (string, error) {
	match, err := regexp.MatchString("[A-Za-z0-9]{8}-([A-Za-z0-9]{4}-){3}[A-Za-z0-9]{12}", id)

	if err != nil || !match {
		return "", fmt.Errorf("BBAuthID incorrectly formatted: %s", err)
	}

	query := "SELECT Id, FirstName, LastName, Email, Phone, Fax, Title, AccountId, AccountName__c," +
		"SFDC_Contact_Status__c, CurrencyIsoCode, BBAuthID__c, BBAuth_Email__c, BBAuth_First_Name__c," +
		"BBAuth_Last_Name__c, Account.Name, Account.Id, Account.Clarify_Site_ID__c," +
		"Account.Business_unit__c, Account.Industry, Account.Payer__c," +
		"Account.Billing_street__c, Account.Billing_City__c, Account.Billing_State_Province__c," +
		"Account.Billing_Zip_Postal_Code__c, Account.Billing_Country__c," +
		"Account.Physical_Street__c, Account.Physical_City__c, Account.Physical_State_Province__c," +
		"Account.Physical_Zip_Postal_Code__c, Account.Physical_Country__c FROM Contact " +
		"WHERE BBAuthID__c = '" + id + "'"

	return query, nil
}

//GetByEmail returns a contact query string that selects contacts with the given
//BBAuth Email.
func (a API) GetByEmail(email string) (string, error) {
	match, err := regexp.MatchString(".+@.+", email)

	if err != nil || !match {
		return "", fmt.Errorf("Email incorrectly formatted: %s", err)
	}

	query := "SELECT Id, FirstName, LastName, Email, Phone, Fax, Title, AccountId, AccountName__c," +
		"SFDC_Contact_Status__c, CurrencyIsoCode, BBAuthID__c, BBAuth_Email__c, BBAuth_First_Name__c," +
		"BBAuth_Last_Name__c, Account.Name, Account.Id, Account.Clarify_Site_ID__c," +
		"Account.Business_unit__c, Account.Industry, Account.Payer__c," +
		"Account.Billing_street__c, Account.Billing_City__c, Account.Billing_State_Province__c," +
		"Account.Billing_Zip_Postal_Code__c, Account.Billing_Country__c," +
		"Account.Physical_Street__c, Account.Physical_City__c, Account.Physical_State_Province__c," +
		"Account.Physical_Zip_Postal_Code__c, Account.Physical_Country__c FROM Contact " +
		"WHERE BBAuth_Email__c = '" + email + "'"

	return query, nil
}

//UpdateContact updates a given contact.
func (a API) UpdateContact(contact *services.ContactDTO) error {
	//This is a bit weird, but we can't update a record if the ID is part of the
	//object and updates fail whenever we try and set the Account field on the
	//contact object. Since we don't have a reason yet to update a contact with
	//a new account, this isn't majorly impacting. However, in the future we'll
	//need to figure out a way around this.
	sfdcContact := SFDCContact{ContactDTO: *contact}
	id := sfdcContact.SalesForceID
	sfdcContact.SalesForceID = ""
	sfdcContact.Account = nil

	return a.client.UpdateSFDCObject(id, sfdcContact)
}
