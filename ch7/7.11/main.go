// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 195.

// Http4 is an e-commerce server that registers the /list and /price
// endpoint by calling http.HandleFunc.
package main

import (
	"log"
	"net/http"

	"github.com/DagmarC/gopl-solutions/ch7/7.11/db"
)

//!+main

func main() {
	db := db.Database{"shoes": 50, "socks": 5}
	http.HandleFunc("/list", db.List)
	http.HandleFunc("/price", db.Price)
	http.HandleFunc("/read", db.Read)
	http.HandleFunc("/create", db.Create)
	http.HandleFunc("/delete", db.Delete)
	http.HandleFunc("/update", db.Update)

	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

//!-main
