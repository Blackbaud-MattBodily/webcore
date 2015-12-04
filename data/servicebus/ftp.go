package servicebus

import (
	"encoding/xml"
	"fmt"

	"github.com/blackbaudIT/webcore/services"
)

var ftpEndpoint = "/Website/WebAccountService"

//GetFTPCredentials retrieves a given user's (identified by their email)
//FTP credentials from the web-db using the azure servicebus.
func (a API) GetFTPCredentials(email string) (*services.FTPCredentialsDTO, error) {
	response := &ftpEnvelope{}
	action := "http://webservices.blackbaud.com/website/webaccount/GetFTPUserName"
	body := "<GetFTPUserName xmlns=\"http://webservices.blackbaud.com/website/webaccount/\">" +
		fmt.Sprintf("<email>%s</email>", email) +
		"</GetFTPUserName>"

	fmt.Println(body)

	data, err := a.Relay.CallEndpoint(ftpEndpoint, action, body)

	if err != nil {
		return response.Body.Response.Result.FTPCreds, err
	}

	err = xml.Unmarshal(data, response)

	return response.Body.Response.Result.FTPCreds, err
}

//The following structs are only for proper unmarshaling of the Soap response that comes back
//froma  request for Case data.
type ftpEnvelope struct {
	XMLName xml.Name
	Body    ftpBody
}

type ftpBody struct {
	XMLName  xml.Name
	Response getFTPUserNameResponse `xml:"GetFTPUserNameResponse"`
}

type getFTPUserNameResponse struct {
	XMLName xml.Name `xml:"GetFTPUserNameResponse"`
	Result  ftpResult
}

type ftpResult struct {
	XMLName  xml.Name `xml:"GetFTPUserNameResult"`
	FTPCreds *services.FTPCredentialsDTO
}
