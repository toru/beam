package main

import (
	"fmt"
	"net/http"

	"github.com/toru/beam/pkg/store"
)

func bookmarksResourceHandlerFunc(db store.Store) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "unimplemented")
	}
}

func loadWebHandlers(db store.Store) {
	http.Handle("/bookmarks", bookmarksResourceHandlerFunc(db))
}
