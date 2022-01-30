package iteration

// Repeat returns a string containing input character n times where n is repeatCount
func Repeat(character string, repeatCount int) string {
	var returnVal string
	for i := 0; i < repeatCount; i++ {
		returnVal += character
	}
	return returnVal
}
