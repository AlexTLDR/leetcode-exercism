package space

type Planet string

func Earth(seconds float64) float64 {
	return seconds / 31557600
}

func Age(seconds float64, planet Planet) float64 {
	switch planet {
	case "Earth":
		return Earth(seconds)
	case "Mercury":
		return Earth(seconds) / 0.2408467
	case "Venus":
		return Earth(seconds) / 0.61519726
	case "Mars":
		return Earth(seconds) / 1.8808158
	case "Jupiter":
		return Earth(seconds) / 11.862615
	case "Saturn":
		return Earth(seconds) / 29.447498
	case "Uranus":
		return Earth(seconds) / 84.016846
	case "Neptune":
		return Earth(seconds) / 164.79132
	}

	return -1
}
