package main

type Dictionary map[string]string

type DictionaryError string

const (
	NotFoundError       = DictionaryError("The term you're searching for does not exist.")
	ExistedTermError    = DictionaryError("It is not possible to add this term, because it already exist")
	nonExistedTermError = DictionaryError("It is not possible to update the term, because it does not exist")
)

func (e DictionaryError) Error() string {
	return string(e)
}

func (d Dictionary) Search(term string) (string, error) {
	definition, found := d[term]

	if !found {
		return "", NotFoundError
	}

	return definition, nil
}

func (d Dictionary) Add(key, term string) error {
	_, err := d.Search(key)

	switch err {
	case NotFoundError:
		d[key] = term
	case nil:
		return ExistedTermError
	default:
		return err
	}

	return nil
}

func (d Dictionary) Update(term, definition string) error {
	_, err := d.Search(term)

	switch err {
	case NotFoundError:
		return nonExistedTermError
	case nil:
		d[term] = definition
	default:
		return err
	}

	return nil
}

func (d Dictionary) Delete(term string) {
	delete(d, term)
}
