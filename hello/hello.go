package hello

const spanish = "Spanish"
const french = "French"
const helloPrefix = "Hello, "
const helloPrefixSpanish = "Hola, "
const helloPrefixFrench = "Bonjour, "

func Hello(name string, language string) string {
	if name == "" {
		name = "World"
	}
	return greetingPrefix(language) + name
}

func greetingPrefix(language string) (prefix string) {
	switch language {
	case spanish:
		prefix = helloPrefixSpanish
		return prefix
	case french:
		prefix = helloPrefixFrench
		return prefix
	default:
		prefix = helloPrefix
		return prefix
	}
}
