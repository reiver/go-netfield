package netfield

import (
	"fmt"
	"io"
	"strings"
)

// parseBody returns the body of a net-field.
func parseBody(scanner io.RuneScanner) (string, error) {

	if nil == scanner {
		return "", errNilRuneScanner
	}

	var body string
	{
		var storage strings.Builder

		for {
			var r1 rune
			{
				r, _, err := scanner.ReadRune()
				if io.EOF == err {
		/////////////////////// BREAK
					break
				}
				if nil != err {
					return storage.String(), fmt.Errorf("problem reading next unicode character (r1), when trying to parse field-body: %w", err)
				}
				if runeError == r {
					return storage.String(), fmt.Errorf("problem reading next unicode character (r1), when trying to parse field-body: %w", errRuneError)
				}

				r1 = r
			}

			if cr != r1 {
				storage.WriteRune(r1)
		/////////////// CONTINUE
				continue
			}

			var r2 rune
			{
				r, _, err := scanner.ReadRune()
				if io.EOF == err {
					storage.WriteRune(r1)
		/////////////////////// BREAK
					break
				}
				if nil != err {
					return storage.String(), fmt.Errorf("problem reading next unicode character (r2), when trying to parse field-body: %w", err)
				}
				if runeError == r {
					return storage.String(), fmt.Errorf("problem reading next unicode character (r2), when trying to parse field-body: %w", errRuneError)
				}

				r2 = r
			}

			if lf != r2 {
				storage.WriteRune(r1)
				storage.WriteRune(r2)
		/////////////// CONTINUE
				continue
			}

			var r3 rune
			{
				r, _, err := scanner.ReadRune()
				if io.EOF == err {
		/////////////////////// BREAK
					break
				}
				if nil != err {
					return storage.String(), fmt.Errorf("problem reading next unicode character (r3), when trying to parse field-body: %w", err)
					}
				if runeError == r {
					return storage.String(), fmt.Errorf("problem reading next unicode character (r3), when trying to parse field-body: %w", errRuneError)
				}

				r3 = r
			}

			if isLinearSpacing(r3) {
				storage.WriteRune(r1)
				storage.WriteRune(r2)
				storage.WriteRune(r3)
		/////////////// CONTINUE
				continue
			}

			err := scanner.UnreadRune()
			if nil != err {
				return storage.String(), fmt.Errorf("problem un-reading unicode character %U, when trying to parse field-body: %w", r3, err)
			}

		/////// BREAK
			break
		}

		body = storage.String()
	}

	return body, nil
}
