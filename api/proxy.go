package api

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"net/url"

	"allmoy/model_provider"
)

type ProxyRequest struct {
	Model string `json:"model"`
}

func ProxyHandler(w http.ResponseWriter, r *http.Request) {
	body, _ := io.ReadAll(r.Body)

	var req ProxyRequest
	json.Unmarshal(body, &req)

	modelInfo := model_provider.GetModel(req.Model)
	if modelInfo == nil {
		http.Error(w, "Unknown model", http.StatusBadRequest)
		return
	}

	endpoint := r.URL.Path[len("/v1/"):]
	path, err := url.JoinPath(modelInfo.APIUrl, endpoint)
	if err != nil {
		http.Error(w, "Error joining path", http.StatusInternalServerError)
		return
	}

	proxyReq, err := http.NewRequest(r.Method, path, io.NopCloser(bytes.NewReader(body)))
	if err != nil {
		http.Error(w, "Request generation error", http.StatusInternalServerError)
		return
	}

	proxyReq.Header.Set("Authorization", "Bearer "+modelInfo.APIKey)
	proxyReq.Header.Set("Content-Type", "application/json")

	resp, err := http.DefaultClient.Do(proxyReq)
	if err != nil {
		http.Error(w, "Backend error", http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()

	responseData, _ := io.ReadAll(resp.Body)

	w.Header().Set("Content-Type", resp.Header.Get("Content-Type"))
	w.WriteHeader(resp.StatusCode)
	w.Write(responseData)
}
