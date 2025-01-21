package main

import (
	"crypto/rand"
	"encoding/base64"
	"encoding/json"
	"log"
	"net/http"
	"strings"

	"urlshorten/store"
)

type URLRequest struct {
	URL string `json:"url"`
}

type ShortenResponse struct {
	Shortcode string `json:"shortcode"`
}
type StatResponse struct {
	StatFound int `json:"statFound"`
}

func GenerateRandomBytes(n int) ([]byte, error) {
	b := make([]byte, n)
	_, err := rand.Read(b)
	if err != nil {
		return nil, err
	}

	return b, nil
}

func GenerateRandomString(s int) (string, error) {
	b, err := GenerateRandomBytes(s - 2)
	return base64.URLEncoding.EncodeToString(b), err
}

func generateShortCode() (code string, err error) {

	for {

		code, err := GenerateRandomString(6)
		if err != nil {
			return "", err
		}

		_, ok := store.DataStore.Load(code)
		if !ok {
			return code, nil
		}

	}

}

func shortenURL(w http.ResponseWriter, r *http.Request) {

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
	shortcode, err := generateShortCode()
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

func redirectURL(w http.ResponseWriter, r *http.Request) {

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

func getStats(w http.ResponseWriter, r *http.Request) {
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

func main() {
	store.InitStore()
	http.HandleFunc("/shorten", shortenURL)
	http.HandleFunc("/stats/", getStats)
	http.HandleFunc("/", redirectURL)

	log.Println("URL Shortener service is running on port 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
