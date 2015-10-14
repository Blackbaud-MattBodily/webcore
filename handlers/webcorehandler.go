package handlers

import (
	"github.com/blackbaudIT/webcore/data/salesforce"
)

//WebcoreHandler is the base struct that contains the config information necessary to perform a webcore API call.
type WebcoreHandler struct {
	API salesforce.API
}

//NewWebcoreHandler creates a new WebcoreHandler with a given API.
func NewWebcoreHandler(api salesforce.API) *WebcoreHandler {
	return &WebcoreHandler{API: api}
}
