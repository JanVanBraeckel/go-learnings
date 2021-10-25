package maps

type Dictionary map[string]string

var (
	ErrNotFound         = DictionaryErr("could not find the word you were looking for")
	ErrWordExists       = DictionaryErr("word already exists in given map")
	ErrWordDoesNotExist = DictionaryErr("word does not exist in map")
)

type DictionaryErr string

func (e DictionaryErr) Error() string {
	return string(e)
}

func (d Dictionary) Search(search string) (string, error) {
	definition, ok := d[search]

	if !ok {
		return "", ErrNotFound
	}

	return definition, nil
}

func (d Dictionary) Add(key, value string) error {
	_, err := d.Search(key)

	switch err {
	case ErrNotFound:
		d[key] = value
	case nil:
		return ErrWordExists
	default:
		return err
	}

	return nil
}

func (d Dictionary) Update(key, value string) error {
	_, err := d.Search(key)

	switch err {
	case ErrNotFound:
		return ErrWordDoesNotExist
	case nil:
		d[key] = value
	default:
		return nil
	}

	return nil
}

func (d Dictionary) Delete(key string) {
	delete(d, key)
}
