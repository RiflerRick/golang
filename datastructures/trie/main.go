package main

import (
	"fmt"
	"net/http"
)

func getValue(c CacheBackend, key string) (string, error) {
	return c.getValue(key)
}

func putValue(c CacheBackend, key string, data string) error {
	return c.putValue(key, data)
}

func putController(cb CacheBackend, key string, value string) error {
	return putValue(cb, key, value)
}

func getController(cb CacheBackend, key string) (string, error) {
	return getValue(cb, key)
}

func main() {

	cacheBackend := "trie"

	fmt.Printf("CACHE_BACKEND initialized as %s", cacheBackend)

	var cache interface{}

	if cacheBackend == "list" {
		cache = new(ListBackend)
		cache.(CacheBackend).init()
		fmt.Println(cache)
	} else {
		cache = new(TrieBackend)
		cache.(CacheBackend).init()
		fmt.Println(cache)

	}
	// r := mux.NewRouter()

	// http.HandleFunc("/", handler)
	// log.Fatal(http.ListenAndServe(":8080", nil))

	http.HandleFunc("/get/", func(w http.ResponseWriter, r *http.Request) {
		key := r.URL.Query().Get("key")
		// value := r.URL.Query().Get("value")
		data, _ := getController(cache.(CacheBackend), key)
		fmt.Fprintf(w, fmt.Sprintf("%s", data))
	})
	http.HandleFunc("/put/", func(w http.ResponseWriter, r *http.Request) {
		key := r.URL.Query().Get("key")
		value := r.URL.Query().Get("value")
		fmt.Println("query params")
		fmt.Println(r.URL.Query())
		fmt.Printf("key: %s\n", key)
		fmt.Printf("value: %s\n", value)
		putController(cache.(CacheBackend), key, value)
	})

	// r.HandleFunc("/put/{key}/{value}", func(w http.ResponseWriter, r *http.Request) {
	// 	vars := mux.Vars(r)
	// 	putController(cache.(CacheBackend), vars["key"], vars["value"])
	// })
	// r.HandleFunc("/get/{key}", func(w http.ResponseWriter, r *http.Request) {
	// 	vars := mux.Vars(r)
	// 	data, err := getController(cache.(CacheBackend), vars["key"])
	// 	fmt.Fprintf(w, fmt.Sprintf("%s, %s", data, err))
	// })
	http.ListenAndServe(":1337", nil)
}
