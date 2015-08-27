package services

import (
	"strconv"
	"testing"

	"github.com/blackbaudIT/webcore/data/salesforce/Godeps/_workspace/src/github.com/blackbaudIT/webcore/entities"
	. "github.com/blackbaudIT/webcore/data/salesforce/Godeps/_workspace/src/github.com/smartystreets/goconvey/convey"
)

var accountDTO = AccountDTO{
	Name:            "Test Org Name",
	SalesForceID:    "001d000001TwuXwAAJ",
	SiteID:          "12345",
	BusinessUnit:    "GMBU",
	Industry:        "Cause & Cure",
	Payer:           "001d000001TwuXwAAZ",
	PrimaryStreet:   "Primary Address Ln",
	PrimaryCity:     "Primary City",
	PrimaryState:    "SC",
	PrimaryZipCode:  "29492",
	PrimaryCountry:  "USA",
	BillingStreet:   "Billing Address Ln",
	BillingCity:     "BillingCity",
	BillingState:    "NC",
	BillingZipCode:  "99999",
	BillingCountry:  "USA",
	ShippingStreet:  "Shipping Address Ln",
	ShippingCity:    "Shiptown",
	ShippingState:   "FL",
	ShippingZipCode: "11111",
	ShippingCountry: "SWE",
}

var accountService = AccountService{AccountRepo: mockAccountRepository{}}

type mockAccountRepository struct {
}

func (m mockAccountRepository) GetAccount(id string) (*AccountDTO, error) {
	return &accountDTO, nil
}

func (m mockAccountRepository) CreateAccount(account *entities.Account) (id string, siteID int, err error) {
	return "001d000001TwuXwAAJ", 12345, nil
}

func (m mockAccountRepository) UpdateAccount(account *entities.Account) error {
	return nil
}

