package main

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"strings"

	"github.com/toru/beam/pkg/bookmark"
	"github.com/toru/beam/pkg/store"
)

func getOnly(next http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			render404(w)
			return
		}
		next.ServeHTTP(w, r)
	})
}

// BeamApp is an experimental struct that implements the Handler interface.
type BeamApp struct {
	cfg config
	db  store.Store
	mux *http.ServeMux
}

func NewBeamApp(cfg config, db store.Store) *BeamApp {
	app := &BeamApp{cfg: cfg, db: db}
	app.mux = http.NewServeMux()
	app.mux.HandleFunc("/bookmarks", app.handleBookmark)
	app.mux.HandleFunc("/bookmarks/", app.handleBookmark)
	app.mux.HandleFunc("/bookmarks/count", getOnly(app.handleBookmarkCount))
	return app
}

func (app *BeamApp) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("server", "beam")
	app.mux.ServeHTTP(w, r)
}

func (app *BeamApp) getBookmarks() []bookmark.Item {
	bookmarks := make([]bookmark.Item, 0, app.db.BookmarkCount())
	for _, bookmark := range app.db.Bookmarks(0) {
		bookmarks = append(bookmarks, bookmark)
	}
	return bookmarks
}

func (app *BeamApp) handleBookmark(w http.ResponseWriter, r *http.Request) {
	tokens := splitPath(r.URL.Path)
	numTokens := len(tokens)
	switch r.Method {
	case http.MethodPost:
		if numTokens > 2 {
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
		if err := app.db.WriteBookmark(item); err != nil {
			render400(w)
			return
		}
		buf, err := json.Marshal(item)
		if err != nil {
			log.Println(err)
			render500(w)
			return
		}
		w.Write(buf)
	case http.MethodGet:
		if numTokens == 1 {
			buf, err := json.Marshal(app.getBookmarks())
			if err != nil {
				log.Println(err)
				render500(w)
				return
			}
			w.Write(buf)
		} else if numTokens == 2 {
			id := tokens[1]
			item, ok := app.db.GetBookmark(id)
			if !ok {
				render404(w)
				return
			}
			buf, err := json.Marshal(item)
			if err != nil {
				log.Println(err)
				render500(w)
				return
			}
			w.Write(buf)
		} else {
			render404(w)
			return
		}
	default:
		render404(w)
	}
}

func (app *BeamApp) handleBookmarkCount(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(strconv.Itoa(app.db.BookmarkCount())))
}

func render500(w http.ResponseWriter) {
	http.Error(w, strconv.Quote("table flip"), http.StatusInternalServerError)
}

func render404(w http.ResponseWriter) {
	http.Error(w, strconv.Quote("not found"), http.StatusNotFound)
}

func render401(w http.ResponseWriter) {
	http.Error(w, strconv.Quote("unauthorized"), http.StatusUnauthorized)
}

func render400(w http.ResponseWriter) {
	http.Error(w, strconv.Quote("bad request"), http.StatusBadRequest)
}

func splitPath(path string) []string {
	return strings.FieldsFunc(path, func(c rune) bool {
		return c == '/'
	})
}
