package main

import "fmt"

type bot interface {
	getGreeting() string
}

type englishBot struct{}
type spanishBot struct{}
type italianBot struct{}
type frenchBot struct{}

func main() {
	bl := []bot{englishBot{}, spanishBot{}, italianBot{}, frenchBot{}}

	for _, b := range bl {
		printGreeting(b)
	}

}

func printGreeting(b bot) {
	fmt.Println(b.getGreeting())
}

func (englishBot) getGreeting() string {
	return "Hi There!"
}

func (spanishBot) getGreeting() string {
	return "Hola!"
}
func (italianBot) getGreeting() string {
	return "Ciao!"
}

func (frenchBot) getGreeting() string {
	return "Salut!"
}
