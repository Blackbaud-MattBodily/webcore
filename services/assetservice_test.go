package services

import (
	"fmt"
	"testing"
	"time"

	. "github.com/smartystreets/goconvey/convey"
)

var tempDate, _ = time.Parse(customDateLayout, "2029-11-10")
var endDate = CustomDate{tempDate}

var assetDTO = AssetDTO{
	ProductLine:  "BBIS",
	EndDate:      endDate,
	MaterialType: "Software",
}
var mock = mockAssetRepository{mockAssetQueryBuilder{}}
var assetService = AssetService{AssetRepo: mock}

type mockAssetRepository struct {
	AssetQueryBuilder
}

func (m mockAssetRepository) QueryAssets(query string) ([]*AssetDTO, error) {
	return []*AssetDTO{&assetDTO}, nil
}

type mockAssetQueryBuilder struct {
}

func (m mockAssetQueryBuilder) BuildAssetsByAccountIDQuery(accountID string) string {
	return fmt.Sprintf("SELECT Product_Line__c, End_Date__c, Material_Type__c "+
		"FROM Client_Asset__c WHERE Account__r.Id = '%s'", accountID)
}

func TestQueryAssets(t *testing.T) {
	Convey("Given a valid account ID", t, func() {
		id := "001d000001TwuXwAAJ"
		Convey("When assets are requested from the AssetService", func() {
			assets, err := assetService.GetAssetsByAccountID(id)
			Convey("Then an Asset Data Transfer Object is returned", func() {
				So(err, ShouldBeNil)
				So(assets[0].ProductLine, ShouldEqual, assetDTO.ProductLine)
				So(assets[0].EndDate.Time.Equal(assetDTO.EndDate.Time), ShouldBeTrue)
				So(assets[0].MaterialType, ShouldEqual, assetDTO.MaterialType)
			})
		})
	})
}

func TestQueryBuilder(t *testing.T) {
	Convey("Given a valid account ID", t, func() {
		id := "001d000001TwuXwAAJ"
		Convey("When a query string is requested from the AssetService", func() {
			query := assetService.AssetRepo.BuildAssetsByAccountIDQuery(id)
			Convey("Then a query is returned", func() {
				So(query, ShouldEqual, "SELECT Product_Line__c, End_Date__c, Material_Type__c FROM Client_Asset__c WHERE Account__r.Id = '001d000001TwuXwAAJ'")
			})
		})
	})
}