func TestAccountDTOToEntity(t *testing.T) {
	Convey("Given an Account Data Transfer Object with an empty name", t, func() {
		accountDTOCopy := accountDTO
		accountDTOCopy.Name = ""
		Convey("When it is converted to an Account entity", func() {
			_, err := accountDTOCopy.toEntity()
			Convey("Then the conversion should fail", func() {
				So(err, ShouldNotBeNil)
			})
		})
	})
	Convey("Given an Account Data Transfer Object with an invalid SiteID", t, func() {
		accountDTOCopy := accountDTO
		accountDTOCopy.SiteID = "-1234"
		Convey("When it is converted to an Account entity", func() {
			_, err := accountDTOCopy.toEntity()
			Convey("Then the conversion should fail", func() {
				So(err, ShouldNotBeNil)
			})
		})
	})
	Convey("Given an Account Data Transfer Object with a non-int SiteID", t, func() {
		accountDTOCopy := accountDTO
		accountDTOCopy.SiteID = "a1234"
		Convey("When it is converted to an Account entity", func() {
			_, err := accountDTOCopy.toEntity()
			Convey("Then the conversion should fail", func() {
				So(err, ShouldNotBeNil)
			})
		})
	})
	Convey("Given an Account Data Transfer Object with an invalid BU", t, func() {
		accountDTOCopy := accountDTO
		accountDTOCopy.BusinessUnit = "INVALID"
		Convey("When it is converted to an Account entity", func() {
			_, err := accountDTOCopy.toEntity()
			Convey("Then the conversion should fail", func() {
				So(err, ShouldNotBeNil)
			})
		})
	})
	Convey("Given an Account DTO with fully populated fields", t, func() {
		Convey("When it is converted to an Account Entity", func() {
			accountEntity, err := accountDTO.toEntity()
			Convey("Then the conversion should succeed and the fields should match", func() {
				So(err, ShouldBeNil)
				So(accountEntity.Name(), ShouldEqual, accountDTO.Name)
				So(strconv.Itoa(accountEntity.SiteID()), ShouldEqual, accountDTO.SiteID)
				So(accountEntity.BusinessUnit(), ShouldEqual, accountDTO.BusinessUnit)
				So(accountEntity.Industry, ShouldEqual, accountDTO.Industry)
				So(accountEntity.Payer, ShouldEqual, accountDTO.Payer)
				So(accountEntity.PrimaryAddress.Street, ShouldEqual, accountDTO.PrimaryStreet)
				So(accountEntity.PrimaryAddress.City, ShouldEqual, accountDTO.PrimaryCity)
				So(accountEntity.PrimaryAddress.State, ShouldEqual, accountDTO.PrimaryState)
				So(accountEntity.PrimaryAddress.ZipCode, ShouldEqual, accountDTO.PrimaryZipCode)
				So(accountEntity.PrimaryAddress.Country, ShouldEqual, accountDTO.PrimaryCountry)
				So(accountEntity.BillingAddress.Street, ShouldEqual, accountDTO.BillingStreet)
				So(accountEntity.BillingAddress.City, ShouldEqual, accountDTO.BillingCity)
				So(accountEntity.BillingAddress.State, ShouldEqual, accountDTO.BillingState)
				So(accountEntity.BillingAddress.ZipCode, ShouldEqual, accountDTO.BillingZipCode)
				So(accountEntity.BillingAddress.Country, ShouldEqual, accountDTO.BillingCountry)
				So(accountEntity.ShippingAddress.Street, ShouldEqual, accountDTO.ShippingStreet)
				So(accountEntity.ShippingAddress.City, ShouldEqual, accountDTO.ShippingCity)
				So(accountEntity.ShippingAddress.State, ShouldEqual, accountDTO.ShippingState)
				So(accountEntity.ShippingAddress.ZipCode, ShouldEqual, accountDTO.ShippingZipCode)
				So(accountEntity.ShippingAddress.Country, ShouldEqual, accountDTO.ShippingCountry)
			})
		})
	})
	Convey("Given an Account DTO with partially populated fields", t, func() {
		Convey("When it is converted to an Account Entity", func() {
			accountDTOCopy := accountDTO
			accountDTOCopy.PrimaryStreet = ""
			accountDTOCopy.PrimaryCity = ""
			accountDTOCopy.PrimaryState = ""
			accountDTOCopy.PrimaryZipCode = ""
			accountDTOCopy.PrimaryCountry = ""
			accountDTOCopy.BillingStreet = ""
			accountDTOCopy.BillingCity = ""
			accountDTOCopy.BillingState = ""
			accountDTOCopy.BillingZipCode = ""
			accountDTOCopy.BillingCountry = ""
			accountDTOCopy.ShippingStreet = ""
			accountDTOCopy.ShippingCity = ""
			accountDTOCopy.ShippingState = ""
			accountDTOCopy.ShippingZipCode = ""
			accountDTOCopy.ShippingCountry = ""
			accountEntity, err := accountDTOCopy.toEntity()
			Convey("Then the conversion should succeed and the fields should match", func() {
				So(err, ShouldBeNil)
				So(accountEntity.Name(), ShouldEqual, accountDTO.Name)
				So(strconv.Itoa(accountEntity.SiteID()), ShouldEqual, accountDTO.SiteID)
				So(accountEntity.BusinessUnit(), ShouldEqual, accountDTO.BusinessUnit)
				So(accountEntity.Industry, ShouldEqual, accountDTO.Industry)
				So(accountEntity.Payer, ShouldEqual, accountDTO.Payer)
				So(accountEntity.PrimaryAddress, ShouldBeNil)
				So(accountEntity.BillingAddress, ShouldBeNil)
				So(accountEntity.ShippingAddress, ShouldBeNil)
			})
		})
	})
}

