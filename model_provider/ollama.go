package model_provider

import (
	"encoding/json"
	"log"
	"net/http"
	"net/url"

	"allmoy/config"
)

func loadOllama(p config.Provider) []ModelInfo {

	path, err := url.JoinPath(p.APIUrl, "api/tags")
	if err != nil {
		log.Println("Error joining path:", err)
		return nil
	}

	resp, err := http.Get(path)
	if err != nil {
		log.Println("Request generation error:", err)
		return nil
	}
	defer resp.Body.Close()

	var result struct {
		Models []struct {
			Name string `json:"name"`
		} `json:"models"`
	}
	json.NewDecoder(resp.Body).Decode(&result)

	var models []ModelInfo
	for _, model := range result.Models {
		models = append(models, ModelInfo{
			ID:      model.Name,
			Object:  "model",
			OwnedBy: p.Name,
			Source:  p.Name,
			APIUrl:  p.APIUrl,
			APIKey:  p.APIKey,
		})
	}
	return models
}
