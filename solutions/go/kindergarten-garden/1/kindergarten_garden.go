package kindergarten

import (
	"fmt"
	"sort"
	"strings"
)

// Garden represents the mapping of children to the plants they own.
type Garden map[string][]string

var plantNames = map[rune]string{
	'R': "radishes",
	'C': "clover",
	'G': "grass",
	'V': "violets",
}

// NewGarden creates a new kindergarten garden.
func NewGarden(diagram string, children []string) (*Garden, error) {
	// Sort and check for duplicate child names.
	sortedChildren := make([]string, len(children))
	copy(sortedChildren, children)
	sort.Strings(sortedChildren)
	for i := 1; i < len(sortedChildren); i++ {
		if sortedChildren[i] == sortedChildren[i-1] {
			return nil, fmt.Errorf("Child is listed twice: %s", sortedChildren[i])
		}
	}

	// Initialize the garden.
	garden := Garden{}
	if len(diagram) <= 0 || diagram[0] != '\n' {
		return nil, fmt.Errorf("Not a valid garden format")
	}
	rowLen := -1
	for _, row := range strings.Split(diagram[1:], "\n") {
		if rowLen == -1 {
			rowLen = len(row)
		} else if rowLen != len(row) {
			return nil, fmt.Errorf("Not a valid garden format: %q", diagram)
		}
		for p, plantCode := range row {
			plantName, ok := plantNames[plantCode]
			if !ok {
				return nil, fmt.Errorf("Not a valid plant code: %c", plantCode)
			}
			if len(sortedChildren) <= p/2 {
				return nil, fmt.Errorf("Not enough children")
			}
			child := sortedChildren[p/2]
			plants, ok := garden[child]
			if !ok {
				plants = []string{}
			}
			garden[child] = append(plants, plantName)
		}
	}

	return &garden, nil
}

// Plants lists the plants owned by a child in the garden.
func (g Garden) Plants(child string) ([]string, bool) {
	plants, ok := g[child]
	return plants, ok
}
