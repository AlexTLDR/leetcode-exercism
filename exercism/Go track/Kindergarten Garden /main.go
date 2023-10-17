// https://exercism.org/tracks/go/exercises/kindergarten-garden

package kindergarten

import (
	"errors"
	"strings"
)

// Garden represents the garden with plant assignments.
type Garden struct {
	rows     []string
	children []string
}

// Define a map to translate plant abbreviations to their full names.
var plantMap = map[rune]string{
	'V': "Violets",
	'R': "Radishes",
	'C': "Clover",
	'G': "Grass",
}

// NewGarden creates a new garden with the given diagram and children.
func NewGarden(diagram string, children []string) (*Garden, error) {
	rows := strings.Split(diagram, "\n")
	if len(rows) != 2 {
		return nil, errors.New("Invalid garden diagram format")
	}

	return &Garden{
		rows:     rows,
		children: children,
	}, nil
}

// Plants returns the plants assigned to the specified child.
func (g *Garden) Plants(child string) ([]string, bool) {
	childIndex := -1
	for i, name := range g.children {
		if name == child {
			childIndex = i
			break
		}
	}

	if childIndex == -1 {
		return nil, false
	}

	startCol := childIndex * 2
	childPlants := make([]string, 0)
	for i := 0; i < 2; i++ {
		row := g.rows[i]
		for j := startCol; j < startCol+2; j++ {
			plantAbbreviation := rune(row[j])
			plantName := plantMap[plantAbbreviation]
			childPlants = append(childPlants, plantName)
		}
	}

	return childPlants, true
}
