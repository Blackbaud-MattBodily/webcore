package services

import (
	"testing"

	. "github.com/blackbaudIT/webcore/Godeps/_workspace/src/github.com/smartystreets/goconvey/convey"
	"github.com/blackbaudIT/webcore/entities"
)

var contactDTO = ContactDTO{
	Name:              "Erik Tate",
	SalesForceID:      "003d0000026MOlUAAW",
	Email:             "erik.tate@blackbaud.com",
	Phone:             "(843)654-2566",
	Fax:               "(843)654-2566",
	Title:             "Application Developer II",
	Account:           &accountDTO,
	SFDCContactStatus: "Active",
	Currency:          "USD - U.S. Dollar",
	BBAuthID:          "32FBC72D-C0FE-4B50-B0F4-EDCEFD7B4DEF",
	BBAuthFirstName:   "Erik",
	BBAuthLastName:    "Tate",
}

var contactService = ContactService{ContactRepo: mockContactRepository{}}

type mockContactRepository struct {
}

func (m mockContactRepository) GetContact(id string) (*ContactDTO, error) {
	return &contactDTO, nil
}

func (m mockContactRepository) QueryContacts(query string) ([]*ContactDTO, error) {
	return nil, nil
}

func (m mockContactRepository) UpdateContact(contact *entities.Contact) error {
	return nil
}

func TestContactDTOToEntity(t *testing.T) {
	Convey("Given a Contact Data Transfer object with an empty name", t, func() {
		contactDTOCopy := contactDTO
		contactDTOCopy.Name = ""
		Convey("When attempting to convert to a Contact entity", func() {
			_, err := contactDTOCopy.toEntity()
			Convey("An error should occur", func() {
				So(err, ShouldNotBeNil)
			})
		})
	})
	Convey("Given a Contact Data Transfer object with an empty account", t, func() {
		contactDTOCopy := contactDTO
		contactDTOCopy.Account = nil
		Convey("when attempting to convert to a Contact entity", func() {
			_, err := contactDTOCopy.toEntity()
			Convey("An error should occur", func() {
				So(err, ShouldNotBeNil)
			})
		})
	})
	Convey("Given a Contact Data Transfer object with no currency", t, func() {
		contactDTOCopy := contactDTO
		contactDTOCopy.Currency = ""
		Convey("When attempting to convert to a Contact entity", func() {
			_, err := contactDTOCopy.toEntity()
			Convey("An error should occur", func() {
				So(err, ShouldNotBeNil)
			})
		})
	})
	Convey("Given a valid Contact Data Transfer object", t, func() {
		contactDTOCopy := contactDTO
		Convey("When attempting to convert to a Contact entity", func() {
			contact, err := contactDTOCopy.toEntity()
			Convey("No error should occur", func() {
				So(err, ShouldBeNil)
				So(contact, ShouldNotBeNil)
			})
		})
	})
}

func TestConvertContactEntityToContactDTO(t *testing.T) {
	Convey("Given a valid contact entity", t, func() {
		contact, _ := contactDTO.toEntity()
		Convey("When attempting to convert to a ContactDTO", func() {
			convertedContactDTO := ConvertContactEntityToContactDTO(contact)
			Convey("Then a ContactDTO should be returned", func() {
				So(convertedContactDTO, ShouldNotBeNil)
			})
		})
	})
}

func TestNewContactService(t *testing.T) {
	Convey("Given a valid ContactRepo", func() {
		contactRepo := mockContactRepository{}
		Convey("When a new contact service is requested", func() {
			cs := NewContactService(contactRepo)
		})
	})
}
