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
			externalId := contact.ExternalIdApiName()
			Convey("Then 'Username' should be returned", func() {
				So(externalId, ShouldEqual, "eBus_Contact_ID__c")
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
