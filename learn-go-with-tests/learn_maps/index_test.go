package main

import "testing"

func TestSearch(t *testing.T) {
	dictionary := Dictionary{"test": "This is a test"}

	result := dictionary.Search("test")
	expected := "This is a test"

	stringCompare(t, result, expected)
}

func stringCompare(t *testing.T, result, expected string) {
	t.Helper()

	if result != expected {
		t.Errorf("result '%s', expected '%s', data '%s'", result, expected, "test")
	}
}
