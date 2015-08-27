# Blackbaud WebCore

See https://godoc.org/github.com/blackbaudIT/webcore for documentation.

See the examples directory for a working example.

#### Setup
In order to connect to SalesForce you must:

1. Setup an OAuth Connected App
https://developer.salesforce.com/page/Digging_Deeper_into_OAuth_2.0_on_Force.com
2. Configure these environmental variables:
  * BBWEBCORE_SFDCVERSION (ex. "v32.0")
  * BBWEBCORE_SFDCCLIENTID (can be found in your SFDC Connected App settings)
  * BBWEBCORE_SFDCCLIENTSECRET (can be found in your SFDC Connected App settings)
  * BBWEBCORE_SFDCUSERNAME
  * BBWEBCORE_SFDCPASSWORD
  * BBWEBCORE_SFDCTOKEN
  * BBWEBCORE_SFDCENVIRONMENT (use either "sandbox" or "production")
