package services

import (
	"errors"
	"fmt"

	"github.com/blackbaudIT/webcore/entities"
)

//ContactRepository is an interface for accessing Contact data
type ContactRepository interface {
	ContactQueryBuilder
	GetContact(id string) (*ContactDTO, error)
	QueryContacts(query string) ([]*ContactDTO, error)
	UpdateContact(contact *ContactDTO) error
}

//ContactQueryBuilder is an interface for building Contact queries.
type ContactQueryBuilder interface {
	GetByAuthID(id string) (string, error)
	GetByEmail(email string) (string, error)
}

//ContactDTO is a data transfer object for entities.Contact
type ContactDTO struct {
	//AccountName       string `json:"accountName,omitempty" force:"AccountName__c,omitempty"`
	Salutation      string      `json:"salutation,omitempty" force:"Salutation,omitempty"`
	FirstName       string      `json:"firstName,omitempty" force:"FirstName,omitempty"`
	LastName        string      `json:"lastName,omitempty" force:"LastName,omitempty"`
	SalesForceID    string      `json:"salesForceID,omitempty" force:"Id,omitempty"`
	Email           string      `json:"email,omitempty" force:"Email,omitempty"`
	Phone           string      `json:"phone,omitempty" force:"Phone,omitempty"`
	Fax             string      `json:"fax,omitempty" force:"Fax,omitempty"`
	Title           string      `json:"title,omitempty" force:"Title,omitempty"`
	Account         *AccountDTO `json:"account,omitempty" force:"Account,omitempty"`
	DefaultAccount  string      `json:"defaultAccount,omitempty" force:"Default_Account__c,omitempty"`
	Status          string      `json:"status,omitempty" force:"SFDC_Contact_Status__c,omitempty"`
	Currency        string      `json:"currency,omitempty" force:"CurrencyIsoCode"`
	BBAuthID        string      `json:"bbAuthId,omitempty" force:"BBAuthID__c,omitempty"`
	BBAuthEmail     string      `json:"bbAuthEmail,omitempty" force:"BBAuth_Email__c,omitempty"`
	BBAuthFirstName string      `json:"bbAuthFirstName,omitempty" force:"BBAuth_First_Name__c,omitempty"`
	BBAuthLastName  string      `json:"bbAuthLastName,omitempty" force:"BBAuth_Last_Name__c,omitempty"`
}

//ToEntity converts a ContactDTO into a Contact entity.
func (c *ContactDTO) ToEntity() (*entities.Contact, error) {
	if c.Account == nil {
		return nil, errors.New("Nil value passed as AccountDTO")
	}

	account, err := c.Account.toEntity()

	if err != nil {
		return nil, errors.New("Failed to convert AccountDTO to account")
	}

	name, err := entities.BuildName(c.Salutation, c.FirstName, c.LastName)

	if err != nil {
		return nil, fmt.Errorf("Failed to build name: %s", err)
	}

	contact, err := entities.NewContact(name, account, entities.CurrencyType(c.Currency))

	if err != nil {
		return contact,
			fmt.Errorf("Error converting to contact Entity: %v", err.Error())
	}

	contact.Phone = c.Phone
	contact.Fax = c.Fax
	contact.Title = c.Title
	contact.SetID(c.SalesForceID)
	contact.SetDefaultAccount(c.DefaultAccount)
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
		Salutation:      contact.Name.Salutation,
		FirstName:       contact.Name.FirstName,
		LastName:        contact.Name.LastName(),
		SalesForceID:    contact.ID(),
		Email:           contact.Email(),
		Phone:           contact.Phone,
		Fax:             contact.Fax,
		Title:           contact.Title,
		Account:         ConvertAccountEntityToAccountDTO(contact.Account()),
		DefaultAccount:  contact.DefaultAccount(),
		Status:          contact.Status(),
		BBAuthID:        contact.BBAuthID(),
		BBAuthEmail:     contact.BBAuthEmail(),
		BBAuthFirstName: contact.BBAuthFirstName(),
		BBAuthLastName:  contact.BBAuthLastName(),
	}
	return dto
}

//ContactService provides interaction with Contact data.
type ContactService struct {
	ContactRepo ContactRepository
}

//NewContactService returns a pointer to a valid ContactService given a
//ContactRepository and a ContactQueryBuilder.
func NewContactService(repo ContactRepository) *ContactService {
	return &ContactService{ContactRepo: repo}
}

//GetContact returns a Contact entity by SFDC ID.
func (cs *ContactService) GetContact(id string) (*ContactDTO, error) {
	c, err := cs.ContactRepo.GetContact(id)

	return c, err
}

//GetContactsByEmail returns a slice of contacts that share the same BBAuth email.
func (cs *ContactService) GetContactsByEmail(email string) ([]*ContactDTO, error) {
	query, err := cs.ContactRepo.GetByEmail(email)

	if err != nil {
		return make([]*ContactDTO, 0), err
	}

	contacts, err := cs.ContactRepo.QueryContacts(query)

	return contacts, err
}

//GetContactsByAuthID returns all contact records associated with a given BBAuthID
func (cs *ContactService) GetContactsByAuthID(authID string) ([]*ContactDTO, error) {
	query, err := cs.ContactRepo.GetByAuthID(authID)

	if err != nil {
		return make([]*ContactDTO, 0), err
	}
	contacts, err := cs.ContactRepo.QueryContacts(query)

	return contacts, err
}

//QueryContacts returns all contact records that result from the given query.
func (cs *ContactService) QueryContacts(query string) ([]*ContactDTO, error) {
	contacts, err := cs.ContactRepo.QueryContacts(query)

	return contacts, err
}

//UpdateContact updates a contact..
func (cs *ContactService) UpdateContact(contactDTO *ContactDTO) error {
	contact, err := contactDTO.ToEntity()

	if err != nil {
		return err
	}

	err = cs.ContactRepo.UpdateContact(ConvertContactEntityToContactDTO(contact))
	return err
}
