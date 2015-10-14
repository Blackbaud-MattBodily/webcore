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

//SFDCContactQueryResponse contains the response for contact queries.
type SFDCContactQueryResponse struct {
	SFDCQueryResponse

	Records []*SFDCContact `json:"Records" force:"records"`
}

//ApiName is the SFDC ApiName of the Contact object.
func (s SFDCContact) ApiName() string {
	return "Contact"
}

//ExternalIdApiName is the SFDC external id for the Contact object.
func (s SFDCContact) ExternalIdApiName() string {
	return "Username"
}

//GetContact returns a Salesforce contact given an SFDC ID or a BBAuthID.
func (a API) GetContact(id string) (*services.ContactDTO, error) {
	contact := &SFDCContact{}
	contactLookupFunc, err := a.getForceAPIContactLookupFunction(id)

	if err != nil {
		return &contact.ContactDTO, fmt.Errorf("Error validating id parameter: %s \n", err)
	}

	err = contactLookupFunc(contact)
	fmt.Println(contact)
	if err != nil {
		return &contact.ContactDTO, fmt.Errorf("Error querying SFDC: %s \n", err)
	}
	return &contact.ContactDTO, nil
}

func (a API) getForceAPIContactLookupFunction(id string) (func(*SFDCContact) error, error) {
	if id == "" {
		return nil, errors.New("id cannot be an empty string")
	}

	if len(id) == 15 || len(id) == 18 {
		return func(contact *SFDCContact) error {
			err := a.client.GetSFDCObject(id, contact)
			return err
		}, nil
	}

	return func(contact *SFDCContact) error {
		//Put the contact pointer into the inital Records slice to eliminate the need for a deep copy.
		records := make([]*SFDCContact, 1)
		records[0] = contact
		queryResponse := &SFDCContactQueryResponse{Records: records}

		query := ("SELECT Id, Name, Email, Phone, Fax, Title, AccountId, AccountName__c," +
			"SFDC_Contact_Status__c, BBAuthID__c, BBAuth_Email__c, BBAuth_First_Name__c," +
			"BBAuth_Last_Name__c FROM Contact WHERE BBAuthID__c = '" + id + "'")

		err := a.client.QuerySFDCObject(query, queryResponse)

		if err != nil {
			fmt.Println(err)
			return err
		}

		if queryResponse.TotalSize > 0 {
			err = errors.New("The ID provided did not resolve to any contact records.")
		}
		return err
	}, nil
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
