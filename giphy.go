package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
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

func sendGiphyReq(title string) (*giphyReq, error) {
	url := giphyURL + "?api_key=" + giphyAPIKey + "&q=" + title + "&limit=1"

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	// regular request isnt working, trying with custom header values
	req.Header.Set("Accept", "text/html,application/xhtml+xml,application/xml;q=0.9,image/webp,*/*;q=0.8")
	req.Header.Set("Accept-Language", "en-GB,en;q=0.5")
	req.Header.Set("Upgrade-Insecure-Requests", "1")
	req.Header.Set("Cache-Control", "no-cache")
	req.Header.Set("Connection", "keep-alive")
	req.Header.Set("Pragma", "no-cache")
	req.Header.Set("User-Agent", "Mozilla/5.0 (X11; Ubuntu; Linux x86_64; rv:81.0) Gecko/20100101 Firefox/81.0")

	client := http.Client{
		Timeout: time.Second,
	}
	resp, err := client.Do(req)

	//resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	fmt.Println(resp)

	if s := resp.StatusCode; s != http.StatusOK {
		return nil, fmt.Errorf("status '%d' encountered with giphy url: '%s'", s, url)
	}

	raw, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	fmt.Println(string(raw))

	g := &giphyReq{}
	err = json.Unmarshal(raw, g)
	if err != nil {
		return nil, err
	}
	return g, nil
}
