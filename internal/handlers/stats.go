package handlers

import (
	"encoding/json"
	"net/http"
	"strings"
	"urlshorten/internal/utils"
)

type StatResponse struct {
	StatFound int `json:"statFound"`
}

func GetStats(context *utils.AppContext, w http.ResponseWriter, r *http.Request) (int, error) {
	if r.Method != "GET" {
		http.Error(w, "Method Invalid: "+r.Method, http.StatusBadRequest)
		return 400, nil
	}

	path := r.URL.Path
	shortcode := strings.TrimPrefix(path, "/stats/")

	if shortcode == "" {
		http.Error(w, "Shortcode not provided", http.StatusBadRequest)
		return 400, nil
	}
	statFound, ok := context.Store.Stats.Load(shortcode)
	intStatFound, _ := statFound.(int)
	if ok {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(StatResponse{StatFound: intStatFound})

	}
	return 200, nil

}
