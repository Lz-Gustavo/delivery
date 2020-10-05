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

	g := &giphyReq{}
	return compileResponseFields(ingredients, r, g)
}

// compileResponseFields ...
func compileResponseFields(keyw []string, r *recipeReq, g *giphyReq) (*ResponseJSON, error) {
	// if len(r.Results) != len(g.Data) {
	// 	return nil, errors.New("different size recipe and gif slices")
	// }

	recps := make([]Recipe, 0, len(r.Results))
	for _, r := range r.Results {
		rp := Recipe{
			Title:       r.Title,
			Link:        r.Link,
			Ingredients: strings.Split(r.Ingredients, ", "),
			Gif:         "TODO",
		}
		recps = append(recps, rp)
	}

	return &ResponseJSON{
		Keywords: keyw,
		Recipes:  recps,
	}, nil
}
