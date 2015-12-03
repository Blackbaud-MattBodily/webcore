package servicebus

import (
	"encoding/xml"
	"fmt"

	"github.com/blackbaudIT/webcore/services"
)

var caseEndpoint = "/Clarify/CaseService"

type CaseEnvelope struct {
	XMLName xml.Name
	Body    CaseBody
}

type CaseBody struct {
	XMLName  xml.Name
	Response GetCaseByClarifySiteIdResponse `xml:"GetCasesByClarifySiteIdResponse"`
}

type GetCaseByClarifySiteIdResponse struct {
	XMLName xml.Name `xml:"GetCasesByClarifySiteIdResponse"`
	Message CaseMessage
}

type CaseMessage struct {
	XMLName   xml.Name `xml:"CaseMessage"`
	CasesElem Cases
}

type Cases struct {
	XMLName   xml.Name            `xml:"Cases"`
	CaseSlice []*services.CaseDTO `xml:"Case"`
}

//GetCasesBySiteID returns a slice of CaseDTO pointers from the last 30 days for the given site ID.
func (a API) GetCasesBySiteID(siteID int) ([]*services.CaseDTO, error) {
	response := &CaseEnvelope{}
	action := "http://webservices.blackbaud.com/clarify/case/GetCasesByClarifySiteId"
	body := "<GetCasesByClarifySiteId xmlns=\"http://webservices.blackbaud.com/clarify/case/\">" +
		fmt.Sprintf("<siteId>%d</siteId>", siteID) +
		"<daysBeforeToday>30</daysBeforeToday>" +
		"<condition></condition><family></family>" +
		"</GetCasesByClarifySiteId>"

	data, err := a.Relay.CallEndpoint(caseEndpoint, action, body)

	if err != nil {
		return response.Body.Response.Message.CasesElem.CaseSlice, err
	}

	err = xml.Unmarshal(data, response)

	if err != nil {
		fmt.Println("Failed to unmarshal")
	}

	fmt.Println(response)
	return response.Body.Response.Message.CasesElem.CaseSlice, err
}
