package netfield

import (
	"github.com/reiver/go-utf8"

	"errors"
	"io"
	"strings"

	"testing"
)

func TestParseBody(t *testing.T) {

	tests := []struct{
		Data string
		Expected string
	}{
		{
			Data:
				"apple banana cherry",
			Expected:
				"apple banana cherry",
		},
		{
			Data:
				"apple banana cherry"+"\r\n",
			Expected:
				"apple banana cherry",
		},



		{
			Data:
				"apple banana cherry"+"\r\n"+
				"something",
			Expected:
				"apple banana cherry",
		},
		{
			Data:
				"apple banana cherry"+"\r\n"+
				"something: here",
			Expected:
				"apple banana cherry",
		},
		{
			Data:
				"apple banana cherry"+"\r\n"+
				"something: here"    +"\r\n",
			Expected:
				"apple banana cherry",
		},



		{
			Data:
				"apple banana cherry"+"\r\n"+
				" ",
			Expected:
				"apple banana cherry\r\n ",
		},
		{
			Data:
				"apple banana cherry"+"\r\n"+
				"\t",
			Expected:
				"apple banana cherry\r\n\t",
		},



		{
			Data:
				"apple banana cherry"+"\r\n"+
				" "+"\r\n",
			Expected:
				"apple banana cherry\r\n ",
		},
		{
			Data:
				"apple banana cherry"+"\r\n"+
				"\t"+"\r\n",
			Expected:
				"apple banana cherry\r\n\t",
		},



		{
			Data:
				"apple banana cherry"+"\r\n"+
				" date",
			Expected:
				"apple banana cherry\r\n date",
		},
		{
			Data:
				"apple banana cherry"+"\r\n"+
				"\tdate",
			Expected:
				"apple banana cherry\r\n\tdate",
		},



		{
			Data:
				"apple banana cherry"+"\r\n"+
				" date"+"\r\n",
			Expected:
				"apple banana cherry\r\n date",
		},
		{
			Data:
				"apple banana cherry"+"\r\n"+
				"\tdate"+"\r\n",
			Expected:
				"apple banana cherry\r\n\tdate",
		},



		{
			Data:
				" ONCE"  +"\r\n"+
				" TWICE" +"\r\n"+
				" THRICE"+"\r\n"+
				" FOURCE",
			Expected:
				" ONCE\r\n TWICE\r\n THRICE\r\n FOURCE",
		},
		{
			Data:
				" ONCE"  +"\r\n"+
				" TWICE" +"\r\n"+
				" THRICE"+"\r\n"+
				" FOURCE"+"\r\n",
			Expected:
				" ONCE\r\n TWICE\r\n THRICE\r\n FOURCE",
		},
		{
			Data:
				" ONCE"  +"\r\n"+
				" TWICE" +"\r\n"+
				" THRICE"+"\r\n"+
				" FOURCE"+"\r\n"+
				"STOP",
			Expected:
				" ONCE\r\n TWICE\r\n THRICE\r\n FOURCE",
		},
	}

	for testNumber, test := range tests {

			var reader io.Reader = strings.NewReader(test.Data)
			wrapped := utf8.RuneScannerWrap(reader)
			var runescanner io.RuneScanner = &wrapped

			actual, err := parseBody(runescanner)
			if nil != err && !errors.Is(err, io.EOF) {
				t.Errorf("For test #%d, did not expect an error but actually got one.", testNumber)
				t.Logf("ERROR: (%T) %q", err, err)
				t.Logf("DATA: %q", test.Data)
				t.Logf("ExPECTED-FIELD-BODY: %q", test.Expected)
				t.Logf("ACTUAL-FIELD_BODY:   %q", actual)
				continue
			}

			var expected string = test.Expected

			if expected != actual {
				t.Errorf("For test #%d, the actual field-body is not what was expected.", testNumber)
				t.Logf("ExPECTED: %q", expected)
				t.Logf("ACTUAL:   %q", actual)
				t.Logf("DATA: %q", test.Data)
				continue
			}
	}
}
