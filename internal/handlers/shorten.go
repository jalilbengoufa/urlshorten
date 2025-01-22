package handlers

import (
	"encoding/json"
	"net/http"
	service "urlshorten/internal/services"
	"urlshorten/internal/store"
)

type URLRequest struct {
	URL string `json:"url"`
}

type ShortenResponse struct {
	Shortcode string `json:"shortcode"`
}

func ShortenURL(w http.ResponseWriter, r *http.Request) {

	if r.Method != "POST" {
		http.Error(w, "Method Invalid: "+r.Method, http.StatusBadRequest)
		return
	}

	var urlRquest URLRequest
	err := json.NewDecoder(r.Body).Decode(&urlRquest)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return

	}
	shortcode, err := service.GenerateShortCode()
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	store.DataStore.Store(shortcode, urlRquest.URL)
	store.DataStoreStat.Store(shortcode, 0)

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(ShortenResponse{Shortcode: shortcode})

}
