package utilities

func isAlphaNumeric(byt byte) bool {
	// Check a-z, A-Z, 0-9
	return (byt >= 'a' && byt <= 'z') || (byt >= 'A' && byt <= 'Z') || (byt >= '0' && byt <= '9')
}

func CountAlphaNumericChars(testString string) int64 {
	var alphaNumericCount int64 = 0
	for id := 0; id < len(testString); id++ {
		if isAlphaNumeric(testString[id]) {
			alphaNumericCount += 1
		}
	}
	return alphaNumericCount
}
