// https://exercism.org/tracks/go/exercises/airport-robot

package main

import "fmt"

func main() {
	SayHello("Alex")
}

type Greeter interface {
	LanguageName() string
	Greet(name string) string
}

func SayHello(name string, greeter Greeter) string {
	return fmt.Sprintf("I can speak %s: %s", greeter.LanguageName(), greeter.Greet(name))
}
