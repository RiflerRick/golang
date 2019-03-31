package main

import "fmt"

// main function is necessary in case of any executable package
func main() {
	//var card string = "Ace of spades" // a value of type string will be assigned here, go is statically typed

	//card := "Ace of spades" // this is a better way of assigning values, the := operator must only be used
	// at the time of intializing a variable, not at the time of reassigning a variable

	// Note that we cannot actually use a variable without initializing it with := operator, we cannot for instance use = to initialize a variable

	// card := newCard()

	//cards := []string{newCard(), newCard()} // this is a slice of type of string

	cards := deck{newCard(), newCard()}
	cards = append(cards, "Six of Spades") // append to cards, the append function as you can see here does not actually modify the existing slice, it returns a new slice. Wrt python this is very different because obviously python's list append actually does modify the list itself

	/*iterating over a slice*/
	for i, card := range cards { // i, card is actually thrown away everytime which is why we need to have
		// the := operator
		fmt.Println(i, card) // i is obviously the index
	}

	// looping through the `deck` of cards and printing them
	cards.print()

	cards = newDeck()

	cards.print()

	fmt.Println("printing deal")
	hand, remainingDeck := deal(cards, 5)
	hand.print()
	fmt.Println("-----")
	remainingDeck.print()
}

func newCard() string { // once again in go, it is necessary to mention what type of data is being returned
	return "Five of Diamonds"
}

