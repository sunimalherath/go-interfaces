package main

type englishBot struct{}
type spanishBot struct{}

func main() {

}

func (eb englishBot) getGreeting() string {
	// this function and the function for spanishBot greeting assumed to have nothing in common
	return "Hello!"
}

func (sb spanishBot) getGreeting() string {
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
