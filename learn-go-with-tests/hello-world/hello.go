package main

import "fmt"

const prefixHello = "Hello, "
const spanish = "Spanish"
const spanishPrefix = "Hola, "
const frenchPrefix = "Bonjour, " 
const french = "French"

func Hello(name string, language string) string {

	if name == "" {
		name = "world!"
	}

	if language == spanish {
		return spanishPrefix + name
	}

	if language == french {
		return frenchPrefix + name	
	}

	return prefixHello + name
}

func main() {
	fmt.Println(Hello("world!", ""))
}