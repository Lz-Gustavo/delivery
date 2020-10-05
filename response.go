package main

import (
	"sort"
	"strings"
)

// ResponseJSON defines the JSON format for the main /recipes endpoint, preserving only
// its requested fields.
type ResponseJSON struct {
	Keywords []string `json:"keywords"`
	Recipes  []Recipe `json:"recipes"`
}

// Recipe defines the format for elements on ResponseJSON's recipes list.
type Recipe struct {
	Title       string   `json:"title"`
	Ingredients []string `json:"ingredients"`
	Link        string   `json:"link"`
	Gif         string   `json:"gif"`
}

// GetResponseJSON returns the matching ResponseJSON for the informed list of ingredients.
func GetResponseJSON(ingredients []string) (*ResponseJSON, error) {
	r, err := sendRecipeReq(ingredients)
	if err != nil {
		return nil, err
	}
	return compileResponseFields(ingredients, r)
}

// compileResponseFields arranges recipes requests from RecipePuppy into the predefined
// ResponseJSON format, querying gifs from Giphy API for each retrieved recipe.
func compileResponseFields(keyw []string, r *recipeReq) (*ResponseJSON, error) {
	recps := make([]Recipe, 0, len(r.Results))
	for _, r := range r.Results {
		g, err := sendGiphyReq(r.Title)
		if err != nil {
			return nil, err
		}

		ings := strings.Split(r.Ingredients, ", ")
		sort.Strings(ings)

		rp := Recipe{
			Title:       r.Title,
			Link:        r.Link,
			Ingredients: ings,
			Gif:         g.Data[0].URL,
		}
		recps = append(recps, rp)
	}

	return &ResponseJSON{
		Keywords: keyw,
		Recipes:  recps,
	}, nil
}
