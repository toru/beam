package main

import (
	"net/http"
	"strconv"
	"strings"

	"github.com/toru/beam/pkg/bookmark"
	"github.com/toru/beam/pkg/store"
)

func render400(w http.ResponseWriter) {
	http.Error(w, strconv.Quote("bad request"), http.StatusBadRequest)
}

func render404(w http.ResponseWriter) {
	http.Error(w, strconv.Quote("not found"), http.StatusNotFound)
}

func splitPath(path string) []string {
	return strings.FieldsFunc(path, func(c rune) bool {
		return c == '/'
	})
}

func bookmarksResourceHandlerFunc(db store.Store) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodPost:
			bookmarkURL := r.PostFormValue("url")
			if len(bookmarkURL) == 0 {
				render400(w)
				return
			}
			item := bookmark.NewItem()
			if err := item.SetURL(bookmarkURL); err != nil {
				render400(w)
				return
			}
			if err := db.WriteBookmark(*item); err != nil {
				render400(w)
				return
			}
			w.Write([]byte(strconv.Quote("lazy ok")))
		default:
			render404(w)
		}
	}
}

func loadWebHandlers(db store.Store) {
	http.Handle("/bookmarks", bookmarksResourceHandlerFunc(db))
}
