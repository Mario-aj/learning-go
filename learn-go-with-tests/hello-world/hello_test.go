package main

import (
	"testing"
)

func TestHello(t *testing.T) {
	verifyCorrectMessage := func(t *testing.T, result, expected string) {
		t.Helper()
		
		if result != expected {
			t.Errorf("result '%s', expected '%s'", result, expected)
		}
	}

	t.Run("Say hello to people", func(t*testing.T){
		result := Hello("Chris", "")
		expected := "Hello, Chris"


		verifyCorrectMessage(t, result, expected)
	})

	t.Run("Say 'Hello, world!' when a empty string will be passed", func(t *testing.T) {
		result:=Hello("", "")
		expected := "Hello, world!"

		verifyCorrectMessage(t, result, expected)
	})

	t.Run("In spanish", func(t *testing.T) {
		result := Hello("Mario", "Spanish")
		expected := "Hola, Mario"

		verifyCorrectMessage(t, result, expected)
	})

	t.Run("Should greet in french", func(t *testing.T) {
		result := Hello("Mario", "French")
		expected := "Bonjour, Mario"
		
		verifyCorrectMessage(t, result, expected)
	})
}