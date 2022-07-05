package db

import (
	"fmt"
	"net/http"
	"strconv"
	"sync"

	"github.com/DagmarC/gopl-solutions/ch7/7.12/template"
)

var mu sync.Mutex

type dollars float32

func (d dollars) String() string { return fmt.Sprintf("$%.2f", d) }

type Database map[string]dollars

func (db Database) List(w http.ResponseWriter, req *http.Request) {
	template.ItemsList.Execute(w, db)
}

func (db Database) Price(w http.ResponseWriter, req *http.Request) {
	item := req.URL.Query().Get("item")
	if price, ok := db[item]; ok {
		fmt.Fprintf(w, "%s\n", price)
	} else {
		w.WriteHeader(http.StatusNotFound) // 404
		fmt.Fprintf(w, "no such item: %q\n", item)
	}
}

func (db Database) Read(w http.ResponseWriter, req *http.Request) {
	item := req.URL.Query().Get("item")
	if price, ok := db[item]; ok {
		fmt.Fprintf(w, "%s: %f\n", item, price)
	} else {
		w.WriteHeader(http.StatusNotFound) // 404
		fmt.Fprintf(w, "no such item: %q\n", item)
	}
}

// Create creates item from request query /create?item=socks&price=6
func (db Database) Create(w http.ResponseWriter, req *http.Request) {
	newItem := req.URL.Query().Get("item")
	price := req.URL.Query().Get("price")

	if _, ok := db[newItem]; ok {
		w.WriteHeader(http.StatusNotModified) // 304
		fmt.Fprintf(w, "%s already exists\n", newItem)
		return
	}

	p, err := strconv.ParseFloat(price, 32)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError) // 500
		fmt.Fprintf(w, "%s in wrong format\n", price)
		return
	}

	db[newItem] = dollars(p)
	w.WriteHeader(http.StatusCreated) // 201
	fmt.Fprintf(w, "Created %s: %s\n", newItem, price)

}

// Delete deletes item from request query /delete?item=socks
func (db Database) Delete(w http.ResponseWriter, req *http.Request) {
	item := req.URL.Query().Get("item")

	if _, ok := db[item]; !ok {
		w.WriteHeader(http.StatusNotFound) // 404
		fmt.Fprintf(w, "%s does not exist\n", item)
		return
	}

	mu.Lock()
	delete(db, item)
	mu.Unlock()

	w.WriteHeader(http.StatusOK) // 200

	fmt.Fprintf(w, "Deleted %s.\n", item)
}

// Delete deletes item from request query /delete?item=socks&price=6
func (db Database) Update(w http.ResponseWriter, req *http.Request) {
	item := req.URL.Query().Get("item")
	price := req.URL.Query().Get("price")

	if _, ok := db[item]; !ok {
		w.WriteHeader(http.StatusNotFound) // 404
		fmt.Fprintf(w, "%s does not exist\n", item)
		return
	}

	p, err := strconv.ParseFloat(price, 32)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError) // 500
		fmt.Fprintf(w, "%s in wrong format\n", price)
		return
	}

	mu.Lock()
	db[item] = dollars(p)
	mu.Unlock()

	w.WriteHeader(http.StatusOK) // 200

	fmt.Fprintf(w, "Updated %s: %s\n", item, price)
}
