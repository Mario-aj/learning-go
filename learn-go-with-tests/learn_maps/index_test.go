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
	t.Run("new term", func(t *testing.T) {
		dic := Dictionary{}
		term := "test"
		definition := "New definition added"

		err := dic.Add(term, definition)

		errorAssertion(t, err, nil)
		compareDefinition(t, dic, term, definition)
	})

	t.Run("existed term", func(t *testing.T) {
		term := "test"
		definition := "This is a test"

		dic := Dictionary{term: definition}

		err := dic.Add(term, definition)

		errorAssertion(t, err, ExistedTermError)
		compareDefinition(t, dic, term, definition)
	})
}

func TestUpdate(t *testing.T) {
	t.Run("exited term", func(t *testing.T) {
		term := "test"
		definition := "This is a test"
		newDefinition := "New definition"

		dic := Dictionary{term: definition}

		err := dic.Update(term, newDefinition)

		errorAssertion(t, err, nil)
		compareDefinition(t, dic, term, newDefinition)
	})

	t.Run("new term", func(t *testing.T) {
		term := "test"
		definition := "This is a test"
		dic := Dictionary{}

		err := dic.Update(term, definition)

		errorAssertion(t, err, nonExistedTermError)
	})
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

func compareDefinition(t *testing.T, dic Dictionary, term, definition string) {
	t.Helper()

	result, err := dic.Search(term)

	if err != nil {
		t.Fatal("Should find the added term:", err)
	}

	if result != definition {
		t.Errorf("result '%s', expected '%s'", result, definition)
	}
}
