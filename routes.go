package main

import (
	"github.com/flosch/pongo2"
	"github.com/gorilla/context"
	"log"
	"net/http"
)

var index_page_tmpl = pongo2.Must(pongo2.FromFile("templates/pages/index.html"))
var index_partial_tmpl = pongo2.Must(pongo2.FromFile("templates/partials/index.html"))
var login_page_tmpl = pongo2.Must(pongo2.FromFile("templates/pages/login.html"))
var flex_page_tmpl = pongo2.Must(pongo2.FromFile("templates/pages/flex.html"))
var webrtc_page_tmpl = pongo2.Must(pongo2.FromFile("templates/pages/webrtc.html"))

func indexHandler(w http.ResponseWriter, r *http.Request) {
	ajax := r.Header.Get("X-PUSH")
	log.Println(context.Get(r, UserKey))
	if ajax != "" {
		// serve the partial
		err := index_partial_tmpl.ExecuteWriter(pongo2.Context{"query": r.FormValue("query")}, w)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	} else {
		// serve a page
		err := index_page_tmpl.ExecuteWriter(pongo2.Context{"query": r.FormValue("query")}, w)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	}
}

func loginHandler(w http.ResponseWriter, r *http.Request) {
	err := login_page_tmpl.ExecuteWriter(pongo2.Context{"query": r.FormValue("query")}, w)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func flexHandler(w http.ResponseWriter, r *http.Request) {
	err := flex_page_tmpl.ExecuteWriter(pongo2.Context{"query": r.FormValue("query")}, w)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func webrtcHandler(w http.ResponseWriter, r *http.Request) {
	err := webrtc_page_tmpl.ExecuteWriter(pongo2.Context{"query": r.FormValue("query")}, w)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
