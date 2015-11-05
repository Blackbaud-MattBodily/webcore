package services

// AssetQueryBuilder is an interface for generating asset query strings
type AssetQueryBuilder interface {
	BuildAssetsByAccountIDQuery(accountID string) string
}

// AssetRepository is an inteface for accessing asset data for an account
type AssetRepository interface {
	AssetQueryBuilder
	QueryAssets(query string) ([]*AssetDTO, error)
}

// AssetDTO is a data transfer obect for assets
type AssetDTO struct {
	ProductLine  string     `json:"productLine,omitempty" force:"Product_Line__c,omitempty"`
	EndDate      CustomDate `json:"endDate,omitempty" force:"End_Date__c,omitempty"`
	MaterialType string     `json:"materialType,omitempty" force:"Material_Type__c,omitempty"`
}

// AssetService provides interaction with Asset data
type AssetService struct {
	AssetRepo AssetRepository
}

// QueryAssets returns the assets for the provides query
func (as *AssetService) QueryAssets(query string) ([]*AssetDTO, error) {
	assets, err := as.AssetRepo.QueryAssets(query)
	return assets, err
}

// GetAssetsByAccountID returns assets for the given accountID
func (as *AssetService) GetAssetsByAccountID(accountID string) ([]*AssetDTO, error) {
	query := as.AssetRepo.BuildAssetsByAccountIDQuery(accountID)
	assets, err := as.QueryAssets(query)
	return assets, err
}
