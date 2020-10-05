package main

import (
	"encoding/json"
	"net/http"
	"strings"

	"github.com/gorilla/mux"
)

func createRouter() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/recipes/", recipesHandler).
		Methods("GET").
		Schemes("", "http")
	return r
}

func recipesHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		http.Error(w, "", http.StatusBadRequest)
		return
	}

	args := r.URL.Query()
	if _, ok := args["i"]; !ok {
		http.Error(w, "no ingredients informed", http.StatusBadRequest)
		return
	}

	ings := strings.Split(args["i"][0], ",")
	if len(ings) > 3 {
		http.Error(w, "more than 3 ingredients informed", http.StatusBadRequest)
		return
	}

	rec, err := GetResponseJSON(ings)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	raw, err := json.Marshal(rec)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	_, err = w.Write(raw)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	return
}
