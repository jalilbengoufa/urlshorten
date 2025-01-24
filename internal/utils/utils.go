package utils

import (
	"net/http"
	"urlshorten/internal/store"
)

type AppContext struct {
	Store *store.Store
}

type AppHandler struct {
	Context *AppContext
	Handler func(*AppContext, http.ResponseWriter, *http.Request) (int, error)
}

func (appHandler AppHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	status, err := appHandler.Handler(appHandler.Context, w, r)
	if err != nil {

		switch status {
		case http.StatusNotFound:
			http.NotFound(w, r)
		default:
			http.Error(w, http.StatusText(status), status)
		}

	}
}
