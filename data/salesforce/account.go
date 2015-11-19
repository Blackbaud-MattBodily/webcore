package salesforce

import (
	"errors"
	"fmt"
	"strconv"

	"github.com/blackbaudIT/webcore/entities"
	"github.com/blackbaudIT/webcore/services"
)

// SFDCAccount wraps the Account Data Transfer Object so that SFDC fields can
// be mapped onto it
type SFDCAccount struct {
	services.AccountDTO
}

//SFDCAccountQueryResponse wraps the base SFDCQueryResponse and attaches a slice
//of SFDCAccount pointers that will be written into
type SFDCAccountQueryResponse struct {
	SFDCQueryResponse

	Records []*services.AccountDTO `json:"Records" force:"records"`
}

// ApiName is the SFDC ApiName of the Account object
func (s SFDCAccount) ApiName() string {
	return "Account"
}

// ExternalIdApiName returns the External ID in SFDC for the Account object
func (s SFDCAccount) ExternalIdApiName() string {
	return "Clarify_Site_ID__c"
}

// GetAccount returns a SalesForce account for the ID specified
func (a API) GetAccount(id string) (*services.AccountDTO, error) {
	account := &SFDCAccount{}
	count, _ := a.GetContactCount(id)
	account.ContactCount = count

	accountLookupFunc, err := a.getForceAPILookupFunction(id)
	if err != nil {
		return &account.AccountDTO, fmt.Errorf("Error validating id parameter: %s \n", err)
	}

	err = accountLookupFunc(account)
	if err != nil {
		return &account.AccountDTO, fmt.Errorf("Error querying SFDC: %s \n", err)
	}
	return &account.AccountDTO, nil
}

func (a API) getForceAPILookupFunction(id string) (func(*SFDCAccount) error, error) {

	i, err := strconv.Atoi(id)

	// if we couldn't convert to int, assume this is an SFDC Id
	if err != nil {
		if id == "" {
			return nil, errors.New("id cannot be an emtpy string")
		}
		if len(id) == 15 || len(id) == 18 {
			return func(account *SFDCAccount) error {
				err := a.client.GetSFDCObject(id, account)
				return err
			}, nil
		}

		return nil, errors.New("SFDC Id's must be an alphanumeric string with a" +
			"length of 15 or 18")
	}

	// if converted to an int, then assume this is the Clarify Site ID
	if i > 0 {
		return func(account *SFDCAccount) error {
			err := a.client.GetSFDCObjectByExternalID(id, account)
			return err
		}, nil
	}

	return nil, errors.New("id cannot be 0")
}

//QueryAccounts returns a slice of AccountDTO pointers that represent the
//results of the given query.
func (a API) QueryAccounts(query string) ([]*services.AccountDTO, error) {
	//TODO: Need to retrieve the contact count for these accounts before returing slice.
	queryResponse := &SFDCAccountQueryResponse{}

	err := a.client.QuerySFDCObject(query, queryResponse)

	return queryResponse.Records, err
}

// CreateAccount creates a new SFDC Account and returns the Clarify Site ID
func (a API) CreateAccount(account *entities.Account) (string, int, error) {
	dto := services.ConvertAccountEntityToAccountDTO(account)

	sfdcAccount := SFDCAccount{AccountDTO: *dto}
	resp, err := a.client.InsertSFDCObject(sfdcAccount)
	if err != nil {
		return "", 0, fmt.Errorf("Error creating account in SFDC: %s", err)
	}
	if !resp.Success {
		return "", 0,
			fmt.Errorf("Error creating account in SFDC: %s", resp.ErrorMessage)
	}

	newAccount := &SFDCAccount{}
	err = a.client.GetSFDCObject(resp.ID, newAccount)
	if err != nil {
		return "", 0, fmt.Errorf("Error getting newly created account: %s", err)
	}

	var siteID int
	if newAccount.SiteID != "" {
		siteID, err = strconv.Atoi(newAccount.SiteID)
		if err != nil {
			return "", 0,
				fmt.Errorf("Error getting SiteID for newly created account: %s", err)
		}
	}

	return resp.ID, siteID, nil
}

// UpdateAccount updates an SFDC Account
func (a API) UpdateAccount(account *entities.Account) error {
	if account.SiteID() <= 0 {
		return fmt.Errorf("A valid SiteID is required to update an account (SiteID: %v)",
			account.SiteID())
	}

	dto := services.ConvertAccountEntityToAccountDTO(account)

	// this is necessary because we are using the SiteID as the external ID for
	// the upsert, so it can't be included in the field list. if we don't clear
	// it out, SFDC will error with: "The Clarify_Site_ID__c field should not be
	// specified in the sobject data"
	siteID := dto.SiteID
	dto.SiteID = ""

	sfdcAccount := SFDCAccount{AccountDTO: *dto}
	err := a.client.UpsertSFDCObjectByExternalID(siteID, sfdcAccount)
	if err != nil {
		return fmt.Errorf("Error updating account in SFDC: %s", err)
	}

	return nil
}

//GetContactCount returns the number of salesforce contacts currently associated with an account.
func (a API) GetContactCount(accountID string) (int, error) {
	queryResponse := &SFDCContactQueryResponse{}
	query := "SELECT count() FROM Contact WHERE AccountId = '" + accountID + "'"

	err := a.client.QuerySFDCObject(query, queryResponse)

	return int(queryResponse.TotalSize), err
}
