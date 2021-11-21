package mydict

import "errors"

// Dictionary type
type Dictionary map[string]string

var (
	errNotFound      = errors.New("not found")
	errAlreadyExists = errors.New("already exists")
)

func (d Dictionary) Search(word string) (string, error) {
	value, exists := d[word]
	if exists {
		return value, nil
	}
	return "", errNotFound
}

// Add a word to the dictionary
func (d Dictionary) Add(word, def string) error {
	switch _, err := d.Search(word); err {
	case errNotFound:
		d[word] = def
	case nil:
		return errAlreadyExists
	}
	return nil
}

func (d Dictionary) Update(word, def string) error {
	_, err := d.Search(word)
	switch err {
	case errNotFound:
		return err
	case nil:
		d[word] = def
	}
	return nil
}

func (d Dictionary) Delete(word string) error {
	_, err := d.Search(word)
	switch err {
	case errNotFound:
		return err
	case nil:
		delete(d, word)
	}
	return nil
}
