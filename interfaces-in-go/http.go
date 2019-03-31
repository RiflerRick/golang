package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {
	resp, err := http.Get("http://google.com")
	if err != nil {
		fmt.Println("Error :", err)
		os.Exit(1)
	}

	bs := make([]byte, 99999) // make is a built in function that takes the type and number of elements
	// we want the type to be initialized with

	resp.Body.Read(bs)
	/*
		A lot is going on here in this previous line. resp which is the first argument returned by Get is actually a pointer to a struct having Body as one of its members. Body is an interface of type ReadCloser. ReadCloser on the other hand is an interface that encapsulates Reader and Closer interfaces. The Reader interface uses the `Read` function.

		Now since Body is an interface of type ReadCloser which itself is an interface of Reader and Closer and Reader implements Read function, we can directly use Read function.

		Also we are passing a byte slice to the read function for the read function to populate.
		Read function will not automatically resize the byte slice btw and this is something to look out for.
	*/

	fmt.Println(string(bs))
}
