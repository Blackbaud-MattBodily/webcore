package salesforce

import (
	"testing"

	. "github.com/blackbaudIT/webcore/Godeps/_workspace/src/github.com/smartystreets/goconvey/convey"
)

func TestAssetApiName(t *testing.T) {
	Convey("Given an SFDCAsset object", t, func() {
		asset := SFDCClientAsset{}
		Convey("When the API Name is requested", func() {
			apiName := asset.ApiName()
			Convey("Then 'Client_Asset__c' should be returned", func() {
				So(apiName, ShouldEqual, "Client_Asset__c")
			})
		})
	})
}

func TestAssetExternalIdApiName(t *testing.T) {
	Convey("Given an SFDCAsset object", t, func() {
		asset := SFDCClientAsset{}
		Convey("When the External Id API Name is requested", func() {
			apiName := asset.ExternalIdApiName()
			Convey("Then 'FI_Reference_ID__c' should be returned", func() {
				So(apiName, ShouldEqual, "FI_Reference_ID__c")
			})
		})
	})
}
