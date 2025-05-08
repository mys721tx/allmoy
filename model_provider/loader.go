package model_provider

import (
	"sync"

	"allmoy/config"
)

type ModelInfo struct {
	ID      string `json:"id"`
	Object  string `json:"object"`
	OwnedBy string `json:"owned_by"`
	Source  string `json:"-"`
	APIUrl  string `json:"-"`
	APIKey  string `json:"-"`
}

var (
	models []ModelInfo
	mu     sync.RWMutex
)

func AddModels(newModels []ModelInfo) {
	mu.Lock()
	defer mu.Unlock()

	filtered := []ModelInfo{}
	for _, m := range models {
		if m.Source != newModels[0].Source {
			filtered = append(filtered, m)
		}
	}

	models = append(filtered, newModels...)
}

func GetAllModels() []ModelInfo {
	mu.RLock()

	models = []ModelInfo{}
	for _, p := range config.LoadedConfig.Providers {
		models = append(models, loadProviderModels(p)...)
	}

	defer mu.RUnlock()

	return models
}

func GetModel(modelID string) *ModelInfo {
	mu.RLock()
	defer mu.RUnlock()

	for _, m := range models {
		if m.ID == modelID {
			return &m
		}
	}
	return nil
}

func loadProviderModels(p config.Provider) []ModelInfo {
	switch p.Type {
	case "openai":
		return loadOpenAI(p)
	case "ollama":
		return loadOllama(p)
	default:
		return nil
	}
}
