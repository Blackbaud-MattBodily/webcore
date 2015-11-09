package services

import (
	"errors"
	"fmt"

	"github.com/blackbaudIT/webcore/entities"
)

//ContactRepository is an interface for accessing Contact data
type ContactRepository interface {
	GetContact(id string) (*ContactDTO, error)
	GetContactsByAuthID(id string) ([]*ContactDTO, error)
	GetContactsByEmail(email string) ([]*ContactDTO, error)
	QueryContacts(query string) ([]*ContactDTO, error)
	GetContactCount(accountID string) (int, error)
	CreateContact(contact *entities.Contact) (id, name string, err error)
	UpdateContact(contact *entities.Contact) error
}

//ContactDTO is a data transfer object for entities.Contact
type ContactDTO struct {
	//AccountName       string `json:"accountName,omitempty" force:"AccountName__c,omitempty"`
	Name            string      `json:"name,omitempty" force:"Name,omitempty"`
	SalesForceID    string      `json:"salesForceID,omitempty" force:"Id,omitempty"`
	Email           string      `json:"email,omitempty" force:"Email,omitempty"`
	Phone           string      `json:"phone,omitempty" force:"Phone,omitempty"`
	Fax             string      `json:"fax,omitempty" force:"Fax,omitempty"`
	Title           string      `json:"title,omitempty" force:"Title,omitempty"`
	Account         *AccountDTO `json:"account,omitempty" force:"Account,omitempty"`
	Status          string      `json:"status,omitempty" force:"SFDC_Contact_Status__c,omitempty"`
	Currency        string      `json:"currency,omitempty" force:"CurrencyIsoCode"`
	BBAuthID        string      `json:"bbAuthId,omitempty" force:"BBAuthID__c,omitempty"`
	BBAuthEmail     string      `json:"bbAuthEmail,omitempty" force:"BBAuth_Email__c,omitempty"`
	BBAuthFirstName string      `json:"bbAuthFirstName,omitempty" force:"BBAuth_First_Name__c,omitempty"`
	BBAuthLastName  string      `json:"bbAuthLastName,omitempty" force:"BBAuth_Last_Name__c,omitempty"`
}

func (c *ContactDTO) toEntity() (*entities.Contact, error) {
	account, err := c.Account.toEntity()

	if err != nil {
		return nil, errors.New("Failed to convert AccountDTO to account")
	}

	contact, err := entities.NewContact(c.Name, account, entities.CurrencyType(c.Currency))

	if err != nil {
		return contact,
			fmt.Errorf("Error converting to contact Entity: %v", err.Error())
	}

	contact.Phone = c.Phone
	contact.Fax = c.Fax
	contact.Title = c.Title
	contact.SetEmail(c.Email)
	contact.SetStatus(c.Status)
	contact.SetBBAuthID(c.BBAuthID)
	contact.SetBBAuthEmail(c.BBAuthEmail)
	contact.SetBBAuthFirstName(c.BBAuthFirstName)
	contact.SetBBAuthLastName(c.BBAuthLastName)

	return contact, err
}

//ConvertContactEntityToContactDTO converts an entity.Contact into a ContactDTO.
func ConvertContactEntityToContactDTO(contact *entities.Contact) *ContactDTO {
	dto := &ContactDTO{
		Name:            contact.Name(),
		Email:           contact.Email(),
		Phone:           contact.Phone,
		Fax:             contact.Fax,
		Title:           contact.Title,
		Account:         ConvertAccountEntityToAccountDTO(contact.Account()),
		Status:          contact.Status(),
		BBAuthID:        contact.BBAuthID(),
		BBAuthEmail:     contact.BBAuthEmail(),
		BBAuthFirstName: contact.BBAuthFirstName(),
		BBAuthLastName:  contact.BBAuthLastName(),
	}
	return dto
}

//ContactService provides interaction with Contact data
type ContactService struct {
	ContactRepo ContactRepository
}

//GetContact returns a Contact by SFDC ID.
func (cs *ContactService) GetContact(id string) (*ContactDTO, error) {
	c, err := cs.ContactRepo.GetContact(id)
	return c, err
}

//GetContactsByEmail returns a slice of contacts that share the same BBAuth email.
func (cs *ContactService) GetContactsByEmail(email string) ([]*ContactDTO, error) {
	contacts, err := cs.ContactRepo.GetContactsByEmail(email)

	return contacts, err
}

//GetContactsByAuthID returns all contact records associated with a given BBAuthID
func (cs *ContactService) GetContactsByAuthID(authID string) ([]*ContactDTO, error) {
	contacts, err := cs.ContactRepo.GetContactsByAuthID(authID)

	if err != nil {
		fmt.Println(err)
	}

	return contacts, err
}

//GetContactCount returns the number of contacts currently associated with an account.
func (cs *ContactService) GetContactCount(accountID string) (int, error) {
	count, err := cs.ContactRepo.GetContactCount(accountID)

	return count, err
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
