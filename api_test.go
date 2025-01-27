package main

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestAPIReponse(t *testing.T) {
	testCases := []struct {
		url          string
		expectedCode int
	}{
		{
			"/recipes/?i=onion,tomato",
			http.StatusOK,
		},
		{
			"/recipes/",
			http.StatusBadRequest,
		},
		{
			"/recipes/?i=onion,tomato,garlic,cheese",
			http.StatusBadRequest,
		},
	}

	for _, tc := range testCases {
		req, err := http.NewRequest("GET", tc.url, nil)
		if err != nil {
			t.Fatal(err)
		}

		rr := httptest.NewRecorder()
		hf := http.HandlerFunc(recipesHandler)
		hf.ServeHTTP(rr, req)
		if s := rr.Code; s != tc.expectedCode {
			body, _ := ioutil.ReadAll(rr.Body)
			t.Logf("failed endpoint '%s' request, got status code '%d', message '%s'", tc.url, s, string(body))
			t.Fail()
		}
	}
}
