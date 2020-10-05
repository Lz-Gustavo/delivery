package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

const (
	recipeURL = "http://www.recipepuppy.com/api/"
)

type recipeReq struct {
	Results []result `json:"results"`
}

type result struct {
	Title       string `json:"title"`
	Link        string `json:"href"`
	Ingredients string `json:"ingredients"`
}

// sendRecipeReq searches for recipes on RecipePuppy API, matching the informed ingredients and
// returning the deserialized JSON as a 'recipeReq' structure. Unnecessary fields from request
// are discarded.
func sendRecipeReq(ingredients []string) (*recipeReq, error) {
	url := recipeURL + "?i=" + url.QueryEscape(strings.Join(ingredients, ","))
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if s := resp.StatusCode; s != http.StatusOK {
		return nil, fmt.Errorf("status '%d' encountered with recipe puppy url: '%s'", s, url)
	}

	raw, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	r := &recipeReq{}
	err = json.Unmarshal(raw, r)
	if err != nil {
		return nil, err
	}
	return r, nil
}
