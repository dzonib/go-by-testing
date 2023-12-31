package intro

import "fmt"

const (
	spanish = "Spanish"
	french  = "French"

	englishPrefix = "Hello"
	spanishPrefix = "Hola"
	frenchPrefix  = "YO"
)

func Hello(name string, language string) string {

	if name == "" {
		name = "Sir/Madam"
	}

	return greetingPrefix(language) + " " + name
}

func greetingPrefix(language string) (prefix string) {
	switch language {
	case spanish:
		prefix = spanishPrefix
	case french:
		prefix = frenchPrefix
	default:
		prefix = englishPrefix
	}
	return
}

func main() {
	fmt.Println(Hello("Mr", "Spanish"))
}
