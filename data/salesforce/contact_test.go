package salesforce

import (
	"fmt"
	"testing"

	. "github.com/blackbaudIT/webcore/Godeps/_workspace/src/github.com/smartystreets/goconvey/convey"
)

func TestContactApiName(t *testing.T) {
	Convey("Given an SFDCContact object", t, func() {
		contact := SFDCContact{}
		Convey("When the API Name is requested", func() {
			apiName := contact.ApiName()
			Convey("Then 'Contact' should be returned", func() {
				So(apiName, ShouldEqual, "Contact")
			})
		})
	})
}

func TestContactExternalIdApiName(t *testing.T) {
	Convey("Given an SFDCContact object", t, func() {
		contact := SFDCContact{}
		Convey("When the External ID API name is requested", func() {
			externalID := contact.ExternalIdApiName()
			Convey("Then 'Username' should be returned", func() {
				So(externalID, ShouldEqual, "eBus_Contact_ID__c")
			})
		})
	})
}

func TestGetContact(t *testing.T) {
	Convey("Given a valid SFDC Id", t, func() {
		id := "003d0000027LKPQAA4"
		Convey("When requesting a contact", func() {
			contact, err := api.GetContact(id)
			Convey("Then a ContactDTO is returned", func() {
				So(err, ShouldBeNil)
				So(contact.SalesForceID, ShouldEqual, id)
			})
		})
	})
	Convey("Given an invalid ID", t, func() {
		cases := map[string]string{
			"(an empty string)":                              "",
			"(a non-SFDC string - length less than 15)":      "aaaa",
			"(a non-SFDC string - length between 15 and 18)": "aaaaaaaaaaaaaa",
			"(a non-SFDC string - length greater than 18)":   "aaaaaaaaaaaaaaaaaaaa",
		}

		for description, test := range cases {
			Convey(fmt.Sprintf("When requesting a contact with an ID that is %s", description), func() {
				_, err := api.GetContact(test)
				Convey("Then an error is returned", func() {
					So(err, ShouldNotBeNil)
				})
			})
		}
	})
}

func TestQueryContacts(t *testing.T) {
	Convey("Given a valid query", t, func() {
		id := "32FBC72D-C0FE-4B50-B0F4-EDCEFD7B4DEF"
		Convey("When requesting a list of contacts", func() {
			contacts, err := api.QueryContacts("select Id, Name, Account, CurrencyISOCode, BBAuthID__c from Contact " +
				"where BBAuthID__c = " + id)
			Convey("Then a list of Contacts should be returned", func() {
				So(len(contacts), ShouldBeGreaterThan, 0)
				So(err, ShouldBeNil)
			})
		})
	})
	Convey("Given an invalid query", t, func() {
		id := ""
		Convey("When requesting a list of contacts", func() {
			contacts, err := api.QueryContacts("delect Id, Name, Account, CurrencyISOCode, BBAuthID__c from Contact " +
				"where BBAuthID__c = " + id)
			Convey("Then an empty list and an error should be returned", func() {
				So(len(contacts), ShouldBeZeroValue)
				So(err, ShouldNotBeNil)
			})
		})
	})
}

func TestGetByAuthID(t *testing.T) {
	Convey("Given a valid Auth ID", t, func() {
		id := "32FBC72D-C0FE-4B50-B0F4-EDCEFD7B4DEF"
		Convey("When requesting a contact query string", func() {
			query, err := api.GetByAuthID(id)
			Convey("Then a query string should be returned", func() {
				So(query, ShouldNotBeEmpty)
				So(err, ShouldBeNil)
			})
		})
	})
	Convey("Given an invalid AuthID", t, func() {
		id := "12345"
		Convey("When requesting a contact query string", func() {
			query, err := api.GetByAuthID(id)
			Convey("Then an error should be returned", func() {
				So(query, ShouldBeEmpty)
				So(err, ShouldNotBeNil)
			})
		})
	})
}

func TestGetByEmail(t *testing.T) {
	Convey("Given a valid email", t, func() {
		email := "erik.tate@blackbaud.com"
		Convey("When requesting a contact query string", func() {
			query, err := api.GetByEmail(email)
			Convey("Then a query string should be returned", func() {
				So(query, ShouldNotBeEmpty)
				So(err, ShouldBeNil)
			})
		})
	})
	Convey("Given an invalid email", t, func() {
		email := "erik.tateblackbaud.com"
		Convey("When requesting a contact query string", func() {
			query, err := api.GetByEmail(email)
			Convey("Then an error should be returned", func() {
				So(query, ShouldBeEmpty)
				So(err, ShouldNotBeNil)
			})
		})
	})
}
