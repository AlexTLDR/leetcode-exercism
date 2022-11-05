// https://exercism.org/tracks/go/exercises/two-fer

package main

import "fmt"

func main() {
	fmt.Println(ShareWith("Bob"))
}

func ShareWith(name string) string {
	if name == "" {
		return "One for you, one for me."
	}
	return "One for " + name + ", one for me."
}
