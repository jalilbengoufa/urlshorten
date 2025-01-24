package handlers

import (
	"encoding/json"
	"net/http"
	"urlshorten/internal/services"
	"urlshorten/internal/utils"
)

type URLRequest struct {
	URL string `json:"url"`
}

type ShortenResponse struct {
	Shortcode string `json:"shortcode"`
}

func ShortenURL(context *utils.AppContext, w http.ResponseWriter, r *http.Request) (int, error) {

	if r.Method != "POST" {
		http.Error(w, "Method Invalid: "+r.Method, http.StatusBadRequest)
		return 401, nil
	}

	var urlRquest URLRequest
	err := json.NewDecoder(r.Body).Decode(&urlRquest)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return 401, err

	}
	shortcode, err := services.GenerateShortCode(context)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return 401, err
	}
	context.Store.Codes.Store(shortcode, urlRquest.URL)
	context.Store.Stats.Store(shortcode, 0)

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(ShortenResponse{Shortcode: shortcode})

	return 200, nil

}
