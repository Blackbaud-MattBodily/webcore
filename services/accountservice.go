// Package services provides services and data transfer objects that can be
// used to interact with entities package.
package services

import (
	"fmt"
	"strconv"

	"github.com/blackbaudIT/webcore/entities"
)

// AccountRepository is an inteface for accessing Account data
type AccountRepository interface {
	GetAccount(id string) (*AccountDTO, error)
	QueryAccounts(query string) ([]*AccountDTO, error)
	CreateAccount(account *entities.Account) (id string, siteID int, err error)
	UpdateAccount(account *entities.Account) error
}

// AccountDTO is an data transfer object for entities.Account
type AccountDTO struct {
	Name            string `json:"name,omitempty" force:"Name,omitempty"`
	SalesForceID    string `json:"salesForceID,omitempty" force:"Id,omitempty"`
	SiteID          string `json:"siteId,omitempty" force:"Clarify_Site_ID__c,omitempty"`
	BusinessUnit    string `json:"businessUnit,omitempty" force:"Business_Unit__c,omitempty"`
	Industry        string `json:"industry,omitempty" force:"Industry,omitempty"`
	Payer           string `json:"payer,omitempty" force:"Payer__c,omitempty"`
	PrimaryStreet   string `json:"primaryStreet,omitempty" force:"-"`
	PrimaryCity     string `json:"primaryCity,omitempty" force:"-"`
	PrimaryState    string `json:"primaryState,omitempty" force:"-"`
	PrimaryZipCode  string `json:"primaryZipCode,omitempty" force:"-"`
	PrimaryCountry  string `json:"primaryCountry,omitempty" force:"-"`
	BillingStreet   string `json:"billingStreet,omitempty" force:"Billing_Street__c,omitempty"`
	BillingCity     string `json:"billingCity,omitempty" force:"Billing_City__c,omitempty"`
	BillingState    string `json:"billingState,omitempty" force:"Billing_State_Province__c,omitempty"`
	BillingZipCode  string `json:"billingZipCode,omitempty" force:"Billing_Zip_Postal_Code__c,omitempty"`
	BillingCountry  string `json:"billingCountry,omitempty" force:"Billing_Country__c,omitempty"`
	ShippingStreet  string `json:"shippingStreet,omitempty" force:"Physical_Street__c,omitempty"`
	ShippingCity    string `json:"shippingCity,omitempty" force:"Physical_City__c,omitempty"`
	ShippingState   string `json:"shippingState,omitempty" force:"Physical_State_Province__c,omitempty"`
	ShippingZipCode string `json:"shippingZipCode,omitempty" force:"Physical_Zip_Postal_Code__c,omitempty"`
	ShippingCountry string `json:"shippingCountry,omitempty" force:"Physical_Country__c,omitempty"`
}

