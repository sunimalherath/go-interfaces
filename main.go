package main

import "fmt"

type bot interface {
	getGreeting() string // any types that use this function will be of type 'bot'
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

func (eb englishBot) getGreeting() string {
	// Since englishBot is using the getGreeting function, they also become type 'bot'
	// this function and the function for spanishBot greeting assumed to have nothing in common
	return "Hello!"
}

func (sb spanishBot) getGreeting() string {
	// Since spanishBot is using the getGreeting function, they also become type 'bot'
	return "Hola!"
}

/*
Creating the same function name with different parameters to print the greetings.
But Go won't allow function overloading.
Both printGreeting functions have the same signature, so can be refactor to use a common function instead of
having 2 seperate functions.
*/

/*

func printGreeting(eb englishBot) {
	fmt.Println(eb.getGreeting())
}

func printGreeting(sb spanishBot) {
	fmt.Println(sb.getGreeting())
}

*/
