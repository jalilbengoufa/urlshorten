package main

import (
	"log"
	"net/http"

	"urlshorten/internal/handlers"
	"urlshorten/internal/store"
	"urlshorten/internal/utils"
)

func main() {

	store := store.NewStore()
	context := &utils.AppContext{Store: store}

	http.Handle("/shorten", utils.AppHandler{Context: context, Handler: handlers.ShortenURL})
	http.Handle("/stats/", utils.AppHandler{Context: context, Handler: handlers.GetStats})
	http.Handle("/", utils.AppHandler{Context: context, Handler: handlers.RedirectURL})

	log.Println("URL Shortener service is running on port 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
