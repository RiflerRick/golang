package main

import (
	"fmt"
	"net/http"
)

func getValue(c CacheBackend, key string, resp chan interface{}) {
	c.getValue(key, resp)
}

func putValue(c CacheBackend, key string, data string, resp chan interface{}) {
	c.putValue(key, data, resp)
}

func putController(cb CacheBackend, key string, value string, resp chan interface{}) {
	putValue(cb, key, value, resp)
}

func getController(cb CacheBackend, key string, resp chan interface{}) {
	getValue(cb, key, resp)
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
	putResponse := make(chan interface{}, 1)
	getResponse := make(chan interface{}, 1)

	http.HandleFunc("/get/", func(w http.ResponseWriter, r *http.Request) {
		key := r.URL.Query().Get("key")
		// value := r.URL.Query().Get("value")
		getController(cache.(CacheBackend), key, getResponse)
		fmt.Println("responding now")
		fmt.Fprintf(w, fmt.Sprintf("%s", <-getResponse))
		fmt.Println("responded")
	})
	http.HandleFunc("/put/", func(w http.ResponseWriter, r *http.Request) {
		key := r.URL.Query().Get("key")
		value := r.URL.Query().Get("value")
		fmt.Println("query params")
		fmt.Println(r.URL.Query())
		fmt.Printf("key: %s\n", key)
		fmt.Printf("value: %s\n", value)
		putController(cache.(CacheBackend), key, value, putResponse)
		fmt.Println("responding now")
		fmt.Fprintf(w, fmt.Sprintf("%s", <-putResponse))
		fmt.Println("responded")
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
