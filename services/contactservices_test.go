package services

import (
	"errors"
	"testing"

	. "github.com/blackbaudIT/webcore/Godeps/_workspace/src/github.com/smartystreets/goconvey/convey"
)

var contactDTO = ContactDTO{
	Salutation:      "Mr.",
	FirstName:       "Erik",
	LastName:        "Tate",
	SalesForceID:    "003d0000026MOlUAAW",
	Email:           "erik.tate@blackbaud.com",
	Phone:           "(843)654-2566",
	Fax:             "(843)654-2566",
	Title:           "Application Developer II",
	Account:         &accountDTO,
	Status:          "Active",
	Currency:        "USD - U.S. Dollar",
	BBAuthID:        "32FBC72D-C0FE-4B50-B0F4-EDCEFD7B4DEF",
	BBAuthFirstName: "Erik",
	BBAuthLastName:  "Tate",
}

var contactService = ContactService{ContactRepo: mockContactRepository{}}

type mockContactRepository struct {
}

func (m mockContactRepository) GetContact(id string) (*ContactDTO, error) {
	if len(id) > 0 {
		return &contactDTO, nil
	}
	return nil, errors.New("An ID must be provided to get a contact")
}

func (m mockContactRepository) QueryContacts(query string) ([]*ContactDTO, error) {
	var contacts []*ContactDTO
	err := errors.New("Bad query")

	if query == "success!" {
		contacts = append(contacts, &contactDTO)
		err = nil
	}

	return contacts, err
}

func (m mockContactRepository) UpdateContact(contact *ContactDTO) error {
	return nil
}

func (m mockContactRepository) GetByAuthID(authID string) (string, error) {
	if len(authID) > 0 {
		return "success!", nil
	}

	return "", errors.New("Must provide a valid authID")
}

func (m mockContactRepository) GetByEmail(email string) (string, error) {
	if len(email) > 0 {
		return "success!", nil
	}

	return "", errors.New("Must provide a valid email address")
}

func (m mockContactRepository) GetByIDs(ids []string) (string, error) {
	if len(ids) > 0 {
		return "success!", nil
	}

	return "", errors.New("Must provide a slice of IDs")
}

func TestContactDTOToEntity(t *testing.T) {
	Convey("Given a Contact Data Transfer object with an empty last name", t, func() {
		contactDTOCopy := contactDTO
		contactDTOCopy.LastName = ""
		Convey("When attempting to convert to a Contact entity", func() {
			_, err := contactDTOCopy.ToEntity()
			Convey("An error should occur", func() {
				So(err, ShouldNotBeNil)
			})
		})
	})
	Convey("Given a Contact Data Transfer object with an empty account", t, func() {
		contactDTOCopy := contactDTO
		contactDTOCopy.Account = nil
		Convey("when attempting to convert to a Contact entity", func() {
			_, err := contactDTOCopy.ToEntity()
			Convey("An error should occur", func() {
				So(err, ShouldNotBeNil)
			})
		})
	})
	Convey("Given a Contact Data Transfer object with no currency", t, func() {
		contactDTOCopy := contactDTO
		contactDTOCopy.Currency = ""
		Convey("When attempting to convert to a Contact entity", func() {
			_, err := contactDTOCopy.ToEntity()
			Convey("An error should occur", func() {
				So(err, ShouldNotBeNil)
			})
		})
	})
	Convey("Given a valid Contact Data Transfer object", t, func() {
		contactDTOCopy := contactDTO
		Convey("When attempting to convert to a Contact entity", func() {
			contact, err := contactDTOCopy.ToEntity()
			Convey("No error should occur", func() {
				So(err, ShouldBeNil)
				So(contact, ShouldNotBeNil)
			})
		})
	})
}

func TestConvertContactEntityToContactDTO(t *testing.T) {
	Convey("Given a valid contact entity", t, func() {
		contact, _ := contactDTO.ToEntity()
		Convey("When attempting to convert to a ContactDTO", func() {
			convertedContactDTO := ConvertContactEntityToContactDTO(contact)
			Convey("Then a ContactDTO should be returned", func() {
				So(convertedContactDTO, ShouldNotBeNil)
			})
		})
	})
}

func TestNewContactService(t *testing.T) {
	Convey("Given a valid ContactRepo", t, func() {
		contactRepo := mockContactRepository{}
		Convey("When a new contact service is requested", func() {
			cs := NewContactService(contactRepo)
			Convey("A pointer to a contact is returned", func() {
				So(cs, ShouldNotBeNil)
			})
		})
	})
}

