package main

import (
	"fmt"
	"net/http"
	"os"
)

type logWriter struct{}

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

	/*
		There is another way we can do the same thing which is also a better way of doing this.
		A process typically can be in one of 2 phases: CPU bound or IO bound.
		IO being a major portion of any process, the Reader and Writer interfaces can help us deal with this problem better.
		Whenever we want to pass around data from one place to another we can use these interfaces as helpers.
	*/

	fmt.Println(string(bs))
}

func (logWriter) Write(bs []byte) (int, error) {
	// by simply calling the function as Write and using the same function signature as the original
	// write function, we are basically implementing the Writer interface with this function
	fmt.Println(string(bs))
	fmt.Println("Just wrote this many bytes:", len(bs))
	return len(bs), nil
}