func TestAccountEntityToDTO(t *testing.T) {
	Convey("Given an Account with fully populated fields", t, func() {
		accountEntity, _ := accountDTO.toEntity()
		Convey("When it is converted to an Account DTO", func() {
			dto := ConvertAccountEntityToAccountDTO(accountEntity)
			Convey("Then the fields should match", func() {
				So(dto.Name, ShouldEqual, accountEntity.Name())
				So(dto.SiteID, ShouldEqual, strconv.Itoa(accountEntity.SiteID()))
				So(dto.BusinessUnit, ShouldEqual, string(accountEntity.BusinessUnit()))
				So(dto.Industry, ShouldEqual, accountEntity.Industry)
				So(dto.Payer, ShouldEqual, accountEntity.Payer)
				So(dto.PrimaryStreet, ShouldEqual, accountEntity.PrimaryAddress.Street)
				So(dto.PrimaryCity, ShouldEqual, accountEntity.PrimaryAddress.City)
				So(dto.PrimaryState, ShouldEqual, accountEntity.PrimaryAddress.State)
				So(dto.PrimaryZipCode, ShouldEqual, accountEntity.PrimaryAddress.ZipCode)
				So(dto.PrimaryCountry, ShouldEqual, accountEntity.PrimaryAddress.Country)
				So(dto.BillingStreet, ShouldEqual, accountEntity.BillingAddress.Street)
				So(dto.BillingCity, ShouldEqual, accountEntity.BillingAddress.City)
				So(dto.BillingState, ShouldEqual, accountEntity.BillingAddress.State)
				So(dto.BillingZipCode, ShouldEqual, accountEntity.BillingAddress.ZipCode)
				So(dto.BillingCountry, ShouldEqual, accountEntity.BillingAddress.Country)
				So(dto.ShippingStreet, ShouldEqual, accountEntity.ShippingAddress.Street)
				So(dto.ShippingCity, ShouldEqual, accountEntity.ShippingAddress.City)
				So(dto.ShippingState, ShouldEqual, accountEntity.ShippingAddress.State)
				So(dto.ShippingZipCode, ShouldEqual, accountEntity.ShippingAddress.ZipCode)
				So(dto.ShippingCountry, ShouldEqual, accountEntity.ShippingAddress.Country)
			})
		})
	})
	Convey("Given an Account with partially populated fields", t, func() {
		accountEntity, _ := accountDTO.toEntity()
		accountEntity.PrimaryAddress = nil
		accountEntity.BillingAddress = nil
		accountEntity.ShippingAddress = nil
		Convey("When it is converted to an Account DTO", func() {
			dto := ConvertAccountEntityToAccountDTO(accountEntity)
			Convey("Then the fields should match", func() {
				So(dto.Name, ShouldEqual, accountEntity.Name())
				So(dto.SiteID, ShouldEqual, strconv.Itoa(accountEntity.SiteID()))
				So(dto.BusinessUnit, ShouldEqual, string(accountEntity.BusinessUnit()))
				So(dto.Industry, ShouldEqual, accountEntity.Industry)
				So(dto.Payer, ShouldEqual, accountEntity.Payer)
				So(dto.PrimaryStreet, ShouldBeEmpty)
				So(dto.PrimaryCity, ShouldBeEmpty)
				So(dto.PrimaryState, ShouldBeEmpty)
				So(dto.PrimaryZipCode, ShouldBeEmpty)
				So(dto.PrimaryCountry, ShouldBeEmpty)
				So(dto.BillingStreet, ShouldBeEmpty)
				So(dto.BillingCity, ShouldBeEmpty)
				So(dto.BillingState, ShouldBeEmpty)
				So(dto.BillingZipCode, ShouldBeEmpty)
				So(dto.BillingCountry, ShouldBeEmpty)
				So(dto.ShippingStreet, ShouldBeEmpty)
				So(dto.ShippingCity, ShouldBeEmpty)
				So(dto.ShippingState, ShouldBeEmpty)
				So(dto.ShippingZipCode, ShouldBeEmpty)
				So(dto.ShippingCountry, ShouldBeEmpty)
			})
		})
	})
}

func TestGetAccount(t *testing.T) {
	Convey("Given a valid account ID and an AccountService", t, func() {
		id := accountDTO.SiteID
		Convey("When an account is requested from the AccountService", func() {
			account, _ := accountService.GetAccount(id)
			Convey("Then an Account Data Transfer Object is returned", func() {
				So(account, ShouldPointTo, &accountDTO)
			})
		})
	})
}

func TestCreateAccount(t *testing.T) {
	Convey("Given a valid Account DTO", t, func() {
		Convey("When an account is created through the AccountService", func() {
			id, siteID, _ := accountService.CreateAccount(accountDTO)
			Convey("Then an ID is returned", func() {
				So(id, ShouldNotBeEmpty)
				So(siteID, ShouldBeGreaterThan, 0)
			})
		})
	})
	Convey("Given an invalid Account DTO", t, func() {
		accountDTOCopy := accountDTO
		accountDTOCopy.Name = ""
		Convey("When an account is created through the AccountService", func() {
			id, siteID, err := accountService.CreateAccount(accountDTOCopy)
			Convey("Then an error should occur", func() {
				So(id, ShouldEqual, "")
				So(siteID, ShouldEqual, 0)
				So(err, ShouldNotBeNil)
			})
		})
	})
}

func TestUpdateAccount(t *testing.T) {
	Convey("Given a valid Account DTO", t, func() {
		Convey("When an account is updated through the AccountService", func() {
			err := accountService.UpdateAccount(accountDTO)
			Convey("Then there should not be an error", func() {
				So(err, ShouldBeNil)
			})
		})
	})
	Convey("Given an invalid Account DTO", t, func() {
		accountDTOCopy := accountDTO
		accountDTOCopy.Name = ""
		Convey("When an account is updated through the AccountService", func() {
			err := accountService.UpdateAccount(accountDTOCopy)
			Convey("Then an error should occur", func() {
				So(err, ShouldNotBeNil)
			})
		})
	})
}
