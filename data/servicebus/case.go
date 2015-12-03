package servicebus

import (
	"encoding/xml"
	"fmt"

	"github.com/blackbaudIT/webcore/services"
)

var caseEndpoint = "/Clarify/CaseService"

//GetCasesBySiteID returns a slice of CaseDTO pointers from the last 30 days for the given site ID.
func (a API) GetCasesBySiteID(siteID int) ([]*services.CaseDTO, error) {
	cases := []*services.CaseDTO{}
	action := "http://webservices.blackbaud.com/clarify/case/GetCasesByClarifySiteID"
	body := "<GetCasesByClarifySiteId xmlns=\"http://webservices.blackbaud.com/clarify/case/\">" +
		fmt.Sprintf("<siteID>%d</siteID>", siteID) +
		"<daysBeforeToday>30</daysBeforeToday>" +
		"</GetCasesByClarifySiteId>"

	data, err := a.Relay.CallEndpoint(caseEndpoint, action, body)

	if err != nil {
		return cases, err
	}

	err = xml.Unmarshal(data, cases)

	return cases, err
}
