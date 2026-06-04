package main

import "testing"

func TestSearch(t *testing.T) {
	dictionary := Dictionary{"test": "This is a test"}

	t.Run("know term", func(t *testing.T) {
		result, _ := dictionary.Search("test")
		expected := "This is a test"

		stringCompare(t, result, expected)
	})

	t.Run("Unknown term", func(t *testing.T) {
		_, err := dictionary.Search("unknown")

		errorAssertion(t, err, NotFoundError)
	})
}

func TestAddTerm(t *testing.T) {
	dic := Dictionary{}

	dic.Add("added", "New definition added")

	result, err := dic.Search("added")
	expected := "New definition added"

	if err != nil {
		t.Fatal("It was not possible to find the added term:", err)
	}

	if result != expected {
		t.Errorf("result '%s', expected '%s'", result, expected)
	}
}

func stringCompare(t *testing.T, result, expected string) {
	t.Helper()

	if result != expected {
		t.Errorf("result '%s', expected '%s', data '%s'", result, expected, "test")
	}
}

func errorAssertion(t *testing.T, err, expected error) {
	t.Helper()

	if err != expected {
		t.Errorf("result '%s', expected '%s'", err, expected)
	}

}
