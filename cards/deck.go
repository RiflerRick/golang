package main

import (
	"fmt"
	"os"
	"rand"
	"strings" // this is the way to import
	"time"
	"io/ioutil"
)

// Create a new type of 'deck' which is slice of strings

type deck []string // this new type called deck extends behaviours of strings
/*
Keep in mind that words like subclass and extends and the like are not actually words that are used in the context of go. It is just helping us to understand what deck type basically is
*/

func newDeck() deck {
	// returns a deck
	// basically creates and returns a new deck
	cards := deck{}
	cardSuits := []string{"Spades", "Diamonds", "Hearts", "Clubs"}
	cardValues := []string{"Ace", "Two", "Three", "Four"} // and so on

	for _, suite := range cardSuits { // when we know that there is a variable(like the index here) 
		// that we really do not care about, we can use the underscore
		for _, val := range cardValues {
			cards = append(cards, suite+" of "+val) // string concatenation works the same way
		}
	}

	return cards
}

/*
This is called a receiver function of deck type
(d deck) deck is obviously the type, d is the instance ofcourse
By convention the receiver variable is always a one or two letter variable that matches the type of the receiver
*/
func (d deck) print() { // be mindful here that the following implementation is also valid: 
	// `func (d deck) print() deck` where we are actually printing 
	for i, card := range d { // d here is the receiver variable, this can be considered to be very similar to self in python
		fmt.Println(i, card)
	}
}

func deal(d deck, handSize int) (deck, deck) { // here we are using parameters
	/*
	this function returns 2 values, both of type deck, thats why we are using deck, deck
	*/
	return d[:handSize], d[handSize:]

}

func (d deck) toString() string {
	return strings.Join([]string(d), ",") // join basically joins each element of the slice with the separator
}

func (d deck) saveToFile(filename string) error {
	return ioutil.WriteFile(filename, []byte(deck.toString()), 0666)
}

func newDeckFromFile(filename string) deck {
	bs, err := ioutil.ReadFile(filename)
	// err is a value of type Error
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}

	s := strings.Split(string(bs), ",")
	return deck(s)

}

func (d deck) shuffle {
	source := rand.NewSource(time.Now().UnixNano()) // this is gonna change the seed every time
	r := rand.New(source)

	for i := range d {
		newPosition := r.Intn(len(d) - 1)
		d[i], d[newPosition] = d[newPosition], d[i]
	}
}