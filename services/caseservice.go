package services

import "encoding/xml"

//CaseDTO is a data transfer object for moving account case data around.
type CaseDTO struct {
	XMLName   xml.Name `xml:"Case"`
	ID        string   `json:"id,omitempty" xml:"Id,attr"`
	Title     string   `json:"title,omitempty" xml:"Title,omitempty"`
	Status    string   `json:"status,omitempty" xml:"Status,omitempty"`
	DateAdded string   `json:"dateAdded,omitempty" xml:"DateAdded,omitempty"`
	WebNotes  string   `json:"notes,omitempty" xml:"WebNotes,omitempty"`
}

//CaseService is a struct that stores the CaseRepository that should be
//communicated with as well as functions for manipulating Case data on that
//repository.
type CaseService struct {
	CaseRepo CaseRepository
}

//CaseRepository is an interface that defines the functions required for an
//object to be considered a CaseRepository.
type CaseRepository interface {
	GetCasesBySiteID(siteID, lookback int) ([]*CaseDTO, error)
}

//NewCaseService returns a pointer to a CaseService instantiated with a given
//CaseRepository.
func NewCaseService(repo CaseRepository) *CaseService {
	return &CaseService{CaseRepo: repo}
}

//GetCasesBySiteID tries to retrieve a slice of CaseDTOs given the siteID of
//an account (likely the Clarify Site ID).
func (c *CaseService) GetCasesBySiteID(siteID, lookback int) ([]*CaseDTO, error) {
	cases, err := c.CaseRepo.GetCasesBySiteID(siteID, lookback)

	return cases, err
}