func TestGetContact(t *testing.T) {
	Convey("Given a Contact ID", t, func() {
		id := "003d0000026MOlUAAW"
		Convey("When a contact is requested", func() {
			cs := NewContactService(mockContactRepository{})
			contact, err := cs.GetContact(id)
			Convey("A Contact Data Transfer Object is returned", func() {
				So(contact, ShouldNotBeNil)
				So(err, ShouldBeNil)
			})
		})
	})
	Convey("Given an empty Contact ID", t, func() {
		id := ""
		Convey("When a contact is requested", func() {
			cs := NewContactService(mockContactRepository{})
			_, err := cs.GetContact(id)
			Convey("An error should occur", func() {
				So(err, ShouldNotBeNil)
			})
		})
	})
}

func TestGetContactsByEmail(t *testing.T) {
	Convey("Given an email", t, func() {
		email := "erik.tate@blackbaud.com"
		Convey("When a list of contacts are queried", func() {
			cs := NewContactService(mockContactRepository{})
			contacts, err := cs.GetContactsByEmail(email)
			Convey("A list of contacts should be returned", func() {
				So(contacts, ShouldNotBeEmpty)
				So(err, ShouldBeNil)
			})
		})
	})
	Convey("Given a blank email", t, func() {
		email := ""
		Convey("When a list of contacts are queried", func() {
			cs := NewContactService(mockContactRepository{})
			contacts, err := cs.GetContactsByEmail(email)
			Convey("An error should occur and no contacts should be returned", func() {
				So(contacts, ShouldBeEmpty)
				So(err, ShouldNotBeNil)
			})
		})
	})
}

func TestGetContactsByAuthID(t *testing.T) {
	Convey("Given an auth ID", t, func() {
		id := "32FBC72D-C0FE-4B50-B0F4-EDCEFD7B4DEF"
		Convey("When a list of contacts are queried", func() {
			cs := NewContactService(mockContactRepository{})
			contacts, err := cs.GetContactsByAuthID(id)
			Convey("A list of contacts should be returned", func() {
				So(contacts, ShouldNotBeEmpty)
				So(err, ShouldBeNil)
			})
		})
	})
	Convey("Given an empty auth ID", t, func() {
		id := ""
		Convey("When a list of contacts are queried", func() {
			cs := NewContactService(mockContactRepository{})
			contacts, err := cs.GetContactsByAuthID(id)
			Convey("An error should occur and no contacts should be returned", func() {
				So(contacts, ShouldBeEmpty)
				So(err, ShouldNotBeNil)
			})
		})
	})
}

func TestGetContactsByIDs(t *testing.T) {
	Convey("Given a list of IDs", t, func() {
		ids := []string{"1234", "5678"}
		Convey("when a list of contacts are queried", func() {
			cs := NewContactService(mockContactRepository{})
			contacts, err := cs.GetContactsByIDs(ids)
			Convey("A list of contacts should be returned", func() {
				So(contacts, ShouldNotBeEmpty)
				So(err, ShouldBeNil)
			})
		})
	})
	Convey("Given a blank list of IDs", t, func() {
		ids := []string{}
		Convey("When a list of contacts are qeried", func() {
			cs := NewContactService(mockContactRepository{})
			contacts, err := cs.GetContactsByIDs(ids)
			Convey("An error should occur and no contacts should be returned", func() {
				So(contacts, ShouldBeEmpty)
				So(err, ShouldNotBeNil)
			})
		})
	})
}
func TestQueryContactsByAuthID(t *testing.T) {
	Convey("Given a contact query string", t, func() {
		query := "success!"
		Convey("When a list of contacts are queried", func() {
			cs := NewContactService(mockContactRepository{})
			contacts, err := cs.QueryContacts(query)
			Convey("A list of contacts should be returned", func() {
				So(contacts, ShouldNotBeEmpty)
				So(err, ShouldBeNil)
			})
		})
	})
	Convey("Given a blank contact query string", t, func() {
		query := ""
		Convey("When a list of contacts are queried", func() {
			cs := NewContactService(mockContactRepository{})
			contacts, err := cs.QueryContacts(query)
			Convey("An error should occur and no contacts should be returned", func() {
				So(contacts, ShouldBeEmpty)
				So(err, ShouldNotBeNil)
			})
		})
	})
}

func TestUpdateContact(t *testing.T) {
	Convey("Given a contact DTO", t, func() {
		Convey("When an update is attempted", func() {
			cs := NewContactService(mockContactRepository{})
			err := cs.UpdateContact(&contactDTO)
			Convey("Then no error should occur", func() {
				So(err, ShouldBeNil)
			})
		})
	})
}
