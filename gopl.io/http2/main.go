package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
)

func main() {
	db := database{"shoes": 50, "socks": 3}

	mux := http.NewServeMux()

	/*either option of handlefunc and handle are valid */
	// mux.Handle("/list", http.HandlerFunc(db.list))
	mux.HandleFunc("/list", db.list)
	mux.Handle("/price", http.HandlerFunc(db.price))
	mux.HandleFunc("/update", db.update)
	mux.HandleFunc("/delete", db.delete)
	log.Fatal(http.ListenAndServe("localhost:8000", mux))
}

type database map[string]dollars
type dollars float32

func (db database) list(w http.ResponseWriter, r *http.Request) {
	for item, price := range db {
		fmt.Fprintf(w, "%s: %s\n", item, price)
	}
}
func (d dollars) String() string {
	return fmt.Sprintf("$%.2f", d)
}
func (db database) price(w http.ResponseWriter, r *http.Request) {
	item := r.URL.Query().Get("item")
	price, ok := db[item]
	if !ok {
		w.WriteHeader(http.StatusNotFound) //404
		fmt.Fprintf(w, "no such item %q\n", item)
		return
	}
	fmt.Fprintf(w, "%s\n", price)
}
func (db database) update(w http.ResponseWriter, r *http.Request) {
	item := r.URL.Query().Get("item")
	price := r.URL.Query().Get("price")
	f, err := strconv.ParseFloat(price, 32)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "badly formatted price %q\n", f)
		return
	}
	db[item] = dollars(f)
	fmt.Fprintf(w, "Updated %s: %s\n", item, price)

}
func (db database) delete(w http.ResponseWriter, r *http.Request) {
	item := r.URL.Query().Get("item")
	_, ok := db[item]
	if !ok {
		w.WriteHeader(http.StatusNotFound) //404
		fmt.Fprintf(w, "no such item %q\n", item)
		return
	}
	delete(db, item)
	fmt.Fprintf(w, "deleted item: %s", item)

}
