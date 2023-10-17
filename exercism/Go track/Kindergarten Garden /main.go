// https://exercism.org/tracks/go/exercises/kindergarten-garden

package kindergarten

import (
	"errors"
	"sort"
	"strings"
)

// Garden is the type that contains all the information about the garden.
type Garden map[string][]string

var plants map[rune]string

func init() {
	plants = map[rune]string{
		'C': "clover",
		'G': "grass",
		'R': "radishes",
		'V': "violets",
	}
}

// NewGarden creates a new garden from the string indicating the location.
// of plants and an array of children's names.
func NewGarden(diagram string, children []string) (*Garden, error) {
	garden := Garden{}

	sorted := sortedChildren(children)

	row1, row2, err := interpretRows(diagram, len(sorted))
	if err != nil {
		return nil, err
	}

	for k, child := range sorted {
		if _, ok := garden[child]; ok {
			return nil, errors.New("duplicate child")
		}

		garden[child] = []string{
			row1[k*2],
			row1[k*2+1],
			row2[k*2],
			row2[k*2+1],
		}
	}

	return &garden, nil
}

// Plants returns the plants that a particular child planted.
func (g *Garden) Plants(child string) ([]string, bool) {
	plants, ok := (*g)[child]
	return plants, ok
}

func sortedChildren(children []string) []string {
	newSlice := make([]string, len(children))
	copy(newSlice, children)
	sort.Strings(newSlice)
	return newSlice
}

func interpretRows(diagram string, numStudents int) ([]string, []string, error) {
	lines := strings.Split(diagram, "\n")

	if lines[0] != "" {
		return nil, nil, errors.New("invalid garden format")
	}

	row1 := []string{}
	row2 := []string{}

	for _, line := range lines[1:] {
		if len(line) != numStudents*2 {
			return nil, nil, errors.New("invalid garden format")
		}

		for i := 0; i < len(line); i += 2 {
			row1 = append(row1, string(line[i]))
			row2 = append(row2, string(line[i+1]))
		}
	}

	return row1, row2, nil
}
