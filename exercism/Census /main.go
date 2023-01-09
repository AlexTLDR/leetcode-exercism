// https://exercism.org/tracks/go/exercises/census

package main

import "fmt"

func main() {

	name := "Matthew Sanabria"
	age := 0
	address := make(map[string]string)
	resident := NewResident(name, age, address)
	fmt.Println(resident)
	fmt.Println(resident.HasRequiredInfo())
	resident.Delete()
	fmt.Println(resident)

}

// Resident represents a resident in this city.
type Resident struct {
	Name    string
	Age     int
	Address map[string]string
}

// NewResident registers a new resident in this city.
func NewResident(name string, age int, address map[string]string) *Resident {
	return &Resident{
		Name:    name,
		Age:     age,
		Address: address,
	}
}

// HasRequiredInfo determines if a given resident has all of the required information.
func (r *Resident) HasRequiredInfo() bool {
	if r.Name == "" || r.Address["street"] == "" {
		return false
	}
	return true
}

// Delete deletes a resident's information.
func (r *Resident) Delete() {
	r.Name = ""
	r.Age = 0
	r.Address = nil
}

// Count counts all residents that have provided the required information.
func Count(residents []*Resident) int {
	panic("Please implement Count.")
}
