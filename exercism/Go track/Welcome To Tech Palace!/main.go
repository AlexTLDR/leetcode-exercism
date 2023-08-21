// https://exercism.org/tracks/go/exercises/welcome-to-tech-palace

package main

import (
	"fmt"
	"strings"
)

func main() {
	message := `
**************************
*    BUY NOW, SAVE 10%   *
**************************
`
	fmt.Println(WelcomeMessage("Judy"))
	fmt.Println(AddBorder("Welcome!", 10))
	fmt.Println(CleanupMessage(message))
}

func WelcomeMessage(customer string) string {
	return "Welcome to the Tech Palace, " + strings.ToUpper(customer)
}

func AddBorder(welcomeMsg string, numStarsPerLine int) string {
	return strings.Repeat("*", numStarsPerLine) + "\n" + welcomeMsg + "\n" + strings.Repeat("*", numStarsPerLine)
}

func CleanupMessage(oldMsg string) string {
	return strings.TrimSpace(strings.ReplaceAll(oldMsg, "*", ""))
}
