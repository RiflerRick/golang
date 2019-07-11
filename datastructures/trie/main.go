package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func get_value(c CacheBackend, key string) (string, error) {
	return c.getValue(key)
}

func put_value(c CacheBackend, key string, data string) error {
	return c.putValue(key, data)
}

func put_controller(cb CacheBackend, key string, value string) error {
	return put_value(cb, key, value)
}

func get_controller(cb CacheBackend, key string) (string, error) {
	return get_value(cb, key)
}

func main() {

	CACHE_BACKEND := "trie"

	fmt.Printf("CACHE_BACKEND initialized as %s", CACHE_BACKEND)

	var cache interface{}

	if CACHE_BACKEND == "list" {
		cache = new(ListBackend)
		cache.(CacheBackend).init()
		fmt.Println(cache)

	} else {
		cache = new(TrieBackend)
		cache.(CacheBackend).init()
		fmt.Println(cache)

	}
	r := mux.NewRouter()

	// http.HandleFunc("/", handler)
	// log.Fatal(http.ListenAndServe(":8080", nil))

	r.HandleFunc("/put/{key}/{value}", func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		// fmt.Printf("key: %s, value: %s", vars["key"], vars["value"])
		// fmt.Fprintf(w, fmt.Sprintf("%s", put_controller(cache.(CacheBackend), vars["key"], vars["value"])))
		put_controller(cache.(CacheBackend), vars["key"], vars["value"])
	})
	r.HandleFunc("/get/{key}", func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		data, err := get_controller(cache.(CacheBackend), vars["key"])
		fmt.Fprintf(w, fmt.Sprintf("%s, %s", data, err))
	})
	http.ListenAndServe(":1337", r)
}
