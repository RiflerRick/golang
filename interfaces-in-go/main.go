package main

import "fmt"

type bot interface {
	getGreeting() string // interfaces are typically used to used to define method sets
}

type englishBot struct{}
type spanishBot struct{}

func main() {
	eb := englishBot{}
	sb := spanishBot{}

	printGreeting(eb)
	printGreeting(sb)

}

func printGreeting(b bot) {
	fmt.Println(b.getGreeting())
}

func (englishBot) getGreeting() string { // receiver function for type englishBot
	return "Hi there"
}

func (spanishBot) getGreeting() string {
	return "Ola"
}
