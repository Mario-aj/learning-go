package main

import (
	"fmt"
	"io"
	"net/http"
)

func Greeting(writer io.Writer, name string) {
	fmt.Fprintf(writer, "Hello, %s", name)
}

func HandleMyGreeting(w http.ResponseWriter, r *http.Request) {
	Greeting(w, "World!")
}

func main() {
	err := http.ListenAndServe(":5000", http.HandlerFunc(HandleMyGreeting))

	if err != nil {
		fmt.Println(err)
	}
}
