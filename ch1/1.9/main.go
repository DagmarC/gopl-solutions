// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 16.
//!+

// Fetch prints the content found at each specified URL.
package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
)

func main() {
	for _, url := range os.Args[1:] {
		// Ex 1.8
		const httpP = "http://"
		const httpsP = "https://"
		if !strings.HasPrefix(url, httpP) && !strings.HasPrefix(url, httpsP) {
			url = httpP + url
		}

		resp, err := http.Get(url)
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: %v\n", err)
			os.Exit(1)
		}

		fmt.Println("Status code:", resp.StatusCode) // Ex 1.9
		// Ex 1.8
		_, err = io.Copy(os.Stdout, resp.Body)
		if err != nil {
			fmt.Fprintf(os.Stderr, "ïo.copy: %v\n", err)
		}
		resp.Body.Close()
	}
}

//!-
