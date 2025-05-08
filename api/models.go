package api

import (
	"encoding/json"
	"net/http"

	"allmoy/model_provider"
)

func ModelsHandler(w http.ResponseWriter, r *http.Request) {
	models := model_provider.GetAllModels()
	resp := map[string]interface{}{
		"object": "list",
		"data":   models,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}
