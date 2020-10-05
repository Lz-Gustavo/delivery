package main

import (
	"encoding/json"
	"net/http"

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
		http.Error(w, "no ingredients parameters", http.StatusBadRequest)
		return
	}

	rec, err := sendRecipeReq(args["i"])
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
