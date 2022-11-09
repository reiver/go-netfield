package netfield

import (
	"fmt"
	"io"
	"strings"
)

// parseName returns the name of a net-field.
func parseName(scanner io.RuneScanner) (string, error) {

	if nil == scanner {
		return "", errNilRuneScanner
	}

	var name string
	{
		var storage strings.Builder

		Loop: for {
			r, _, err := scanner.ReadRune()
			if nil != err {
				return storage.String(), fmt.Errorf("problem reading next unicode character, when trying to parse field-name: %w", err)
			}
			if runeError == r {
				return storage.String(), fmt.Errorf("problem reading next unicode character, when trying to parse field-name: %w", errRuneError)
			}

			if ':' == r || isLinearSpacing(r) {
				err := scanner.UnreadRune()
				if nil != err {
					return storage.String(), fmt.Errorf("problem un-reading unicode character %U, when trying to parse field-name: %w", r, err)
				}
		/////////////// BREAK
				break Loop
			}
			storage.WriteRune(r)
		}

		name = storage.String()
	}

	return name, nil
}
