package entities

import (
	"testing"

	. "github.com/blackbaudIT/webcore/Godeps/_workspace/src/github.com/smartystreets/goconvey/convey"
)

var contactEntity = &Contact{
	id:              "123456test",
	Name:            &Name{"Mr.", "Erik", "Tate"},
	email:           "erik.tate@blackbaud.com",
	Phone:           "(843)654-2566",
	Fax:             "(843)654-2566",
	Title:           "Application Developer II",
	account:         &Account{name: "Test Account"},
	defaultAccount:  "123",
	status:          "Active",
	Currency:        "USD - U.S. Dollar",
	bbAuthID:        "123456-1234-1234-1234-12345678",
	bbAuthEmail:     "erik.tate@blackbaud.com",
	bbAuthFirstName: "Erik",
	bbAuthLastName:  "Tate",
}

func TestNewContact(t *testing.T) {
	Convey("Given an empty last name", t, func() {
		name := &Name{"Mr.", "Erik", ""}
		Convey("When a contact creation is attempted", func() {
			account, _ := NewAccount("test")
			_, err := NewContact(name, account, USD)
			Convey("Then an error should occur", func() {
				So(err, ShouldNotBeNil)
			})
		})
	})
	Convey("Given a nil Account", t, func() {
		Convey("When a contact creation is attempted", func() {
			_, err := NewContact(&Name{"", "Erik", "Tate"}, nil, USD)
			Convey("Then an error should occur", func() {
				So(err, ShouldNotBeNil)
			})
		})
	})
	Convey("Given an empty currency type", t, func() {
		account, _ := NewAccount("test")
		Convey("When a contact creation is attempted", func() {
			_, err := NewContact(&Name{"", "Erik", "Tate"}, account, "")
			Convey("Then an error should occur", func() {
				So(err, ShouldNotBeNil)
			})
		})
	})
	Convey("Given a name, account, and currency type", t, func() {
		name := &Name{"Mr.", "Erik", "Tate"}
		account, _ := NewAccount("test")
		currency := USD

		Convey("When a contact creation is attempted", func() {
			contact, err := NewContact(name, account, currency)
			Convey("Then the Contact should be created without error", func() {
				So(err, ShouldBeNil)
				So(contact, ShouldNotBeNil)
			})
			Convey("And the Contact Name, Account, and Currency should equal the given values", func() {
				So(contact.Name.Salutation, ShouldEqual, name.Salutation)
				So(contact.Name.FirstName, ShouldEqual, name.FirstName)
				So(contact.Name.LastName(), ShouldEqual, name.LastName())
				So(contact.Account(), ShouldEqual, account)
				So(contact.Currency, ShouldEqual, currency)
			})
		})
	})
}

func TestContactGettersAndSetters(t *testing.T) {
	Convey("Given a contact entity", t, func() {
		contact := contactEntity
		Convey("When requesting their last name", func() {
			name := contact.Name.LastName()
			Convey("Then a last name should be returned", func() {
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
		Convey("When requesting their default account", func() {
			defaultAccount := contact.DefaultAccount()
			Convey("Then a default account ID should be returned.", func() {
				So(defaultAccount, ShouldNotBeEmpty)
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
		Convey("When attempting to set their ID", func() {
			id := "123456789test"
			contact.SetID(id)
			Convey("Then thec ontact's ID should be changed", func() {
				So(contact.ID(), ShouldEqual, id)
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
		Convey("When attempting to set their default account", func() {
			defaultAccount := "123456"
			contact.SetDefaultAccount(defaultAccount)
			Convey("Then the contact's default account should be changed", func() {
				So(contact.DefaultAccount(), ShouldEqual, defaultAccount)
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

func TestNameSetLastName(t *testing.T) {
	Convey("Given a valid name struct", t, func() {
		name, _ := BuildName("Mr.", "Erik", "Tate")
		Convey("and a valid last name", func() {
			lastName := "User"
			Convey("When attempting to set their last name", func() {
				name.SetLastName(lastName)
				Convey("Then the contact's last name should be changed", func() {
					So(name.LastName(), ShouldEqual, lastName)
				})
			})
		})
		Convey("and an invalid last name", func() {
			lastName := ""
			Convey("When attempting to set their last name", func() {
				err := name.SetLastName(lastName)
				Convey("Then an error should occur", func() {
					So(err, ShouldNotBeNil)
				})
			})
		})
	})
}

func TestBuildName(t *testing.T) {
	Convey("Given some name information with a blank last name", t, func() {
		salutation := "Mr."
		firstName := "Erik"
		lastName := ""
		Convey("When attempting to build a name", func() {
			_, err := BuildName(salutation, firstName, lastName)
			Convey("Then an error should occur", func() {
				So(err, ShouldNotBeNil)
			})
		})
	})
	Convey("Given some name information with a valid last name", t, func() {
		salutation := "Mr."
		firstName := "Erik"
		lastName := "Tate"
		Convey("When attempting to build a name", func() {
			_, err := BuildName(salutation, firstName, lastName)
			Convey("Then no error should occur", func() {
				So(err, ShouldBeNil)
			})
		})
	})
}
