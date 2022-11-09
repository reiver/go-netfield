package netfield

import (
	"github.com/reiver/go-utf8"

	"io"
	"strings"

	"testing"
)

func TestParseName(t *testing.T) {

	tests := []struct{
		Data string
		Expected string
	}{
		{
			Data:     "name: apple banana cherry",
			Expected: "name",
		},

		{
			Data:     "name : apple banana cherry",
			Expected: "name",
		},
		{
			Data:     "name  : apple banana cherry",
			Expected: "name",
		},
		{
			Data:     "name   : apple banana cherry",
			Expected: "name",
		},

		{
			Data:     "name\t: apple banana cherry",
			Expected: "name",
		},
		{
			Data:     "name\t\t: apple banana cherry",
			Expected: "name",
		},
		{
			Data:     "name\t\t\t: apple banana cherry",
			Expected: "name",
		},



		{
			Data:     "Apple-Banana-Cherry: once, twice, thrice, fource",
			Expected: "Apple-Banana-Cherry",
		},

		{
			Data:     "Apple-Banana-Cherry : once, twice, thrice, fource",
			Expected: "Apple-Banana-Cherry",
		},
		{
			Data:     "Apple-Banana-Cherry  : once, twice, thrice, fource",
			Expected: "Apple-Banana-Cherry",
		},
		{
			Data:     "Apple-Banana-Cherry   : once, twice, thrice, fource",
			Expected: "Apple-Banana-Cherry",
		},

		{
			Data:     "Apple-Banana-Cherry\t: once, twice, thrice, fource",
			Expected: "Apple-Banana-Cherry",
		},
		{
			Data:     "Apple-Banana-Cherry\t\t: once, twice, thrice, fource",
			Expected: "Apple-Banana-Cherry",
		},
		{
			Data:     "Apple-Banana-Cherry\t\t\t: once, twice, thrice, fource",
			Expected: "Apple-Banana-Cherry",
		},
	}

	for testNumber, test := range tests {

			var reader io.Reader = strings.NewReader(test.Data)
			wrapped := utf8.RuneScannerWrap(reader)
			var runescanner io.RuneScanner = &wrapped

			actual, err := parseName(runescanner)
			if nil != err {
				t.Errorf("For test #%d, did not expect an error but actually got one.", testNumber)
				t.Logf("ERROR: (%T) %q", err, err)
				t.Logf("DATA: %q", test.Data)
				continue
			}

			var expected string = test.Expected

			if expected != actual {
				t.Errorf("For test #%d, the actual field-name is not what was expected.", testNumber)
				t.Logf("ExPECTED: %q", expected)
				t.Logf("ACTUAL:   %q", actual)
				t.Logf("DATA: %q", test.Data)
				continue
			}
	}
}
