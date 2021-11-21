package mydict

import "errors"

// Dictionary type
type Dictionary map[string]string

var errNotFound = errors.New("not found.")

func (d Dictionary) Search(word string) (string, error) {
	value, exists := d[word]
	if exists {
		return value, nil
	}
	return "", errNotFound
}
