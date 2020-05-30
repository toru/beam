package main

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/toru/beam/pkg/store"
)

func render404(w http.ResponseWriter) {
	http.Error(w, strconv.Quote("not found"), http.StatusNotFound)
}

func bookmarksResourceHandlerFunc(db store.Store) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodPost:
			fmt.Fprintf(w, strconv.Quote("unimplemented"))
		default:
			render404(w)
		}
	}
}

func loadWebHandlers(db store.Store) {
	http.Handle("/bookmarks", bookmarksResourceHandlerFunc(db))
}
