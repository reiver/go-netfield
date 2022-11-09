package netfield

import (
	"github.com/reiver/go-utf8"

	"io"
	"strings"

	"testing"
)

func TestConsumeColon(t *testing.T) {

	tests := []struct{
		Data string
	}{
		{
			Data: ":",
		},
		{
			Data: ": ",
		},
		{
			Data: ": once twice thrice fource",
		},
		{
			Data: ": once twice thrice fource\r\n",
		},
		{
			Data: ":\t",
		},
		{
			Data: ":\r",
		},
		{
			Data: ":\r\n",
		},
		{
			Data: ":\n",
		},
		{
			Data: "::",
		},
		{
			Data: ":something",
		},
		{
			Data: ":something\r\n",
		},
	}

	TestLoop: for testNumber, test := range tests {

		var reader io.Reader = strings.NewReader(test.Data)
		wrapped := utf8.RuneScannerWrap(reader)
		var runescanner io.RuneScanner = &wrapped

		{
			err := consumeColon(runescanner)
			if nil != err {
				t.Errorf("For test #%d, did not expect to get an error, but actually got one", testNumber)
				t.Logf("ERROR: (%T) %q", err, err)
				t.Logf("DATA: %q", test.Data)
	/////////////////////// CONTINUE
				continue TestLoop
			}
		}

		var remaining string
		{
			var storage strings.Builder

			for {
				r, _, err := runescanner.ReadRune()
				if io.EOF == err {
			/////////////// BREAK
					break
				}
				if nil != err {
				t.Errorf("For test #%d, did not expect to get an error, but actually got one", testNumber)
				t.Logf("ERROR: (%T) %q", err, err)
				t.Logf("DATA: %q", test.Data)
	/////////////////////////////// BREAK
					break TestLoop
				}

				storage.WriteRune(r)
			}

			remaining = storage.String()
		}

		{
			var expected string = test.Data[len(":"):]
			var actual   string = remaining

			if expected != actual {
				t.Errorf("For test #%d, the actual value is not what was expected.", testNumber)
				t.Logf("EXPECTED: %q", expected)
				t.Logf("ACTUAL:   %q", actual)
	/////////////////////// CONTINUE
				continue TestLoop
			}
		}
	}
}
