package salesforce

import (
	"errors"
	"fmt"
	"strings"

	"github.com/blackbaudIT/webcore/services"
	"github.com/nimajalali/go-force/force"
)

var api = API{mockClient{}}
var getCommandError = func() error { return nil }
var getQueryError = func() error { return nil }
var getSFDCResposne = func() SFDCResponse {
	return SFDCResponse{
		ID:           "001d000001TweFmAAJ",
		ErrorMessage: "",
		Success:      true,
	}
}

type mockClient struct {
}

func (m mockClient) GetSFDCObject(id string, obj interface{}) (err error) {
	_, ok := obj.(force.SObject)
	if !ok {
		err = fmt.Errorf("unable to convert data to sObject. Unexpected type: %T", obj)
		return err
	}

	sfdcaccount, ok := obj.(*SFDCAccount)

	if ok {
		return getAccount(id, sfdcaccount)
	}

	sfdccontact, ok := obj.(*SFDCContact)

	if ok {
		return getContact(id, sfdccontact)
	}

	return fmt.Errorf("unused sObject given. Webcore does not work with: %T", obj)
}

func getAccount(id string, sobject *SFDCAccount) (err error) {
	sobject.SalesForceID = id

	// used to return a valid SiteID during create test
	if id == "001d000001TweFmAAJ" {
		sobject.SiteID = "5740"
	}

	// used to return an invalid SiteID during create test
	if id == "001d000001TweFmZZZ" {
		sobject.SiteID = "a5740"
	}

	return getQueryError()
}

func getContact(id string, sobject *SFDCContact) (err error) {
	sobject.SalesForceID = id

	return getQueryError()
}

func (m mockClient) GetSFDCObjectByExternalID(id string, obj interface{}) (err error) {
	sobject, ok := obj.(*SFDCAccount)
	if !ok {
		err = fmt.Errorf("unable to convert data to SFDCAccount. Unexpected type: %T", obj)
		return err
	}

	// mock a non-existing account
	if id == "9999999" {
		err = fmt.Errorf("unable to find account with ID: %s", id)
		return err
	}

	sobject.SiteID = id
	return getQueryError()
}

func (m mockClient) InsertSFDCObject(obj interface{}) (resposne SFDCResponse, err error) {
	return getSFDCResposne(), getCommandError()
}

func (m mockClient) UpsertSFDCObjectByExternalID(id string, obj interface{}) (err error) {
	return getCommandError()
}

func (m mockClient) QuerySFDCObject(query string, obj interface{}) (err error) {
	contact, ok := obj.(*SFDCContactQueryResponse)

	if ok {
		return queryContacts(query, contact)
	}

	account, ok := obj.(*SFDCAccountQueryResponse)

	if ok {
		return queryAccounts(query, account)
	}

	return errors.New("obj is not a valid SFDCQueryResponse")
}

func (m mockClient) UpdateSFDCObject(id string, obj interface{}) error {
	return getCommandError()
}

func queryContacts(query string, res *SFDCContactQueryResponse) error {
	res.TotalSize = 0
	if strings.Split(query, " ")[0] == "delect" {
		return errors.New("Invalid query passed to QuerySFDCObject")
	}

	contacts := make([]*services.ContactDTO, 1)
	account := &services.AccountDTO{Name: "Test Account"}
	contact := &services.ContactDTO{Name: "Test Contact", Account: account, Currency: "US"}
	contacts[0] = contact

	res.Records = contacts

	if query == "SELECT count() FROM Contact WHERE AccountId = '001d000001TwgVCAAZ'" {
		res.TotalSize = 1337
	}

	if query == "SELECT count() FROM Contact WHERE AccountId = '12345'" {
		return errors.New("AccountId doesn't exist")
	}
	return nil
}

func queryAccounts(query string, res *SFDCAccountQueryResponse) error {
	if strings.Split(query, " ")[0] == "delect" {
		return errors.New("Invalid query passed to QuerySFDCObject")
	}

	accounts := make([]*services.AccountDTO, 1)
	account := &services.AccountDTO{Name: "Test Account"}
	accounts[0] = account

	res.Records = accounts
	return nil
}
