package salesforce

import (
	"errors"
	"fmt"

	"github.com/blackbaudIT/webcore/entities"
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

type SFDCContactQueryBuilder struct {
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

//GetContactCount returns the number of salesforce contacts currently associated with an account.
func (a API) GetContactCount(accountId string) (int, error) {
	queryResponse := &SFDCContactQueryResponse{}
	query := "SELECT count() FROM Contact WHERE AccountId = '" + accountId + "'"

	err := a.client.QuerySFDCObject(query, queryResponse)

	return int(queryResponse.TotalSize), err
}

//QueryContacts returns a slice of SFDCContacts that represents the results of the SOQL query given.
func (a API) QueryContacts(query string) ([]*services.ContactDTO, error) {
	queryResponse := &SFDCContactQueryResponse{}

	err := a.client.QuerySFDCObject(query, queryResponse)

	return queryResponse.Records, err
}

//CreateContact creates a new SFDC Contact.
func (a API) CreateContact(contact *entities.Contact) (string, string, error) {
	dto := services.ConvertContactEntityToContactDTO(contact)

	sfdcContact := SFDCContact{ContactDTO: *dto}
	resp, err := a.client.InsertSFDCObject(sfdcContact)

	if err != nil {
		return "", "", fmt.Errorf("Error creating contact in SFDC: %s", err)
	}

	if !resp.Success {
		return "", "", fmt.Errorf("Error creating contact in SFDC: %s", resp.ErrorMessage)
	}

	newContact := &SFDCContact{}
	err = a.client.GetSFDCObject(resp.ID, newContact)

	if err != nil {
		return "", "", fmt.Errorf("Error getting newly created contact: %s", err)
	}

	return resp.ID, newContact.Name, nil
}

//UpdateContact updates a given contact.
func (a API) UpdateContact(contact *entities.Contact) error {
	return nil
}

//GetByAuthID returns a contact query string that selects contacts with the given
//BBAuthID.
func (b SFDCContactQueryBuilder) GetByAuthID(id string) string {
	query := "SELECT Id, Name, Email, Phone, Fax, Title, AccountId, AccountName__c," +
		"SFDC_Contact_Status__c, CurrencyIsoCode, BBAuthID__c, BBAuth_Email__c, BBAuth_First_Name__c," +
		"BBAuth_Last_Name__c, Account.Name, Account.Id, Account.Clarify_Site_ID__c," +
		"Account.Business_unit__c, Account.Industry, Account.Payer__c," +
		"Account.Billing_street__c, Account.Billing_City__c, Account.Billing_State_Province__c," +
		"Account.Billing_Zip_Postal_Code__c, Account.Billing_Country__c," +
		"Account.Physical_Street__c, Account.Physical_City__c, Account.Physical_State_Province__c," +
		"Account.Physical_Zip_Postal_Code__c, Account.Physical_Country__c FROM Contact " +
		"WHERE BBAuthID__c = '" + id + "'"

	return query
}

//GetByEmail returns a contact query string that selects contacts with the given
//BBAuth Email.
func (b SFDCContactQueryBuilder) GetByEmail(email string) string {
	query := "SELECT Id, Name, Email, Phone, Fax, Title, AccountId, AccountName__c," +
		"SFDC_Contact_Status__c, CurrencyIsoCode, BBAuthID__c, BBAuth_Email__c, BBAuth_First_Name__c," +
		"BBAuth_Last_Name__c, Account.Name, Account.Id, Account.Clarify_Site_ID__c," +
		"Account.Business_unit__c, Account.Industry, Account.Payer__c," +
		"Account.Billing_street__c, Account.Billing_City__c, Account.Billing_State_Province__c," +
		"Account.Billing_Zip_Postal_Code__c, Account.Billing_Country__c," +
		"Account.Physical_Street__c, Account.Physical_City__c, Account.Physical_State_Province__c," +
		"Account.Physical_Zip_Postal_Code__c, Account.Physical_Country__c FROM Contact " +
		"WHERE BBAuth_Email__c = '" + email + "'"

	return query
}
