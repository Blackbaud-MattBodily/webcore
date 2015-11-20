// Package entities defines the domain entity objects and the business rules
// used to create them in a valid state.
package entities

import (
	"errors"
	"fmt"
)

// Account is a Blackbaud Account entity
type Account struct {
	name            string
	siteID          int
	businessUnit    BusinessUnit
	Industry        string
	Payer           string
	PrimaryAddress  *Address
	BillingAddress  *Address
	ShippingAddress *Address
}

// NewAccount creates a valid Account object (with required fields)
func NewAccount(name string) (*Account, error) {
	if name == "" {
		return nil, errors.New("account name cannot be blank")
	}

	return &Account{name: name}, nil
}

// Name of the Account
func (a *Account) Name() string {
	return a.name
}

// SetName will update the name of the account. Can't be an empty string.
func (a *Account) SetName(name string) error {
	if name == "" {
		return errors.New("account name cannot be blank")
	}

	a.name = name
	return nil
}

// SiteID of the Account
func (a *Account) SiteID() int {
	return a.siteID
}

// SetSiteID will update the site ID of the account. Must be a positive integer.
func (a *Account) SetSiteID(siteID int) error {
	if siteID <= 0 {
		return errors.New("siteID must be greater than 0")
	}

	a.siteID = siteID
	return nil
}

// BusinessUnit of the Account
func (a *Account) BusinessUnit() BusinessUnit {
	return a.businessUnit
}

// SetBusinessUnit will update the business unit of the account.
// Restricted to the enum values of BusinessUnit.
func (a *Account) SetBusinessUnit(businessUnit BusinessUnit) error {
	// ok is true if the business unit is in the map
	_, ok := businessUnitValues[string(businessUnit)]

	if !ok {
		return fmt.Errorf("Invalid business unit: %s.", businessUnit)
	}

	a.businessUnit = businessUnit
	return nil
}

// Address block
type Address struct {
	Street  string
	City    string
	State   string
	ZipCode string
	Country string
}

// BusinessUnit is used for enumeration of the Account field
type BusinessUnit string

// BusinessUnit enumeration values
const (
	GMBU BusinessUnit = "GMBU"
	ECBU BusinessUnit = "ECBU"
	IBU  BusinessUnit = "IBU"
)

// map business unit strings to const. used for setting the business unit
var businessUnitValues = map[string]BusinessUnit{
	"GMBU": GMBU,
	"ECBU": ECBU,
	"IBU":  IBU,
}
