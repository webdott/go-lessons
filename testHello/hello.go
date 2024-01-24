package main

import "fmt"

const (
	spanish = "Spanish"
	french  = "French"

	englishPrefix string = "Hello, "
	spanishPrefix string = "Hola, "
	frenchPrefix  string = "Bonjour, "
)

func Hello(name string, language string) string {
	if name == "" {
		name = "World"
	}

	return getGreetPrefix(language) + name
}

func getGreetPrefix(language string) (prefixToUse string) {
	switch language {
	case spanish:
		prefixToUse = spanishPrefix
	case french:
		prefixToUse = frenchPrefix
	default:
		prefixToUse = englishPrefix
	}
	return
}

func main() {
	fmt.Println(Hello("Uche", "Spanish"))
}
