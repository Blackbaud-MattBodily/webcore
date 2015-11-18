package entities

import "errors"

//Required fields.
//Name, Account, Contact Currency, maybe record type.

//Contact is a Blackbaud Contact entity.
type Contact struct {
	id              string
	name            string
	email           string
	Phone           string
	Fax             string
	Title           string
	account         *Account
	status          string
	Currency        CurrencyType
	bbAuthID        string
	bbAuthEmail     string
	bbAuthFirstName string
	bbAuthLastName  string
}

//CurrencyType is an enum for setting a contact's preferred currency.
type CurrencyType string

//Enumeration for CurrencyType.
const (
	USD CurrencyType = "USD - U.S. Dollar"
	CAD CurrencyType = "CAD - Canadian Dollar"
	EUR CurrencyType = "EUR - Euro"
	GBP CurrencyType = "GBP - British Pound"
	JPY CurrencyType = "JPY - Japanese Yen"
	AUD CurrencyType = "AUD - Australian Dollar"
)

//NewContact creates a valid Contact object with required fields.
func NewContact(name string, account *Account, currency CurrencyType) (*Contact, error) {
	if name == "" {
		return nil, errors.New("Contact name cannot be blank")
	}

	if account == nil {
		return nil, errors.New("Contact must have an account")
	}

	if currency == "" {
		return nil, errors.New("Contact must have a currency type")
	}

	return &Contact{name: name, account: account, Currency: currency}, nil
}

//ID of the contact.
func (c *Contact) ID() string {
	return c.id
}

//Name of the contact.
func (c *Contact) Name() string {
	return c.name
}

//Email of the contact.
func (c *Contact) Email() string {
	return c.email
}

//Account of the contact.
func (c *Contact) Account() *Account {
	return c.account
}

//Status of the cotnact.
func (c *Contact) Status() string {
	return c.status
}

//BBAuthID of the contact.
func (c *Contact) BBAuthID() string {
	return c.bbAuthID
}

//BBAuthEmail of the contact.
func (c *Contact) BBAuthEmail() string {
	return c.bbAuthEmail
}

//BBAuthFirstName of the contact.
func (c *Contact) BBAuthFirstName() string {
	return c.bbAuthFirstName
}

//BBAuthLastName of the contact.
func (c *Contact) BBAuthLastName() string {
	return c.bbAuthLastName
}

//SetName sets the contact's name.
func (c *Contact) SetName(name string) error {
	if name == "" {
		return errors.New("Contact must have a non-empty name")
	}

	c.name = name
	return nil
}

//SetID sets the contact's ID.
func (c *Contact) SetID(id string) error {
	c.id = id
	return nil
}

//SetEmail sets the contact's email.
func (c *Contact) SetEmail(email string) error {
	c.email = email
	return nil
}

//SetAccount sets the contact's account.
func (c *Contact) SetAccount(account *Account) error {
	c.account = account
	return nil
}

//SetStatus sets the contact's status.
func (c *Contact) SetStatus(status string) error {
	c.status = status
	return nil
}

//SetBBAuthID sets the contact's BBAuthID.
func (c *Contact) SetBBAuthID(bbAuthID string) error {
	c.bbAuthID = bbAuthID
	return nil
}

//SetBBAuthEmail sets the contact's BBAuthEmail.
func (c *Contact) SetBBAuthEmail(bbAuthEmail string) error {
	c.bbAuthEmail = bbAuthEmail
	return nil
}

//SetBBAuthFirstName sets the contact's BBAuthFirstname.
func (c *Contact) SetBBAuthFirstName(bbAuthFirstName string) error {
	c.bbAuthFirstName = bbAuthFirstName
	return nil
}

//SetBBAuthLastName sets the contact's BBAuthLastName.
func (c *Contact) SetBBAuthLastName(bbAuthLastName string) error {
	c.bbAuthLastName = bbAuthLastName
	return nil
}
