package salesforce

import (
	"fmt"

	"github.com/blackbaudIT/webcore/services"
)

// SFDCClientAsset wraps the Account data tranfer object so that SFDC fields can be
// mapped onto it
type SFDCClientAsset struct {
	services.AssetDTO
}

// ApiName is the SFDC ApiName of the ClientAsset object
func (s SFDCClientAsset) ApiName() string {
	return "Client_Asset__c"
}

// ExternalIdApiName is the SFDC external id for the ClientAsset object.
func (s SFDCClientAsset) ExternalIdApiName() string {
	return "FI_Reference_ID__c"
}

// BuildAssetsByAccountIDQuery builds the SOQL string to query assets by account
func (a API) BuildAssetsByAccountIDQuery(accountID string) string {
	return fmt.Sprintf("SELECT Product_Line__c, End_Date__c, Material_Type__c "+
		"FROM Client_Asset__c WHERE Account__r.Id = '%s'", accountID)
}

// SFDCClientAssetQueryResponse wraps the base SFDCQueryResponse and attaches a
// slice of SFDCClientAsset pointers that will be written into
type SFDCClientAssetQueryResponse struct {
	SFDCQueryResponse
	Records []*services.AssetDTO `json:"Records" force:"records"`
}

// QueryAssets returns AssetDTO objects based on the query provided
func (a API) QueryAssets(query string) ([]*services.AssetDTO, error) {
	queryResponse := &SFDCClientAssetQueryResponse{}

	err := a.client.QuerySFDCObject(query, queryResponse)
	return queryResponse.Records, err
}
