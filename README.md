# Blackbaud WebCore

[ ![Codeship Status for blackbaudIT/webcore](https://codeship.com/projects/8294cd40-566a-0133-7143-166e04e6001d/status?branch=master)](https://codeship.com/projects/109347)

Package | Documentation
--------|--------------
entities|[![GoDoc](https://godoc.org/github.com/blackbaudIT/webcore/entities?status.svg)](https://godoc.org/github.com/blackbaudIT/webcore/entities)
handlers|[![GoDoc](https://godoc.org/github.com/blackbaudIT/webcore/handlers?status.svg)](https://godoc.org/github.com/blackbaudIT/webcore/handlers)
salesforce|[![GoDoc](https://godoc.org/github.com/blackbaudIT/webcore/data/salesforce?status.svg)](https://godoc.org/github.com/blackbaudIT/webcore/data/salesforce)
services|[![GoDoc](https://godoc.org/github.com/blackbaudIT/webcore/services?status.svg)](https://godoc.org/github.com/blackbaudIT/webcore/services)

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
