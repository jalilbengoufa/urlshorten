package main

import (
	"log"
	"net/http"

	"urlshorten/internal/handlers"
	"urlshorten/internal/store"
)

func main() {
	store.InitStore()
	http.HandleFunc("/shorten", handlers.ShortenURL)
	http.HandleFunc("/stats/", handlers.GetStats)
	http.HandleFunc("/", handlers.RedirectURL)

	log.Println("URL Shortener service is running on port 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
