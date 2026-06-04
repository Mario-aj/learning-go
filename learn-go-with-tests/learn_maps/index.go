package main

import "errors"

type Dictionary map[string]string

var NotFoundError = errors.New("The term you're searching for does not exist.")

func (d Dictionary) Search(term string) (string, error) {
	definition, found := d[term]

	if !found {
		return "", NotFoundError
	}

	return definition, nil
}

func (d Dictionary) Add(key, term string) {
	d[key] = term
}
