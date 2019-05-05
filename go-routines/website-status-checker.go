package main

import (
	"fmt"
	"net/http"
)

func main() {
	websites := []string{
		"http://amazon.com",
		"http://facebook.com",
		"http://golang.org",
		"http://google.com",
	}

	c := make(chan string) // this is how to create a channel

	// c will be treated like any other variable all scoping rules will be applied to it just like any other variable

	for _, link := range websites {
		go checkLink(link, c) // this simple go keyword at the beginning of the function call
		// spawns a new go routine(child go routine)

		// we are passing the channel c here. In order to pass data to and from the main go routine to the child go routine, we
		// need to pass c as an argument

		/*
			With go routines, once the go routines are spawned the only way by which the child go routines can communicate with the parent or main go routine is using channels
			Channels are typed which means that the data that we pass through a channel is always typed which means that we cannot send data of say float type into a channel of type string.
		*/
	}

	// fmt.Println(<-c) // receive whatever is coming from the channel
	/*
		This is a blocking call
	*/

	for { // this is like a while True loop ;)
		go checkLink(<-c, c)
	}

	// The above code is actually equivalent to
	// for l := range c { // we can use range in this way to get a value out of a channel
	// go checkLink(l, c)
	// }
}

func checkLink(link string, c chan string) { // here in the function prototype, we not only have to define the type of c
	// as chan but also declare what type of data will c be passing.
	_, err := http.Get(link)
	if err != nil {
		fmt.Println(link, "might be down")
		c <- link //basically pass that given value into the channel
		return
	}
	fmt.Println(link, "is up!")
	c <- link
}
