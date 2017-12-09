package gcitem

func Parse(value string) (Item, bool) {
	i, ok := ParseInteger(value)

	if ok {
		return i, true
	}

	return nil, false
}
