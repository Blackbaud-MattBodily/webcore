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

//ApiName is the SFDC ApiName of the Contact object.
func (s SFDCContact) ApiName() string {
	return "Contact"
}

//ExternalIdApiName is the SFDC external id for the Contact object.
func (s SFDCContact) ExternalIdApiName() string {
	return "Username"
}

//GetContact returns a Salesforce contact for the ID specified.
func (a API) GetContact(id string) (*services.ContactDTO, error) {
	contact := &SFDCContact{}
	contactLookupFunc, err := a.getForceAPIContactLookupFunction(id)

	if err != nil {
		return &contact.ContactDTO, fmt.Errorf("Error validating id parameter: %s \n", err)
	}

	err = contactLookupFunc(contact)

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

	return nil, errors.New("SFDC Id's must be an alphanumeric string with a" +
		"length of 15 or 18")
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
