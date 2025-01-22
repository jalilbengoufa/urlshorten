package handlers

import (
	"net/http"
	"strings"
	"urlshorten/internal/store"
)

func RedirectURL(w http.ResponseWriter, r *http.Request) {

	if r.Method != "GET" {
		http.Error(w, "Method Invalid: "+r.Method, http.StatusBadRequest)
		return
	}

	path := r.URL.Path
	shortcode := strings.TrimPrefix(path, "/")

	if shortcode == "" {
		http.Error(w, "Shortcode not provided", http.StatusBadRequest)
		return
	}

	urlFound, ok := store.DataStore.Load(shortcode)
	if ok {

		nbRedirect, ok := store.DataStoreStat.Load(shortcode)
		if !ok {
			http.Error(w, "Shortcode not found", http.StatusNotFound)
			return
		}
		intNbRedirect, _ := nbRedirect.(int)
		url, _ := urlFound.(string)
		store.DataStoreStat.Store(shortcode, intNbRedirect+1)
		http.Redirect(w, r, url, http.StatusPermanentRedirect)

	}

}
