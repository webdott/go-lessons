package main

func Repeat(char string, repeatVal int) string {
	retVal := ""

	for i := 0; i < repeatVal; i++ {
		retVal += char
	}

	return retVal
}
