package services

import (
	"fmt"

	"github.com/blackbaudIT/webcore/entities"
)

//ContactRepository is an interface for accessing Contact data
type ContactRepository interface {
	GetContact(id string) (*ContactDTO, error)
	QueryContacts(query string) ([]*ContactDTO, error)
	CreateContact(contact *entities.Contact) (id, name string, err error)
	UpdateContact(contact *entities.Contact) error
}

//ContactDTO is a data transfer object for entities.Contact
type ContactDTO struct {
	Name              string `json:"name,omitempty" force:"Name,omitempty"`
	SalesForceID      string `json:"salesForceID,omitempty" force:"Id,omitempty"`
	Email             string `json:"email,omitempty" force:"Email,omitempty"`
	Phone             string `json:"phone,omitempty" force:"Phone,omitempty"`
	Fax               string `json:"fax,omitempty" force:"Fax,omitempty"`
	Title             string `json:"title,omitempty" force:"Title,omitempty"`
	AccountID         string `json:"accountId,omitempty" force:"AccountId,omitempty"`
	AccountName       string `json:"accountName,omitempty" force:"AccountName__c,omitempty"`
	SFDCContactStatus string `json:"status,omitempty" force:"SFDC_Contact_Status__c,omitempty"`
	BBAuthID          string `json:"bbAuthId,omitempty" force:"BBAuthID__c,omitempty"`
	BBAuthEmail       string `json:"bbAuthEmail,omitempty" force:"BBAuth_Email__c,omitempty"`
	BBAuthFirstName   string `json:"bbFirstName,omitempty" force:"BBAuth_First_Name__c,omitempty"`
	BBAuthLastName    string `json:"bbLastName,omitempty" force:"BBAuth_Last_Name__c,omitempty"`
}

func (c *ContactDTO) toEntity() (*entities.Contact, error) {
	contact, err := entities.NewContact(c.Name)

	if err != nil {
		return contact,
			fmt.Errorf("Error converting to contact Entity: %v", err.Error())
	}

	contact.Email = c.Email
	contact.Phone = c.Phone
	contact.Fax = c.Fax
	contact.Title = c.Title
	contact.AccountID = c.AccountID
	contact.AccountName = c.AccountName
	contact.Status = c.SFDCContactStatus
	contact.BBAuthID = c.BBAuthID
	contact.BBAuthEmail = c.BBAuthEmail
	contact.BBAuthFirstName = c.BBAuthFirstName
	contact.BBAuthLastName = c.BBAuthLastName

	return contact, err
}

//ConvertContactEntityToContactDTO converts an entity.Contact into a ContactDTO.
func ConvertContactEntityToContactDTO(contact *entities.Contact) *ContactDTO {
	dto := &ContactDTO{
		Name:              contact.Name,
		Email:             contact.Email,
		Phone:             contact.Phone,
		Fax:               contact.Fax,
		Title:             contact.Title,
		AccountName:       contact.AccountName,
		SFDCContactStatus: contact.Status,
		BBAuthID:          contact.BBAuthID,
		BBAuthEmail:       contact.BBAuthEmail,
		BBAuthFirstName:   contact.BBAuthFirstName,
		BBAuthLastName:    contact.BBAuthLastName,
	}
	return dto
}

//ContactService provides interaction with Contact data
type ContactService struct {
	ContactRepo ContactRepository
}

//GetContact returns a Contact by SFDC ID or BBAuthID
func (cs *ContactService) GetContact(id string) (*ContactDTO, error) {
	c, err := cs.ContactRepo.GetContact(id)
	return c, err
}

//GetContactByEmail returns a contact from an email.
func (cs *ContactService) GetContactByEmail(email string) (*ContactDTO, error) {
	//TODO: This should probably be based on BBAuth_Email__c, but that field is not indexed yet.
	query := ("SELECT Id, Name, Email, Phone, Fax, Title, AccountId, AccountName__c," +
		"SFDC_Contact_Status__c, BBAuthID__c, BBAuth_Email__c, BBAuth_First_Name__c," +
		"BBAuth_Last_Name__c FROM Contact WHERE Email = '" + email + "'")
	contacts, err := cs.ContactRepo.QueryContacts(query)

	if err != nil {
		return &ContactDTO{}, err
	}

	//TODO: Could probably do something more elegant than returning the first contact in the result slice.
	return contacts[0], err
}

//GetContacts returns all contact records associated with a given BBAuthID
func (cs *ContactService) GetContacts(authID string) ([]*ContactDTO, error) {
	query := ("SELECT Id, Name, Email, Phone, Fax, Title, AccountId, AccountName__c," +
		"SFDC_Contact_Status__c, BBAuthID__c, BBAuth_Email__c, BBAuth_First_Name__c," +
		"BBAuth_Last_Name__c FROM Contact WHERE BBAuthID__c = '" + authID + "'")
	contacts, err := cs.ContactRepo.QueryContacts(query)

	return contacts, err
}

//QueryContacts returns all contact records that result from the given query.
func (cs *ContactService) QueryContacts(query string) ([]*ContactDTO, error) {
	contacts, err := cs.ContactRepo.QueryContacts(query)

	return contacts, err
}

//CreateContact creates a new Contact
func (cs *ContactService) CreateContact(c ContactDTO) (id, name string, err error) {
	contact, err := c.toEntity()

	if err != nil {
		return "", "", err
	}

	id, name, err = cs.ContactRepo.CreateContact(contact)

	return id, name, err
}

//UpdateContact updates a contact represented by a ContactDTO.
func (cs *ContactService) UpdateContact(c ContactDTO) error {
	contact, err := c.toEntity()

	if err != nil {
		return err
	}

	err = cs.ContactRepo.UpdateContact(contact)
	return err
}