func (a *AccountDTO) toEntity() (*entities.Account, error) {
	account, err := entities.NewAccount(a.Name)

	if err != nil {
		return account,
			fmt.Errorf("Error converting to Account Entity: %v", err.Error())
	}

	account.Payer = a.Payer
	account.Industry = a.Industry

	if a.SiteID != "" {
		siteID, err := strconv.Atoi(a.SiteID)
		if err != nil {
			return account,
				fmt.Errorf("Error converting site ID: %v", err.Error())
		}

		err = account.SetSiteID(siteID)
		if err != nil {
			return account,
				fmt.Errorf("Error converting site ID: %v", err.Error())
		}
	}

	if a.BusinessUnit != "" {
		err = account.SetBusinessUnit(entities.BusinessUnit(a.BusinessUnit))
		if err != nil {
			return account,
				fmt.Errorf("Error converting BusinessUnit: %v", err.Error())
		}
	}

	if a.PrimaryStreet != "" || a.PrimaryCity != "" || a.PrimaryState != "" ||
		a.PrimaryZipCode != "" || a.PrimaryCountry != "" {
		account.PrimaryAddress = &entities.Address{}
		account.PrimaryAddress.Street = a.PrimaryStreet
		account.PrimaryAddress.City = a.PrimaryCity
		account.PrimaryAddress.State = a.PrimaryState
		account.PrimaryAddress.ZipCode = a.PrimaryZipCode
		account.PrimaryAddress.Country = a.PrimaryCountry
	}

	if a.BillingStreet != "" || a.BillingCity != "" || a.BillingState != "" ||
		a.BillingZipCode != "" || a.BillingCountry != "" {
		account.BillingAddress = &entities.Address{}
		account.BillingAddress.Street = a.BillingStreet
		account.BillingAddress.City = a.BillingCity
		account.BillingAddress.State = a.BillingState
		account.BillingAddress.ZipCode = a.BillingZipCode
		account.BillingAddress.Country = a.BillingCountry
	}

	if a.ShippingStreet != "" || a.ShippingCity != "" || a.ShippingState != "" ||
		a.ShippingZipCode != "" || a.ShippingCountry != "" {
		account.ShippingAddress = &entities.Address{}
		account.ShippingAddress.Street = a.ShippingStreet
		account.ShippingAddress.City = a.ShippingCity
		account.ShippingAddress.State = a.ShippingState
		account.ShippingAddress.ZipCode = a.ShippingZipCode
		account.ShippingAddress.Country = a.ShippingCountry
	}

	return account, err
}

// ConvertAccountEntityToAccountDTO converts an entity account to a data tranfer
// object. Used when creating or updating objects in the data store
func ConvertAccountEntityToAccountDTO(account *entities.Account) *AccountDTO {
	dto := &AccountDTO{
		Name:         account.Name(),
		SiteID:       strconv.Itoa(account.SiteID()),
		BusinessUnit: string(account.BusinessUnit()),
		Industry:     account.Industry,
		Payer:        account.Payer,
	}

	if account.PrimaryAddress != nil {
		dto.PrimaryStreet = account.PrimaryAddress.Street
		dto.PrimaryCity = account.PrimaryAddress.City
		dto.PrimaryState = account.PrimaryAddress.State
		dto.PrimaryZipCode = account.PrimaryAddress.ZipCode
		dto.PrimaryCountry = account.PrimaryAddress.Country
	}

	if account.BillingAddress != nil {
		dto.BillingStreet = account.BillingAddress.Street
		dto.BillingCity = account.BillingAddress.City
		dto.BillingState = account.BillingAddress.State
		dto.BillingZipCode = account.BillingAddress.ZipCode
		dto.BillingCountry = account.BillingAddress.Country
	}

	if account.ShippingAddress != nil {
		dto.ShippingStreet = account.ShippingAddress.Street
		dto.ShippingCity = account.ShippingAddress.City
		dto.ShippingState = account.ShippingAddress.State
		dto.ShippingZipCode = account.ShippingAddress.ZipCode
		dto.ShippingCountry = account.ShippingAddress.Country
	}

	return dto
}

// AccountService provides interaction with Account data
type AccountService struct {
	AccountRepo AccountRepository
}

// GetAccount returns an account by ID
func (as *AccountService) GetAccount(id string) (*AccountDTO, error) {
	a, err := as.AccountRepo.GetAccount(id)
	return a, err
}

// CreateAccount creates a new account
func (as *AccountService) CreateAccount(a AccountDTO) (id string, siteID int, err error) {
	account, err := a.toEntity()

	if err != nil {
		return "", 0, err
	}

	id, siteID, err = as.AccountRepo.CreateAccount(account)
	return id, siteID, err
}

//QueryAccounts returns a slice of the accunts returned by the query.
func (as *AccountService) QueryAccounts(query string) ([]*AccountDTO, error) {
	accounts, err := as.AccountRepo.QueryAccounts(query)

	return accounts, err
}

// UpdateAccount updatesn account
func (as *AccountService) UpdateAccount(a AccountDTO) error {
	account, err := a.toEntity()

	if err != nil {
		return err
	}

	err = as.AccountRepo.UpdateAccount(account)
	return err
}
