package main

import "testing"

func TestRecipesExtraction(t *testing.T) {
	testCases := []struct {
		ingredients []string
	}{
		{
			[]string{"onion", "tomato"},
		},
		{
			[]string{"cheese"},
		},
		{
			[]string{"car"},
		},
	}

	for _, tc := range testCases {
		_, err := sendRecipeReq(tc.ingredients)
		if err != nil {
			t.Log("failed recipes req with ingredients: '", tc.ingredients, "', error:", err.Error())
			t.Fail()
		}
	}
}

func TestGiphyExtraction(t *testing.T) {
	testCases := []struct {
		title string
	}{
		{
			"Dehydrating Tomatoes",
		},
		{
			"Pepperoni Pizza",
		},
		{
			"Is there an invalid gif search string? Maybe this one",
		},
	}

	for _, tc := range testCases {
		_, err := sendGiphyReq(tc.title)
		if err != nil {
			t.Log("failed giphy req with title: '", tc.title, "', error:", err.Error())
			t.Fail()
		}
	}
}

func TestResponseJSON(t *testing.T) {
	testCases := []struct {
		ingredients []string
	}{
		{
			[]string{"onion", "tomato"},
		},
		{
			[]string{"cheese"},
		},
		{
			[]string{"car"},
		},
	}

	for _, tc := range testCases {
		_, err := GetResponseJSON(tc.ingredients)
		if err != nil {
			t.Log("failed recipes req with ingredients: '", tc.ingredients, "', error:", err.Error())
			t.Fail()
		}
	}
}
