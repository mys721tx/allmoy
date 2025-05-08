package main

import (
	"flag"
	"log"
	"net/http"

	"allmoy/api"
	"allmoy/config"
	"allmoy/model_provider"
)

var (
	address    string
	configPath string
)

func main() {
	flag.StringVar(&address, "address", "0.0.0.0:8080", "Address to bind to")
	flag.StringVar(&configPath, "config", "providers.yaml", "Path to providers.yaml configuration file")
	flag.Parse()

	config.LoadConfig(configPath)
	model_provider.GetAllModels()

	http.HandleFunc("/v1/models", api.ModelsHandler)
	http.HandleFunc("/v1/", api.ProxyHandler)

	log.Println("Allmoy router running on :8080")
	log.Fatal(http.ListenAndServe(address, nil))
}
