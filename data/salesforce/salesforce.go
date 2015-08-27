/*
Package salesforce provides access to SFDC data.

A Connected App must be configured in SalesForce.  See the documentation here:
https://developer.salesforce.com/page/Digging_Deeper_into_OAuth_2.0_on_Force.com

A default client is included (when calling salesforce.NewAPI()) that will read
the necessary connection information from environmental variables on the machine
running the program. These are the expected environmental variables (case
sensitive):

BBWEBCORE_SFDCVERSION (ex. "v32.0")

BBWEBCORE_SFDCCLIENTID (can be found in your SFDC Connected App settings)

BBWEBCORE_SFDCCLIENTSECRET (can be found in your SFDC Connected App settings)

BBWEBCORE_SFDCUSERNAME

BBWEBCORE_SFDCPASSWORD

BBWEBCORE_SFDCTOKEN

BBWEBCORE_SFDCENVIRONMENT (use either "sandbox" or "production")

*/
package salesforce

import (
	"fmt"

	"github.com/blackbaudIT/webcore/data/salesforce/Godeps/_workspace/src/github.com/nimajalali/go-force/force"
	"github.com/blackbaudIT/webcore/data/salesforce/Godeps/_workspace/src/github.com/spf13/viper"
)

//var forceAPI *force.ForceApi

// API provides access to SalesForce Data
type API struct {
	client sfdcClient
}

// NewAPI returns an API object with a default client
func NewAPI() API {
	getConfigSettings()
	fc := forceClient{getForceAPIClient()}
	return API{client: fc}
}

// SFDCResponse contains the SalesForce response info after an insert/update
type SFDCResponse struct {
	ID           string `force:"id,omitempty"`
	ErrorMessage string `force:"error,omitempty"`
	Success      bool   `force:"success,omitempty"`
}

type sfdcClient interface {
	GetSFDCObject(id string, obj interface{}) (err error)
	GetSFDCObjectByExternalID(id string, obj interface{}) (err error)
	InsertSFDCObject(object interface{}) (resposne SFDCResponse, err error)
	UpsertSFDCObjectByExternalID(id string, obj interface{}) (err error)
}

type forceClient struct {
	*force.ForceApi
}

func (f forceClient) GetSFDCObject(id string, obj interface{}) (err error) {
	sobject, ok := obj.(force.SObject)
	if !ok {
		err = fmt.Errorf("unable to convert data to SObject")
		return err
	}

	err = f.GetSObject(id, sobject)
	return err
}

func (f forceClient) GetSFDCObjectByExternalID(id string, obj interface{}) (err error) {
	sobject, ok := obj.(force.SObject)
	if !ok {
		err = fmt.Errorf("unable to convert data to SObject")
		return err
	}

	err = f.GetSObjectByExternalId(id, sobject)
	return err
}

func (f forceClient) InsertSFDCObject(obj interface{}) (resposne SFDCResponse, err error) {
	sobject, ok := obj.(force.SObject)
	if !ok {
		err = fmt.Errorf("unable to convert data to SObject")
		return SFDCResponse{}, err
	}

	resp, err := f.InsertSObject(sobject)

	sfdcResp := SFDCResponse{}
	if resp != nil {
		sfdcResp.ID = resp.Id
		sfdcResp.ErrorMessage = resp.Errors.Error()
		sfdcResp.Success = resp.Success
	}

	return sfdcResp, err
}

func (f forceClient) UpsertSFDCObjectByExternalID(id string, obj interface{}) (err error) {
	sobject, ok := obj.(force.SObject)
	if !ok {
		err = fmt.Errorf("unable to convert data to SObject")
		return err
	}

	// no response object is returned for upserts
	_, err = f.UpsertSObjectByExternalId(id, sobject)

	return err
}

func getConfigSettings() {
	viper.SetEnvPrefix("bbwebcore")
	viper.AutomaticEnv()
}

func getForceAPIClient() *force.ForceApi {
	forceAPI, err := force.Create(
		viper.GetString("sfdcVersion"),
		viper.GetString("sfdcClientId"),
		viper.GetString("sfdcClientSecret"),
		viper.GetString("sfdcUserName"),
		viper.GetString("sfdcPassword"),
		viper.GetString("sfdcToken"),
		viper.GetString("sfdcEnvironment"),
	)

	if err != nil {
		panic(fmt.Errorf("Fatal error creating forceAPI: %s \n", err))
	}

	return forceAPI
}
