package xkcd

import (
	"encoding/json"
	"net/http"
)

const (
	XkcdBaseURL    = "https://xkcd.com/%d"
	XkcdJsonSuffix = "/info.0.json"
	MaxCominNumber = 10
)

type Comic struct {
	Title      string
	Month      string
	Day        string
	Num        int
	SafeTitle  string `json:"safe_title"`
	Transcript string
	Alt        string
	Img        string
}

type ComicIndex map[string]*Comic

func GetComic(url string, c chan *Comic, failed chan string) {
	q := url + XkcdJsonSuffix

	comic := new(Comic)

	req, err := http.NewRequest("GET", q, nil)
	if err != nil {
		failed <- q
		return
	}
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		failed <- q
		return
	}
	// Unmarshalling into the comic structure
	if err := json.NewDecoder(resp.Body).Decode(comic); err != nil {
		failed <- url
	}
	c <- comic
}
