package entities

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestNewContact(t *testing.T) {
	Convey("Given an empty name", t, func() {
		account, _ := NewAccount("test")
		Convey("When a contact creation is attempted", func() {
			_, err := NewContact("", account, USD)
			Convey("Then an error should occur", func() {
				So(err, ShouldNotBeNil)
			})
		})
	})
	Convey("Given a nil Account", t, func() {
		Convey("When a contact creation is attempted", func() {
			_, err := NewContact("Test Two", nil, USD)
			Convey("Then an error should occur", func() {
				So(err, ShouldNotBeNil)
			})
		})
	})
	Convey("Given an empty currency type", t, func() {
		account, _ := NewAccount("test")
		Convey("When a contact creation is attempted", func() {
			_, err := NewContact("Test Three", account, "")
			Convey("Then an error should occur", func() {
				So(err, ShouldNotBeNil)
			})
		})
	})
	Convey("Given a name, account, and currency type", t, func() {
		name := "Test Four"
		account, _ := NewAccount("test")
		currency := USD

		Convey("When a contact creation is attempted", func() {
			contact, err := NewContact(name, account, currency)
			Convey("Then the Contact should be created without error", func() {
				So(err, ShouldBeNil)
				So(contact, ShouldNotBeNil)
			})
			Convey("And the Contact Name, Account, and Currency should equal the given values", func() {
				So(contact.Name(), ShouldEqual, name)
				So(contact.Account(), ShouldEqual, account)
				So(contact.Currency, ShouldEqual, currency)
			})
		})
	})
}
