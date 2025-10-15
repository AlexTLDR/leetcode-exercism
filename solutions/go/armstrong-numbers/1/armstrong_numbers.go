package armstrong

import (
	"math"
	"strconv"
)

func IsNumber(n int) bool {
	var sum float64
	strNum := strconv.Itoa(n)
	for _, v := range strNum {
		sum += math.Pow(float64(int(v-48)), float64(len(strNum)))
	}
	if n == int(sum) {
		return true
	}
	return false
}