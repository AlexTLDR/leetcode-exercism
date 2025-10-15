package etl

import (
	"fmt"
	"strings"
)

func Transform(in map[int][]string) map[string]int {
	out := make(map[string]int)

	for i, v := range in {

		for _, s := range v {
			out[strings.ToLower(s)] = i
		}

	}
	fmt.Println(out)
	return out
}
