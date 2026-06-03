package main

import "testing"

func TestSearch(t *testing.T) {
	dictionary := map[string]string{"test": "This is a test"}

	result := Search(dictionary, "test")
	expected := "This is a test"

	stringCompare(t, result, expected)
}

func stringCompare(t *testing.T, result, expected string) {
	t.Helper()

	if result != expected {
		t.Errorf("result '%s', expected '%s', data '%s'", result, expected, "test")
	}
}
