package handlers

import (
	"encoding/json"
	"net/http"
	"strings"
	"urlshorten/internal/store"
)

type StatResponse struct {
	StatFound int `json:"statFound"`
}

func GetStats(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		http.Error(w, "Method Invalid: "+r.Method, http.StatusBadRequest)
		return
	}

	path := r.URL.Path
	shortcode := strings.TrimPrefix(path, "/stats/")

	if shortcode == "" {
		http.Error(w, "Shortcode not provided", http.StatusBadRequest)
		return
	}
	statFound, ok := store.DataStoreStat.Load(shortcode)
	intStatFound, _ := statFound.(int)
	if ok {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(StatResponse{StatFound: intStatFound})

	}

}
