package entities

import (
	"testing"

	. "github.com/blackbaudIT/webcore/Godeps/_workspace/src/github.com/smartystreets/goconvey/convey"
)

var contactEntity = &Contact{
	name:            "Erik Tate",
	email:           "erik.tate@blackbaud.com",
	Phone:           "(843)654-2566",
	Fax:             "(843)654-2566",
	Title:           "Application Developer II",
	account:         &Account{name: "Test Account"},
	status:          "Active",
	Currency:        "USD - U.S. Dollar",
	bbAuthID:        "123456-1234-1234-1234-12345678",
	bbAuthEmail:     "erik.tate@blackbaud.com",
	bbAuthFirstName: "Erik",
	bbAuthLastName:  "Tate",
}

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

func TestContactGetters(t *testing.T) {
	Convey("Given a contact entity", t, func() {
		contact := contactEntity
		Convey("When requesting their name", func() {
			name := contact.Name()
			Convey("Then a name should be returned", func() {
				So(name, ShouldNotBeEmpty)
			})
		})
		Convey("When requesting their email", func() {
			email := contact.Email()
			Convey("Then an email should be returned", func() {
				So(email, ShouldNotBeEmpty)
			})
		})
		Convey("When requesting their account", func() {
			account := contact.Account()
			Convey("Then an account should be returned.", func() {
				So(account, ShouldNotBeNil)
			})
		})
		Convey("When requesting their status", func() {
			status := contact.Status()
			Convey("Then a status should be returned.", func() {
				So(status, ShouldNotBeEmpty)
			})
		})
		Convey("When requesting their bbAuthID", func() {
			authID := contact.BBAuthID()
			Convey("Then a BBAuthID should be returned", func() {
				So(authID, ShouldNotBeEmpty)
			})
		})
		Convey("When requesting their BBAuth Email", func() {
			authEmail := contact.BBAuthEmail()
			Convey("Then an email should be returned", func() {
				So(authEmail, ShouldNotBeEmpty)
			})
		})
		Convey("When requesting their BBAuth First Name", func() {
			firstName := contact.BBAuthFirstName()
			Convey("Then a first name should be returned", func() {
				So(firstName, ShouldNotBeEmpty)
			})
		})
		Convey("When requesting their BBAuth last Name", func() {
			lastName := contact.BBAuthLastName()
			Convey("Then a last name should be returned", func() {
				So(lastName, ShouldNotBeEmpty)
			})
		})
		Convey("When attempting to set their name", func() {
			name := "Test User"
			contact.SetName(name)
			Convey("Then the contact's name should be changed", func() {
				So(contact.Name(), ShouldEqual, name)
			})
		})
		Convey("When attempting to set their email", func() {
			email := "test@test.com"
			contact.SetEmail(email)
			Convey("Then the contact's email should be changed", func() {
				So(contact.Email(), ShouldEqual, email)
			})
		})
		Convey("When attempting to set their account", func() {
			account, _ := NewAccount("Super Test Account")
			contact.SetAccount(account)
			Convey("Then the contact's account should be changed", func() {
				So(contact.Account(), ShouldEqual, account)
			})
		})
		Convey("When attempting to set their status", func() {
			status := "Inactive"
			contact.SetStatus(status)
			Convey("Then the contact's status should be changed", func() {
				So(contact.Status(), ShouldEqual, status)
			})
		})
		Convey("When attempting to set their BBAuthID", func() {
			authID := "987654-9874-9874-9874-98765432"
			contact.SetBBAuthID(authID)
			Convey("Then the contact's BBAuthID should be changed", func() {
				So(contact.BBAuthID(), ShouldEqual, authID)
			})
		})
		Convey("When attempting to set their BBAuth Email", func() {
			email := "test@test.com"
			contact.SetBBAuthEmail(email)
			Convey("Then the contact's BBAuth Email should be changed", func() {
				So(contact.BBAuthEmail(), ShouldEqual, email)
			})
		})
		Convey("When attempting to set their BBAuth First Name", func() {
			firstName := "Test"
			contact.SetBBAuthFirstName(firstName)
			Convey("Then the contact's BBAuth First Name should be changed", func() {
				So(contact.BBAuthFirstName(), ShouldEqual, firstName)
			})
		})
		Convey("When attempting to set their BBAuth Last Name", func() {
			lastName := "User"
			contact.SetBBAuthLastName(lastName)
			Convey("Then the contact's BBAuth Last Name should be changed", func() {
				So(contact.BBAuthLastName(), ShouldEqual, lastName)
			})
		})
	})
}
