package main

import (
	"fmt"
)

const prefixHello = "Hello, "
const spanish = "Spanish"
const spanishPrefix = "Hola, "
const frenchPrefix = "Bonjour, " 
const french = "French"

func Hello(name string, language string) string {

	if name == "" {
		name = "world!"
	}

	prefix := prefixHello

	switch language {
		case french: 
			prefix = frenchPrefix	
		case spanish:
			prefix = spanishPrefix
	}

	return prefix + name
}

func main() {
	fmt.Println(Hello("world!", ""))
}