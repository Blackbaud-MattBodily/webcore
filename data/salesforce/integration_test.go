package salesforce

import (
	"testing"

	"github.com/blackbaudIT/webcore/services"
	. "github.com/smartystreets/goconvey/convey"
	"github.com/spf13/viper"
)

var qasAPI = NewAPI()

func TestGetSFDCObject(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping test in short mode.")
	}

	Convey("Given a valid ID", t, func() {
		id := "001d000001TweFmAAJ"
		Convey("When requesting an account", func() {
			account, err := qasAPI.GetAccount(id)
			Convey("Then an AccountDTO is returned", func() {
				So(err, ShouldBeNil)
				So(account.SalesForceID, ShouldEqual, id)
			})
		})
	})
	Convey("Given an invalid SObject", t, func() {
		id := "001d000001TweFmAAJ"
		obj := ""
		Convey("When requesting an account", func() {
			err := qasAPI.client.GetSFDCObject(id, obj)
			Convey("Then an error is returned", func() {
				So(err, ShouldNotBeNil)
			})
		})
	})
}

func TestGetSFDCObjectByExternalID(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping test in short mode.")
	}

	Convey("Given a valid ID", t, func() {
		id := "5740"
		Convey("When requesting an account", func() {
			account, err := qasAPI.GetAccount(id)
			Convey("Then an AccountDTO is returned", func() {
				So(err, ShouldBeNil)
				So(account.SiteID, ShouldEqual, id)
			})
		})
	})
	Convey("Given an invalid SObject", t, func() {
		id := "5740"
		obj := ""
		Convey("When requesting an account", func() {
			err := qasAPI.client.GetSFDCObjectByExternalID(id, obj)
			Convey("Then an error is returned", func() {
				So(err, ShouldNotBeNil)
			})
		})
	})
}

func TestInsertSFDCObjectByExternalID(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping test in short mode.")
	}

	Convey("Given the required account fields", t, func() {
		account := services.AccountDTO{
			Name:            "Integration Testing Account",
			BusinessUnit:    "GMBU",
			Industry:        "Cause & Cure",
			ShippingStreet:  "123 Random St",
			ShippingCity:    "Charleston",
			ShippingState:   "SC",
			ShippingZipCode: "29492",
			ShippingCountry: "USA",
		}
		obj := SFDCAccount{account}
		Convey("When inserting an account", func() {
			resp, err := qasAPI.client.InsertSFDCObject(obj)
			Convey("Then no error should occur", func() {
				So(err, ShouldBeNil)
				So(resp.Success, ShouldBeTrue)
				So(resp.ErrorMessage, ShouldBeEmpty)
				So(resp.ID, ShouldContainSubstring, "001") // SFDC Account Id prefix
				Reset(func() {
					// delete the account we just created
					id := resp.ID
					if id != "" {
						fc := forceClient{getForceAPIClient()}
						fc.DeleteSObject(id, obj)
					}
				})
			})
		})
	})
	Convey("Given an invalid SObject", t, func() {
		obj := ""
		Convey("When inserting an account", func() {
			_, err := qasAPI.client.InsertSFDCObject(obj)
			Convey("Then an error is returned", func() {
				So(err, ShouldNotBeNil)
			})
		})
	})
}

func TestUpsertSFDCObjectByExternalID(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping test in short mode.")
	}

	Convey("Given a valid account", t, func() {
		id := "93275"
		account := services.AccountDTO{Name: "Integration Testing Account", ShippingStreet: "123 Main St"}
		obj := SFDCAccount{account}
		Convey("When updating an account", func() {
			err := qasAPI.client.UpsertSFDCObjectByExternalID(id, obj)
			Convey("Then no error should occur", func() {
				So(err, ShouldBeNil)
			})
		})
	})
	Convey("Given an invalid SObject", t, func() {
		id := "93275"
		obj := ""
		Convey("When updating an account", func() {
			err := qasAPI.client.UpsertSFDCObjectByExternalID(id, obj)
			Convey("Then an error is returned", func() {
				So(err, ShouldNotBeNil)
			})
		})
	})
}

func TestInvalidConfig(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping test in short mode.")
	}

	Convey("Given invalid configuration", t, func() {
		viper.Set("sfdcUserName", "")
		viper.Set("sfdcClientId", "")
		viper.Set("sfdcClientSecret", "")
		Convey("When SFDC is accessed", func() {
			f := func() { getForceAPIClient() }
			Convey("Then the application should panic", func() {
				So(f, ShouldPanic)
			})
		})
	})
}