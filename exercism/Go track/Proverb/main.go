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

	var proverb []string

	for i := 0; i < len(rhyme)-1; i++ {
		proverb = append(proverb, fmt.Sprintf("For want of a %s the %s was lost.", rhyme[i], rhyme[i+1]))
	}
	proverb = append(proverb, fmt.Sprintf("And all for the want of a %s.", rhyme[0]))
	return proverb
}
