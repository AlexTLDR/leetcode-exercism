package rotationalcipher

func RotationalCipher(plain string, shiftKey int) string {
	shiftedCapitalAlphabet := make(map[string]string)
	shiftedLowerAlphabet := make(map[string]string)
	shiftedCapitalAlphabet = shift("ABCDEFGHIJKLMNOPQRSTUVWXYZ", shiftKey)
	shiftedLowerAlphabet = shift("abcdefghijklmnopqrstuvwxyz", shiftKey)

	encriptedMessage := ""
	for _, letter := range plain {
		switch {
		case letter <= 65 || (letter > 90 && letter < 97) || (letter > 122):
			encriptedMessage += string(letter)
		case letter >= 'A' && letter <= 'Z':
			encriptedMessage += shiftedCapitalAlphabet[string(letter)]
		case letter >= 'a' && letter <= 'z':
			encriptedMessage += shiftedLowerAlphabet[string(letter)]
		}
	}
	return encriptedMessage
}

func shift(alphabet string, shiftKey int) map[string]string {
	shifted := make(map[string]string)
	for i, letter := range alphabet {
		shifted[string(letter)] = string(alphabet[(i+shiftKey)%26])
	}
	return shifted
}