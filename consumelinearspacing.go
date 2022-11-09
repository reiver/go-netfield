package netfield

import (
	"fmt"
	"io"
)

// consumeLinearSpacing consumes zero to many contiguous linear-spacing characters.
func consumeLinearSpacing(scanner io.RuneScanner) error {

	if nil == scanner {
		return errNilRuneScanner
	}


	for {
		r, _, err := scanner.ReadRune()
		if nil != err {
			return fmt.Errorf("problem reading next unicode character: %w", err)
		}
		if runeError == r {
			return fmt.Errorf("problem reading next unicode character: %w", errRuneError)
		}

		if !isLinearSpacing(r) {
			err := scanner.UnreadRune()
			if nil != err {
				return fmt.Errorf("problem un-reading unicode character %U: %w", r, err)
			}

	/////////////// BREAK
			break
		}
	}

	return nil
}
