package model_provider

import (
	"encoding/json"
	"log"
	"net/http"
	"net/url"

	"allmoy/config"
)

func loadOpenAI(p config.Provider) []ModelInfo {
	path, err := url.JoinPath(p.APIUrl, "models")
	if err != nil {
		log.Println("Error joining path:", err)
		return nil
	}

	req, err := http.NewRequest("GET", path, nil)
	if err != nil {
		log.Println("Request generation error:", err)
		return nil
	}

	req.Header.Set("Authorization", "Bearer "+p.APIKey)

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Println("Request generation error:", err)
		return nil
	}
	defer resp.Body.Close()

	var result map[string][]map[string]interface{}
	json.NewDecoder(resp.Body).Decode(&result)

	var models []ModelInfo
	for _, model := range result["data"] {
		models = append(models, ModelInfo{
			ID:      model["id"].(string),
			Object:  "model",
			OwnedBy: p.Name,
			Source:  p.Name,
			APIUrl:  p.APIUrl,
			APIKey:  p.APIKey,
		},
		)
	}
	return models
}
