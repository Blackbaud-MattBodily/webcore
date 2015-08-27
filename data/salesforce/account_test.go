package salesforce

import (
	"errors"
	"fmt"
	"testing"

	. "github.com/blackbaudIT/webcore/data/salesforce/Godeps/_workspace/src/github.com/smartystreets/goconvey/convey"
	"github.com/blackbaudIT/webcore/entities"
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
	sobject, ok := obj.(*SFDCAccount)
	if !ok {
		err = fmt.Errorf("unable to convert data to SFDCAccount. Unexpected type: %T", obj)
		return err
	}

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

func TestAccountApiName(t *testing.T) {
	Convey("Given an SFDCAccount object", t, func() {
		account := SFDCAccount{}
		Convey("When the API Name is requested", func() {
			apiName := account.ApiName()
			Convey("Then 'Account' should be returned", func() {
				So(apiName, ShouldEqual, "Account")
			})
		})
	})
}

func TestAccountExternalIdApiName(t *testing.T) {
	Convey("Given an SFDCAccount object", t, func() {
		account := SFDCAccount{}
		Convey("When the External Id API Name is requested", func() {
			apiName := account.ExternalIdApiName()
			Convey("Then 'Clarify_Site_ID__c' should be returned", func() {
				So(apiName, ShouldEqual, "Clarify_Site_ID__c")
			})
		})
	})
}

func TestGetAccount(t *testing.T) {
	Convey("Given a valid SFDC Id", t, func() {
		id := "001d000001TweFmAAJ"
		Convey("When requesting an account", func() {
			account, err := api.GetAccount(id)
			Convey("Then an AccountDTO is returned", func() {
				So(err, ShouldBeNil)
				So(account.SalesForceID, ShouldEqual, id)
			})
		})
	})
	Convey("Given a valid Clarify Site ID", t, func() {
		id := "5740"
		Convey("When requesting an account", func() {
			account, err := api.GetAccount(id)
			Convey("Then an AccountDTO is returned", func() {
				So(err, ShouldBeNil)
				So(account.SiteID, ShouldEqual, id)
			})
		})
	})
	Convey("Given an invalid ID", t, func() {
		cases := map[string]string{
			"(an empty string)":                              "",
			"(a zero value int)":                             "0",
			"(an int that doesn't exist in SFDC)":            "9999999",
			"(a float string)":                               "0.0",
			"(a non-SFDC string - length less than 15)":      "aaaa",
			"(a non-SFDC string - length between 15 and 18)": "aaaaaaaaaaaaaa",
			"(a non-SFDC string - length greater than 18)":   "aaaaaaaaaaaaaaaaaaaa",
		}

		for description, test := range cases {
			Convey(fmt.Sprintf("When requesting an account with an ID that is %s", description), func() {
				_, err := api.GetAccount(test)
				Convey("Then an error is returned", func() {
					So(err, ShouldNotBeNil)
				})
			})
		}
	})
}

func TestCreateAccount(t *testing.T) {
	Convey("Given an SFDCAccount object", t, func() {
		account, _ := entities.NewAccount("Test Org Name")
		Convey("When creating an account", func() {
			id, siteID, err := api.CreateAccount(account)
			Convey("Then a successful response should be returned", func() {
				So(err, ShouldBeNil)
				So(id, ShouldEqual, "001d000001TweFmAAJ")
				So(siteID, ShouldEqual, 5740)
			})
		})
		Convey("When an error occurs while executing SFDC command", func() {
			getCommandError = func() error { return errors.New("fake error") }
			_, _, err := api.CreateAccount(account)
			Convey("Then an error should be returned to the client", func() {
				So(err, ShouldNotBeNil)
			})
		})
		Convey("When an error occurs while querying new account from SFDC", func() {
			getQueryError = func() error { return errors.New("fake error") }
			_, _, err := api.CreateAccount(account)
			Convey("Then an error should be returned to the client", func() {
				So(err, ShouldNotBeNil)
			})
		})
		Convey("When an invalid SiteID is queried from SFDC", func() {
			getSFDCResposne = func() SFDCResponse {
				return SFDCResponse{
					ID:           "001d000001TweFmZZZ",
					ErrorMessage: "",
					Success:      true,
				}
			}
			_, _, err := api.CreateAccount(account)
			Convey("Then an error should be returned to the client", func() {
				So(err, ShouldNotBeNil)
			})
		})
		Convey("When an error occurs while persisting to SFDC", func() {
			getSFDCResposne = func() SFDCResponse {
				return SFDCResponse{
					ID:           "001d000001TweFmAAJ",
					ErrorMessage: "fake error",
					Success:      false,
				}
			}
			_, _, err := api.CreateAccount(account)
			Convey("Then an error should be returned to the client", func() {
				So(err, ShouldNotBeNil)
			})
		})
		Reset(func() {
			getCommandError = func() error { return nil }
			getQueryError = func() error { return nil }
			getSFDCResposne = func() SFDCResponse {
				return SFDCResponse{
					ID:           "001d000001TweFmAAJ",
					ErrorMessage: "",
					Success:      true,
				}
			}
		})
	})
}

func TestUpdateAccount(t *testing.T) {
	Convey("Given an SFDCAccount object", t, func() {
		account, _ := entities.NewAccount("Test Org Name")
		Convey("When updating an account without a SiteID", func() {
			err := api.UpdateAccount(account)
			Convey("Then an error should be returned", func() {
				So(err, ShouldNotBeNil)
			})
		})
		Convey("When updating an account with a SiteID", func() {
			account.SetSiteID(5740)
			err := api.UpdateAccount(account)
			Convey("Then the update should succeed", func() {
				So(err, ShouldBeNil)
			})
		})
		Convey("When an error occurs while executing SFDC command", func() {
			getCommandError = func() error { return errors.New("fake error") }
			account.SetSiteID(5740)
			err := api.UpdateAccount(account)
			Convey("Then an error should be returned to the client", func() {
				So(err, ShouldNotBeNil)
			})
		})
		Reset(func() {
			getCommandError = func() error { return nil }
			getQueryError = func() error { return nil }
		})
	})
}
