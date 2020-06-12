package main

import (
	"net/http"
	"strconv"
	"strings"

	"github.com/toru/beam/pkg/bookmark"
	"github.com/toru/beam/pkg/store"
)

// BeamApp is an experimental struct that implements the Handler interface.
type BeamApp struct {
	db  store.Store
	mux *http.ServeMux
}

func NewBeamApp(db store.Store) *BeamApp {
	app := &BeamApp{db: db}
	app.mux = http.NewServeMux()
	app.mux.HandleFunc("/bookmarks", app.handleBookmark)
	app.mux.HandleFunc("/bookmarks/count", app.handleBookmarkCount)
	return app
}

func (app *BeamApp) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("server", "beam")
	app.mux.ServeHTTP(w, r)
}

func (app *BeamApp) handleBookmark(w http.ResponseWriter, r *http.Request) {
	tokens := splitPath(r.URL.Path)
	numTokens := len(tokens)
	switch r.Method {
	case http.MethodPost:
		if numTokens > 1 {
			render404(w)
			return
		}
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
		if err := app.db.WriteBookmark(*item); err != nil {
			render400(w)
			return
		}
		w.Write([]byte(strconv.Quote("lazy ok")))
	default:
		render404(w)
	}
}

func (app *BeamApp) handleBookmarkCount(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		render404(w)
		return
	}
	w.Write([]byte(strconv.Itoa(app.db.BookmarkCount())))
}

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
