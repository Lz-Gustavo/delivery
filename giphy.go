package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

const (
	giphyURL = "https://api.giphy.com/v1/gifs/search"
)

type giphyReq struct {
	Data []gif `json:"data"`
}

type gif struct {
	URL string `json:"url"`
}

// sendGiphyReq searches for a single specific gif on Giphy API, matching the informed 'tittle'
// keyword and returning the deserialized JSON as a 'giphyReq' structure. Unnecessary fields from
// request are discarded.
func sendGiphyReq(title string) (*giphyReq, error) {
	url := giphyURL + "?api_key=" + url.QueryEscape(giphyAPIKey) + "&q=" + url.QueryEscape(title) + "&limit=1"
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

	g := &giphyReq{}
	err = json.Unmarshal(raw, g)
	if err != nil {
		return nil, err
	}
	return g, nil
}
