// https://exercism.org/tracks/go/exercises/proverb

package main

import "fmt"

func main() {
	rhyme := []string{"nail", "shoe", "horse", "rider", "message", "battle", "kingdom"}
	fmt.Println(Proverb(rhyme))
}

func Proverb(rhyme []string) []string {
	if len(rhyme) == 0 {
		return []string{}
	}

	for i := 0; i < len(rhyme)-1; i++ {
		fmt.Printf("For want of a %s the %s was lost.\n", rhyme[i], rhyme[i+1])
	}
	fmt.Printf("And all for the want of a %s.\n", rhyme[0])
	return []string{}
}
