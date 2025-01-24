package handlers

import (
	"net/http"
	"strings"
	"urlshorten/internal/utils"
)

func RedirectURL(context *utils.AppContext, w http.ResponseWriter, r *http.Request) (int, error) {

	if r.Method != "GET" {
		http.Error(w, "Method Invalid: "+r.Method, http.StatusBadRequest)
		return 400, nil
	}

	path := r.URL.Path
	shortcode := strings.TrimPrefix(path, "/")

	if shortcode == "" {
		http.Error(w, "Shortcode not provided", http.StatusBadRequest)
		return 400, nil
	}

	urlFound, ok := context.Store.Codes.Load(shortcode)
	if ok {

		nbRedirect, ok := context.Store.Stats.Load(shortcode)
		if !ok {
			http.Error(w, "Shortcode not found", http.StatusNotFound)
			return 400, nil
		}
		intNbRedirect, _ := nbRedirect.(int)
		url, _ := urlFound.(string)
		context.Store.Stats.Store(shortcode, intNbRedirect+1)
		http.Redirect(w, r, url, http.StatusPermanentRedirect)

	}
	return 200, nil

}
