package entities

import "errors"

//Required fields.
//Name, Account, Contact Currency, maybe record type.

//Contact is a Blackbaud Contact entity.
type Contact struct {
	id              string
	Name            *Name
	email           string
	Phone           string
	Fax             string
	Title           string
	account         *Account
	defaultAccount  string
	roles           []ContactRole
	status          string
	Currency        CurrencyType
	bbAuthID        string
	bbAuthEmail     string
	bbAuthFirstName string
	bbAuthLastName  string
}

//ContactRole is a role for a Blackbaud Contact entity.
type ContactRole struct {
	RoleName   string
	RoleType   string
	RoleStatus string
}

//Name represents the salutation, first name, and last name of a contact.
type Name struct {
	Salutation string
	FirstName  string
	lastName   string
}

//BuildName builds a Name type out of a firstname, lastname, and salutation.
func BuildName(salutation, firstname, lastname string) (*Name, error) {
	if len(lastname) > 0 {
		return &Name{Salutation: salutation, FirstName: firstname, lastName: lastname}, nil
	}
	return nil, errors.New("A last name is required to build a Name struct.")
}

//LastName of a Name struct.
func (n *Name) LastName() string {
	return n.lastName
}

//SetLastName sets the lastName of a Name struct.
func (n *Name) SetLastName(lastName string) error {
	if len(lastName) == 0 {
		return errors.New("Lastname can not be empty.")
	}
	n.lastName = lastName
	return nil
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
func NewContact(name *Name, account *Account, currency CurrencyType) (*Contact, error) {
	if len(name.LastName()) == 0 {
		return nil, errors.New("Contact's Name must have a lastName")
	}
	if account == nil {
		return nil, errors.New("Contact must have an account")
	}

	if currency == "" {
		return nil, errors.New("Contact must have a currency type")
	}

	return &Contact{Name: name, account: account, Currency: currency}, nil
}

//ID of the contact.
func (c *Contact) ID() string {
	return c.id
}

//Email of the contact.
func (c *Contact) Email() string {
	return c.email
}

//Account of the contact.
func (c *Contact) Account() *Account {
	return c.account
}

//DefaultAccount of the contact.
func (c *Contact) DefaultAccount() string {
	return c.defaultAccount
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

//SetDefaultAccount sets the contact's default account.
func (c *Contact) SetDefaultAccount(id string) error {
	c.defaultAccount = id
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
