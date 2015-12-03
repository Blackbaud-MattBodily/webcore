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

type CaseService struct {
	CaseRepo CaseRepository
}

type CaseRepository interface {
	GetCasesBySiteID(siteID int) ([]*CaseDTO, error)
}

func NewCaseService(repo CaseRepository) *CaseService {
	return &CaseService{CaseRepo: repo}
}

func (c *CaseService) GetCasesBySiteID(siteID int) ([]*CaseDTO, error) {
	cases, err := c.CaseRepo.GetCasesBySiteID(siteID)

	return cases, err
}
