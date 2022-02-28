package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strings"
)

type Movie struct {
	Title        string
	Poster       string
	Released     string
	ImdbRating   string
	imdbID       string
	Type         string
	TotalSeasons int
}

// SavePoster
// 1. The response.Body is a stream of data, and implements the Reader interface -
// meaning you can sequentially call Read on it, as if it was an open file.
// 2.  The file I'm opening here implements the Writer interface. This is the opposite
// - it's a stream you can call Write on.
// 3. io.Copy "patches" a reader and a writer, consumes the reader stream and writes its contents to a Writer.
func (m *Movie) SavePoster() error {

	// 1. Get the response.Body
	resp, err := http.Get(m.Poster)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	// 2. Open the file
	f, err := os.Create(strings.Join([]string{m.Title, "jpg"}, "."))
	if err != nil {
		return err
	}
	defer f.Close()

	// 3. Copy respose.Body to a file opened.
	_, err = io.Copy(f, resp.Body)
	if err != nil {
		return err
	}

	fmt.Println("INFO: Poster saved.")
	return nil
}

const OMDbBaseURL = "http://www.omdbapi.com/?t=%s&apikey=acbcffe9"

// usage: go run . <movie name>

func main() {
	if len(os.Args) != 2 {
		log.Fatal("Provide the name of the movie")
	}
	movie, err := fetchMovie(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	err = movie.SavePoster()
	if err != nil {
		log.Fatal(err)
	}
}

func fetchMovie(title string) (*Movie, error) {
	m := &Movie{}
	url := fmt.Sprintf(OMDbBaseURL, title)

	fmt.Println("INFO: Fetching url", url)

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if err = json.NewDecoder(resp.Body).Decode(m); err != nil {
		return nil, err
	}

	return m, nil
}
