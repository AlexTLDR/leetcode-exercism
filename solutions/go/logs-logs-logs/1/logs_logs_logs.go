package logs

// Application identifies the application emitting the given log.
func Application(log string) string {
	for _, char := range log {
		switch {
		case char == '❗':
			return "recommendation"
		case char == '🔍':
			return "search"
		case char == '☀':
			return "weather"
		}
	}
	return "default"
}

// Replace replaces all occurrences of old with new, returning the modified log
// to the caller.
func Replace(log string, oldRune, newRune rune) string {
	logFixed := ""
	for _, char := range log {
		switch {
		case char == oldRune:
			char = newRune
		}
		logFixed += string(char)
	}
	return logFixed
}

// WithinLimit determines whether or not the number of characters in log is
// within the limit.
func WithinLimit(log string, limit int) bool {
	for i := range log {
		if i+1 > limit {
			return false
		}
	}
	return true
}
