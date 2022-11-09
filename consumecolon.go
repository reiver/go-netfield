package netfield

import (
	"fmt"
	"io"
)

// consumeColon consumes a single colon character (i.e., ":").
func consumeColon(scanner io.RuneScanner) error {

	if nil == scanner {
		return errNilRuneScanner
	}

	{
		r, _, err := scanner.ReadRune()
		if nil != err {
			return fmt.Errorf("problem reading next unicode character, when trying to consume colon character(':'): %w", err)
		}
		if runeError == r {
			return fmt.Errorf("problem reading next unicode character, when trying to consume colon character(':'): %w", errRuneError)
		}

		if ':' != r {
			err := scanner.UnreadRune()
			if nil != err {
				return fmt.Errorf("problem un-reading unicode character %U, when trying to consume colon character(':'): %w", r, err)
			}

			return fmt.Errorf("expected a colon character (':') but actually got a %q character: %w", string(r), errColonNotFound)
		}
	}

	return nil
}
