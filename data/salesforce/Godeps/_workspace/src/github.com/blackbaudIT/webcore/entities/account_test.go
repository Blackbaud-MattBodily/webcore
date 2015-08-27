package entities

import (
	"testing"

	. "github.com/blackbaudIT/webcore/data/salesforce/Godeps/_workspace/src/github.com/smartystreets/goconvey/convey"
)

func TestNewAccount(t *testing.T) {
	Convey("Given an empty account name", t, func() {
		Convey("When an account creation is attempted", func() {
			_, err := NewAccount("")
			Convey("Then an error should occur", func() {
				So(err, ShouldNotBeNil)
			})
		})
	})
	Convey("Given an non-empty account name", t, func() {
		name := "Test Org Name"
		Convey("When an account creation is attempted", func() {
			account, err := NewAccount(name)
			Convey("Then the Account should be created without error", func() {
				So(err, ShouldBeNil)
				So(account, ShouldNotBeNil)
			})
			Convey("And the Account Name should equal the name provided", func() {
				So(account.Name(), ShouldEqual, name)
			})
		})
	})
}

func TestAccountSetName(t *testing.T) {
	Convey("Given an existing account", t, func() {
		name := "Test Org Name"
		account, err := NewAccount(name)
		Convey("When the account name is updated with an empty string", func() {
			err = account.SetName("")
			Convey("Then an error should occur", func() {
				So(err, ShouldNotBeNil)
			})
			Convey("And the account name should not be changed", func() {
				So(account.Name(), ShouldEqual, name)
			})
		})
		Convey("When the account name is updated with a non-empty string", func() {
			newName := "Changed Org Name"
			err = account.SetName(newName)
			Convey("Then the Account should be updated without error", func() {
				So(err, ShouldBeNil)
			})
			Convey("And the Account Name should equal the name provided", func() {
				So(account.Name(), ShouldEqual, newName)
			})
		})
	})
}

func TestAccountSetSiteID(t *testing.T) {
	Convey("Given an existing account", t, func() {
		name := "Test Org Name"
		account, err := NewAccount(name)
		Convey("When the site ID is updated with an invalid int", func() {
			err = account.SetSiteID(0)
			Convey("Then an error should occur", func() {
				So(err, ShouldNotBeNil)
			})
		})
		Convey("When the site ID is updated valid int", func() {
			siteID := 5740
			err = account.SetSiteID(siteID)
			Convey("Then the Account should be updated without error", func() {
				So(err, ShouldBeNil)
			})
			Convey("And the Account SiteID should equal the id provided", func() {
				So(account.SiteID(), ShouldEqual, siteID)
			})
		})
	})
}

func TestAccountSetBusinessUnit(t *testing.T) {
	Convey("Given an existing account", t, func() {
		name := "Test Org Name"
		account, err := NewAccount(name)
		Convey("When the business unit is updated with an invalid enum", func() {
			err = account.SetBusinessUnit(BusinessUnit("NOBU"))
			Convey("Then an error should occur", func() {
				So(err, ShouldNotBeNil)
			})
		})
		Convey("When the business unit is updated valid enum", func() {
			bu := GMBU
			err = account.SetBusinessUnit(bu)
			Convey("Then the Account should be updated without error", func() {
				So(err, ShouldBeNil)
			})
			Convey("And the Account BusinessUnit should equal the BU provided", func() {
				So(account.BusinessUnit(), ShouldEqual, bu)
			})
		})
	})
}
