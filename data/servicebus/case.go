package servicebus

import (
	"encoding/xml"
	"fmt"

	"github.com/blackbaudIT/webcore/services"
)

var caseEndpoint = "/Clarify/CaseService"

//GetCasesBySiteID returns a slice of CaseDTO pointers. The lookback provided is
//the number of days the provider needs to look back when retrieving case data,
//and the siteID is the Site ID (likely Clarify) of the account that the cases
//should be retrieved for.
func (a API) GetCasesBySiteID(siteID, lookback int) ([]*services.CaseDTO, error) {
	response := &caseEnvelope{}
	action := "http://webservices.blackbaud.com/clarify/case/GetCasesByClarifySiteId"
	body := "<GetCasesByClarifySiteId xmlns=\"http://webservices.blackbaud.com/clarify/case/\">" +
		fmt.Sprintf("<siteId>%d</siteId>", siteID) +
		fmt.Sprintf("<daysBeforeToday>%d</daysBeforeToday>", lookback) +
		"<condition></condition><family></family>" +
		"</GetCasesByClarifySiteId>"

	data, err := a.Relay.CallEndpoint(caseEndpoint, action, body)

	if err != nil {
		return response.Body.Response.Message.CasesElem.CaseSlice, err
	}

	err = xml.Unmarshal(data, response)

	return response.Body.Response.Message.CasesElem.CaseSlice, err
}

//The following structs are only for proper unmarshaling of the Soap response that comes back
//froma  request for Case data.
type caseEnvelope struct {
	XMLName xml.Name
	Body    caseBody
}

type caseBody struct {
	XMLName  xml.Name
	Response getCaseByClarifySiteIdResponse `xml:"GetCasesByClarifySiteIdResponse"`
}

type getCaseByClarifySiteIdResponse struct {
	XMLName xml.Name `xml:"GetCasesByClarifySiteIdResponse"`
	Message caseMessage
}

type caseMessage struct {
	XMLName   xml.Name `xml:"CaseMessage"`
	CasesElem cases
}

type cases struct {
	XMLName   xml.Name            `xml:"Cases"`
	CaseSlice []*services.CaseDTO `xml:"Case"`
}
