package main

import (
	"flag"
	"fmt"
	"strings"

	"github.com/DagmarC/gopl-solutions/ch4/4.12/xkcd"
)

var s = flag.String("s", "sleep", "Search term that will be used to search in transcript")
var total = flag.Int("total", xkcd.MaxCominNumber, "Total number of comics being indexed from day 1 up to total.")

// Example: go run . --total=2000 -s away
func main() {
	flag.Parse()
	fmt.Println("Search tag:", *s)
	fmt.Println("Total tag:", *total)


	var comicIndex xkcd.ComicIndex
	var failedURLs []string
	comicIndex, failedURLs = fetchComics()

	fmt.Println("Failed", failedURLs)
	n := 0
	for u, c := range comicIndex {
		if strings.Contains(c.Transcript, *s) {
			fmt.Println("\nMATCH number", n)
			fmt.Println(u, c.Num, c.Transcript)
			fmt.Println("---------------")
			n++
		}
	}
}

func fetchComics() (xkcd.ComicIndex, []string) {

	comicIndex := make(xkcd.ComicIndex, xkcd.MaxCominNumber)
	failedURLs := make([]string, 0)

	comics := make(chan *xkcd.Comic)
	failedURLch := make(chan string)

	defer close(comics)
	defer close(failedURLch)

	// Fetch comics
	for i := 1; i <= *total; i++ {
		go xkcd.GetComic(fmt.Sprintf(xkcd.XkcdBaseURL, i), comics, failedURLch)
	}

	for i := 1; i <= *total; i++ {
		select {
		case url := <-failedURLch:
			failedURLs = append(failedURLs, url)
		case c := <-comics:
			comicIndex[fmt.Sprintf(xkcd.XkcdBaseURL, c.Num)] = c
		}
	}

	return comicIndex, failedURLs
}
