// https://exercism.org/tracks/go/exercises/party-robot

package main

import "fmt"

func main() {
	fmt.Println(Welcome("Christiane"))
	fmt.Println(HappyBirthday("Frank", 58))
	fmt.Println(AssignTable("Christiane", 27, "Frank", "on the left", 23.7834298))
}

func Welcome(name string) string {
	return fmt.Sprintf("Welcome to my party, %s!", name)
}

// HappyBirthday wishes happy birthday to the birthday person and exclaims their age.
func HappyBirthday(name string, age int) string {
	return fmt.Sprintf("Happy birthday %s! You are now %d years old!", name, age)
}

// AssignTable assigns a table to each guest.
func AssignTable(name string, table int, neighbor, direction string, distance float64) string {
	return fmt.Sprintf("Welcome to my party, %s!\nYou have been assigned to table %03d. Your table is %s, exactly %.1f meters from here.\nYou will be sitting next to %s.",
		name, table, direction, distance, neighbor)
}
