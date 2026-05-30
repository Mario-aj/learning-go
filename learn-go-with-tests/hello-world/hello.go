package main

import (
	"fmt"
)

const prefixHello = "Hello, "
const spanish = "Spanish"
const spanishPrefix = "Hola, "
const frenchPrefix = "Bonjour, " 
const french = "French"
const angolan = "Angolan"
const angolanPrefix = "É como, "

func Hello(name string, language string) string {

	if name == "" {
		name = "world!"
	}
	
	return greetingPrefix(language) + name
}

func greetingPrefix(language string) (prefix string) {
	switch language {
	case french:
		prefix = frenchPrefix
	case spanish:
		prefix = spanishPrefix
	case angolan:
		prefix = angolanPrefix

	default:
		prefix = prefixHello	
	}

	return
}

func main() {
	fmt.Println(Hello("world!", ""))
}