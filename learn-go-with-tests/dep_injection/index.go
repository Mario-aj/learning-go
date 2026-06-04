package main

import (
	"fmt"
	"io"
	"os"
)

func Greeting(writer io.Writer, name string) {
	fmt.Fprintf(writer, "Hello, %s", name)
}

func main() {
	Greeting(os.Stdout, "Mario")
}
