package entities

import "errors"

//Contact is a Blackbaud Contact entity.
type Contact struct {
	Name            string
	Email           string
	Phone           string
	Fax             string
	Title           string
	AccountID       string
	AccountName     string
	Status          string
	BBAuthID        string
	BBAuthEmail     string
	BBAuthFirstName string
	BBAuthLastName  string
}

//NewContact creates a valid Contact object with required fields.
func NewContact(name string) (*Contact, error) {
	if name == "" {
		return nil, errors.New("Contact name cannot be blank")
	}

	return &Contact{Name: name}, nil
}

//Account returns the account related to a Contact given an AccountService.
/*func (c *Contact) Account(service *services.AccountService) (*Account, error) {
	account := NewAccount(AccounName)
	accountDTO, err := service.GetAccount(c.AccountID)

	if err != nil {
		return account, err
	}

}*/
