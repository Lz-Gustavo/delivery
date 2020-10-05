package main

import (
	"strings"
)

// ResponseJSON ...
type ResponseJSON struct {
	Keywords []string `json:"keywords"`
	Recipes  []Recipe `json:"recipes"`
}

// Recipe ...
type Recipe struct {
	Title       string   `json:"title"`
	Ingredients []string `json:"ingredients"`
	Link        string   `json:"link"`
	Gif         string   `json:"gif"`
}

// getResponseJSON ...
func getResponseJSON(ingredients []string) (*ResponseJSON, error) {
	r, err := sendRecipeReq(ingredients)
	if err != nil {
		return nil, err
	}
	return compileResponseFields(ingredients, r)
}

// compileResponseFields ...
func compileResponseFields(keyw []string, r *recipeReq) (*ResponseJSON, error) {
	recps := make([]Recipe, 0, len(r.Results))
	for _, r := range r.Results {
		g, err := sendGiphyReq(r.Title)
		if err != nil {
			return nil, err
		}

		rp := Recipe{
			Title:       r.Title,
			Link:        r.Link,
			Ingredients: strings.Split(r.Ingredients, ", "),
			Gif:         g.Data[0].URL,
		}
		recps = append(recps, rp)
	}

	return &ResponseJSON{
		Keywords: keyw,
		Recipes:  recps,
	}, nil
}
