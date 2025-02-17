package main

const (
	englishHelloPrefix = "Hello, "
	spanish            = "Spanish"
	spanishHelloPrefix = "Hola, "
	german             = "German"
	germanHelloPrefix  = "Hallo, "
)

// uppercase - public functions
func Hello(name string, lang string) string {
	if name == "" {
		name = "World"
	}

	return greetingPrefix(lang) + name
}

// lowercase private functions
func greetingPrefix(lang string) (prefix string) {
	switch lang {
	case spanish:
		prefix = spanishHelloPrefix

	case german:
		prefix = germanHelloPrefix

	default:
		prefix = englishHelloPrefix
	}

	return
}
