package handlers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/blackbaudIT/webcore/Godeps/_workspace/src/github.com/gorilla/mux"
	"github.com/blackbaudIT/webcore/services"
)

// AssetHandler holds an AssetService and uses it to handle standard
// http requests related to Assets.
type AssetHandler struct {
	assetService services.AssetService
}

//NewAssetHandler creates a new AssetHandler given an AssetRepository and AssetService
func NewAssetHandler(service services.AssetService) *AssetHandler {
	return &AssetHandler{assetService: service}
}

// GetAssetsByAccountID returns all Assets for the accountID provided
func (h *AssetHandler) GetAssetsByAccountID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	accountID, ok := vars["accountId"]
	if !ok {
		log.Println("AssetHandler.GetAssetsByAccountID - missing accountID parameter.")
		http.Error(w, http.StatusText(400), 400)
		return
	}

	assets, err := h.assetService.GetAssetsByAccountID(accountID)
	if err != nil {
		log.Printf("AssetHandler.GetAssetsByAccountID failed: %s", err)
		http.Error(w, http.StatusText(500), 500)
		return
	}

	data, err := json.Marshal(assets)
	if err != nil {
		log.Printf("AssetHandler.GetAssetsByAccountID failed to marshal result: %s", err)
		http.Error(w, http.StatusText(500), 500)
		return
	}

	w.Write(data)
}
