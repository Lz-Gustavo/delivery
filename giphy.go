package main

const (
	giphyURL = "http://www.recipepuppy.com/api/"
	apiKEY   = "kkkk"
)

type giphyReq struct {
	Data []gif `json:"data"`
}

type gif struct {
	URL string `json:"url"`
}

func sendGiphyReq(ingredients []string) (*giphyReq, error) {
	return nil, nil
}
