package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
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

func sendRecipeReq(ingredients []string) (*recipeReq, error) {
	url := recipeURL + "?i=" + strings.Join(ingredients, ",")
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if s := resp.StatusCode; s != http.StatusOK {
		return nil, fmt.Errorf("status '%d' encountered with giphy url: '%s'", s, url)
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
