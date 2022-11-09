package netfield

import (
	"io"
)

// Parse return the field-name & field-body of a net-field — i.e., the type of fields found in HTTP headers & trailers, as well as SMTP headers.
func Parse(scanner io.RuneScanner) (name string, body string, err error) {

	if nil == scanner {
		err = errNilRuneScanner
		return
	}

	{
		name, err = parseName(scanner)
		if nil != err {
			return
		}
	}

	{
		err = consumeLinearSpacing(scanner)
		if nil != err {
			return
		}
	}

	{
		err = consumeColon(scanner)
		if nil != err {
			return
		}
	}

	{
		body, err = parseBody(scanner)
		if nil != err {
			return
		}
	}

	return
}
