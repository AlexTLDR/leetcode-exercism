package rotationalcipher

func RotationalCipher(plain string, shiftKey int) string {
	capitalAlphabet := "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	lowerAlphabet := "abcdefghijklmnopqrstuvwxyz"
	shiftedCapitalAlphabet := make(map[string]string)
	shiftedLowerAlphabet := make(map[string]string)
	for i, letter := range capitalAlphabet {
		shiftedCapitalAlphabet[string(letter)] = string(capitalAlphabet[(i+shiftKey)%26])
	}

	for i, letter := range lowerAlphabet {
		shiftedLowerAlphabet[string(letter)] = string(lowerAlphabet[(i+shiftKey)%26])
	}

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
