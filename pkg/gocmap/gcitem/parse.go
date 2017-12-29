package gcitem

import "fmt"

func Parse(value string) (Item, error) {
	i, err := ParseInteger(value)
	if err != nil {
		return nil, fmt.Errorf("could not parse value %v: %v", value, err)
	}

	return i, nil
}
