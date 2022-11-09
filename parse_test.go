package netfield_test

import (
	"github.com/reiver/go-netfield"

	"github.com/reiver/go-utf8"

	"io"
	"strings"

	"testing"
)

func TestParse(t *testing.T) {

	tests := []struct{
		Data string
		ExpectedFieldName string
		ExpectedFieldBody string
	}{
		{
			Data:              "the-name:the-body",
			ExpectedFieldName: "the-name",
			ExpectedFieldBody:          "the-body",
		},
		{
			Data:              "the-name: the-body",
			ExpectedFieldName: "the-name",
			ExpectedFieldBody:          " the-body",
		},
		{
			Data:              "the-name:  the-body",
			ExpectedFieldName: "the-name",
			ExpectedFieldBody:          "  the-body",
		},
		{
			Data:              "the-name :the-body",
			ExpectedFieldName: "the-name",
			ExpectedFieldBody:           "the-body",
		},
		{
			Data:              "the-name  :the-body",
			ExpectedFieldName: "the-name",
			ExpectedFieldBody:            "the-body",
		},
		{
			Data:              "the-name : the-body",
			ExpectedFieldName: "the-name",
			ExpectedFieldBody:           " the-body",
		},
		{
			Data:              "the-name :  the-body",
			ExpectedFieldName: "the-name",
			ExpectedFieldBody:           "  the-body",
		},
		{
			Data:              "the-name  : the-body",
			ExpectedFieldName: "the-name",
			ExpectedFieldBody:            " the-body",
		},
		{
			Data:              "the-name  :  the-body",
			ExpectedFieldName: "the-name",
			ExpectedFieldBody:            "  the-body",
		},



		{
			Data:              "the-name:the-body"+"\r",
			ExpectedFieldName: "the-name",
			ExpectedFieldBody:          "the-body\r",
		},
		{
			Data:              "the-name: the-body"+"\r",
			ExpectedFieldName: "the-name",
			ExpectedFieldBody:          " the-body\r",
		},
		{
			Data:              "the-name:  the-body"+"\r",
			ExpectedFieldName: "the-name",
			ExpectedFieldBody:          "  the-body\r",
		},
		{
			Data:              "the-name :the-body"+"\r",
			ExpectedFieldName: "the-name",
			ExpectedFieldBody:           "the-body\r",
		},
		{
			Data:              "the-name  :the-body"+"\r",
			ExpectedFieldName: "the-name",
			ExpectedFieldBody:           "the-body\r",
		},
		{
			Data:              "the-name : the-body"+"\r",
			ExpectedFieldName: "the-name",
			ExpectedFieldBody:           " the-body\r",
		},
		{
			Data:              "the-name :  the-body"+"\r",
			ExpectedFieldName: "the-name",
			ExpectedFieldBody:           "  the-body\r",
		},
		{
			Data:              "the-name  : the-body"+"\r",
			ExpectedFieldName: "the-name",
			ExpectedFieldBody:            " the-body\r",
		},
		{
			Data:              "the-name  :  the-body"+"\r",
			ExpectedFieldName: "the-name",
			ExpectedFieldBody:            "  the-body\r",
		},



		{
			Data:              "the-name:the-body"+"\n",
			ExpectedFieldName: "the-name",
			ExpectedFieldBody:          "the-body\n",
		},
		{
			Data:              "the-name: the-body"+"\n",
			ExpectedFieldName: "the-name",
			ExpectedFieldBody:          " the-body\n",
		},
		{
			Data:              "the-name:  the-body"+"\n",
			ExpectedFieldName: "the-name",
			ExpectedFieldBody:          "  the-body\n",
		},
		{
			Data:              "the-name :the-body"+"\n",
			ExpectedFieldName: "the-name",
			ExpectedFieldBody:           "the-body\n",
		},
		{
			Data:              "the-name  :the-body"+"\n",
			ExpectedFieldName: "the-name",
			ExpectedFieldBody:           "the-body\n",
		},
		{
			Data:              "the-name : the-body"+"\n",
			ExpectedFieldName: "the-name",
			ExpectedFieldBody:           " the-body\n",
		},
		{
			Data:              "the-name :  the-body"+"\n",
			ExpectedFieldName: "the-name",
			ExpectedFieldBody:           "  the-body\n",
		},
		{
			Data:              "the-name  : the-body"+"\n",
			ExpectedFieldName: "the-name",
			ExpectedFieldBody:            " the-body\n",
		},
		{
			Data:              "the-name  :  the-body"+"\n",
			ExpectedFieldName: "the-name",
			ExpectedFieldBody:            "  the-body\n",
		},



		{
			Data:              "the-name:the-body"+"\r\n",
			ExpectedFieldName: "the-name",
			ExpectedFieldBody:          "the-body",
		},
		{
			Data:              "the-name: the-body"+"\r\n",
			ExpectedFieldName: "the-name",
			ExpectedFieldBody:          " the-body",
		},
		{
			Data:              "the-name:  the-body"+"\r\n",
			ExpectedFieldName: "the-name",
			ExpectedFieldBody:          "  the-body",
		},
		{
			Data:              "the-name :the-body"+"\r\n",
			ExpectedFieldName: "the-name",
			ExpectedFieldBody:           "the-body",
		},
		{
			Data:              "the-name  :the-body"+"\r\n",
			ExpectedFieldName: "the-name",
			ExpectedFieldBody:           "the-body",
		},
		{
			Data:              "the-name : the-body"+"\r\n",
			ExpectedFieldName: "the-name",
			ExpectedFieldBody:           " the-body",
		},
		{
			Data:              "the-name :  the-body"+"\r\n",
			ExpectedFieldName: "the-name",
			ExpectedFieldBody:           "  the-body",
		},
		{
			Data:              "the-name  : the-body"+"\r\n",
			ExpectedFieldName: "the-name",
			ExpectedFieldBody:            " the-body",
		},
		{
			Data:              "the-name  :  the-body"+"\r\n",
			ExpectedFieldName: "the-name",
			ExpectedFieldBody:            "  the-body",
		},



		{
			Data:              "the-name:the-body"+"\r"+
			                   "something: nothing",
			ExpectedFieldName: "the-name",
			ExpectedFieldBody:          "the-body\rsomething: nothing",
		},
		{
			Data:              "the-name: the-body"+"\r"+
			                   "something: nothing",
			ExpectedFieldName: "the-name",
			ExpectedFieldBody:          " the-body\rsomething: nothing",
		},
		{
			Data:              "the-name:  the-body"+"\r"+
			                   "something: nothing",
			ExpectedFieldName: "the-name",
			ExpectedFieldBody:          "  the-body\rsomething: nothing",
		},
		{
			Data:              "the-name :the-body"+"\r"+
			                   "something: nothing",
			ExpectedFieldName: "the-name",
			ExpectedFieldBody:           "the-body\rsomething: nothing",
		},
		{
			Data:              "the-name  :the-body"+"\r"+
			                   "something: nothing",
			ExpectedFieldName: "the-name",
			ExpectedFieldBody:            "the-body\rsomething: nothing",
		},
		{
			Data:              "the-name : the-body"+"\r"+
			                   "something: nothing",
			ExpectedFieldName: "the-name",
			ExpectedFieldBody:           " the-body\rsomething: nothing",
		},
		{
			Data:              "the-name :  the-body"+"\r"+
			                   "something: nothing",
			ExpectedFieldName: "the-name",
			ExpectedFieldBody:           "  the-body\rsomething: nothing",
		},
		{
			Data:              "the-name  : the-body"+"\r"+
			                   "something: nothing",
			ExpectedFieldName: "the-name",
			ExpectedFieldBody:            " the-body\rsomething: nothing",
		},
		{
			Data:              "the-name  :  the-body"+"\r"+
			                   "something: nothing",
			ExpectedFieldName: "the-name",
			ExpectedFieldBody:            "  the-body\rsomething: nothing",
		},



		{
			Data:              "the-name:the-body"+"\n"+
			                   "something: nothing",
			ExpectedFieldName: "the-name",
			ExpectedFieldBody:          "the-body\nsomething: nothing",
		},
		{
			Data:              "the-name: the-body"+"\n"+
			                   "something: nothing",
			ExpectedFieldName: "the-name",
			ExpectedFieldBody:          " the-body\nsomething: nothing",
		},
		{
			Data:              "the-name:  the-body"+"\n"+
			                   "something: nothing",
			ExpectedFieldName: "the-name",
			ExpectedFieldBody:          "  the-body\nsomething: nothing",
		},
		{
			Data:              "the-name :the-body"+"\n"+
			                   "something: nothing",
			ExpectedFieldName: "the-name",
			ExpectedFieldBody:           "the-body\nsomething: nothing",
		},
		{
			Data:              "the-name  :the-body"+"\n"+
			                   "something: nothing",
			ExpectedFieldName: "the-name",
			ExpectedFieldBody:            "the-body\nsomething: nothing",
		},
		{
			Data:              "the-name : the-body"+"\n"+
			                   "something: nothing",
			ExpectedFieldName: "the-name",
			ExpectedFieldBody:           " the-body\nsomething: nothing",
		},
		{
			Data:              "the-name :  the-body"+"\n"+
			                   "something: nothing",
			ExpectedFieldName: "the-name",
			ExpectedFieldBody:           "  the-body\nsomething: nothing",
		},
		{
			Data:              "the-name  : the-body"+"\n"+
			                   "something: nothing",
			ExpectedFieldName: "the-name",
			ExpectedFieldBody:            " the-body\nsomething: nothing",
		},
		{
			Data:              "the-name  :  the-body"+"\n"+
			                   "something: nothing",
			ExpectedFieldName: "the-name",
			ExpectedFieldBody:            "  the-body\nsomething: nothing",
		},



		{
			Data:              "the-name:the-body"+"\r\n"+
			                   "something: nothing",
			ExpectedFieldName: "the-name",
			ExpectedFieldBody:          "the-body",
		},
		{
			Data:              "the-name: the-body"+"\r\n"+
			                   "something: nothing",
			ExpectedFieldName: "the-name",
			ExpectedFieldBody:          " the-body",
		},
		{
			Data:              "the-name:  the-body"+"\r\n"+
			                   "something: nothing",
			ExpectedFieldName: "the-name",
			ExpectedFieldBody:          "  the-body",
		},
		{
			Data:              "the-name :the-body"+"\r\n"+
			                   "something: nothing",
			ExpectedFieldName: "the-name",
			ExpectedFieldBody:           "the-body",
		},
		{
			Data:              "the-name  :the-body"+"\r\n"+
			                   "something: nothing",
			ExpectedFieldName: "the-name",
			ExpectedFieldBody:            "the-body",
		},
		{
			Data:              "the-name : the-body"+"\r\n"+
			                   "something: nothing",
			ExpectedFieldName: "the-name",
			ExpectedFieldBody:           " the-body",
		},
		{
			Data:              "the-name :  the-body"+"\r\n"+
			                   "something: nothing",
			ExpectedFieldName: "the-name",
			ExpectedFieldBody:           "  the-body",
		},
		{
			Data:              "the-name  : the-body"+"\r\n"+
			                   "something: nothing",
			ExpectedFieldName: "the-name",
			ExpectedFieldBody:            " the-body",
		},
		{
			Data:              "the-name  :  the-body"+"\r\n"+
			                   "something: nothing",
			ExpectedFieldName: "the-name",
			ExpectedFieldBody:            "  the-body",
		},



		{
			Data:              "The-Name:the, body"+"\r\n"+
			                   "key: value"+"\r\n",
			ExpectedFieldName: "The-Name",
			ExpectedFieldBody:          "the, body",
		},
		{
			Data:              "The-Name: the, body"+"\r\n"+
			                   "key: value"+"\r\n",
			ExpectedFieldName: "The-Name",
			ExpectedFieldBody:          " the, body",
		},
		{
			Data:              "The-Name:  the, body"+"\r\n"+
			                   "key: value"+"\r\n",
			ExpectedFieldName: "The-Name",
			ExpectedFieldBody:          "  the, body",
		},
		{
			Data:              "The-Name :the, body"+"\r\n"+
			                   "key: value"+"\r\n",
			ExpectedFieldName: "The-Name",
			ExpectedFieldBody:           "the, body",
		},
		{
			Data:              "The-Name  :the, body"+"\r\n"+
			                   "key: value"+"\r\n",
			ExpectedFieldName: "The-Name",
			ExpectedFieldBody:            "the, body",
		},
		{
			Data:              "The-Name : the, body"+"\r\n"+
			                   "key: value"+"\r\n",
			ExpectedFieldName: "The-Name",
			ExpectedFieldBody:           " the, body",
		},
		{
			Data:              "The-Name :  the, body"+"\r\n"+
			                   "key: value"+"\r\n",
			ExpectedFieldName: "The-Name",
			ExpectedFieldBody:           "  the, body",
		},
		{
			Data:              "The-Name  : the, body"+"\r\n"+
			                   "key: value"+"\r\n",
			ExpectedFieldName: "The-Name",
			ExpectedFieldBody:            " the, body",
		},
		{
			Data:               "The-Name  :  the, body"+"\r\n"+
			                   "key: value"+"\r\n",
			ExpectedFieldName: "The-Name",
			ExpectedFieldBody:             "  the, body",
		},


		{
			Data:
				"One: 1"+"\r\n"+
				"     2"+"\r\n"+
				"\t3"+"\r\n"+
				"Two: a"+"\r\n"+
				"     b"+"\r\n"+
				"     c"+"\r\n",
			ExpectedFieldName: "One",
			ExpectedFieldBody: " 1\r\n     2\r\n\t3",
		},
	}

	for testNumber, test := range tests {

		var reader io.Reader = strings.NewReader(test.Data)
		wrapped := utf8.RuneScannerWrap(reader)
		var runescanner io.RuneScanner = &wrapped

		actualFieldName, actualFieldBody, err := netfield.Parse(runescanner)
		if nil != err {
			t.Errorf("For test #%d, did not expect an error, but actually got one.", testNumber)
			t.Logf("ERROR: (%T) %q", err, err)
			t.Logf("DATA: %q", test.Data)
			t.Logf("EXPECTED-FIELD-NAME: %q", test.ExpectedFieldName)
			t.Logf("EXPECTED-FIELD-BODY: %q", test.ExpectedFieldBody)
			continue
		}

		{
			var expected string = test.ExpectedFieldName
			var actual   string = actualFieldName

			if expected != actual {
				t.Errorf("For test #%d, the actual value for the field-name is not what was expected.", testNumber)
				t.Logf("EXPECTED: %q", expected)
				t.Logf("ACTUAL:   %q", actual)
				t.Logf("DATA: %q", test.Data)
				t.Logf("EXPECTED-FIELD-BODY: %q", test.ExpectedFieldBody)
				t.Logf("ACTUAL-FIELD-BODY:   %q", actualFieldBody)
				continue
			}
		}

		{
			var expected string = test.ExpectedFieldBody
			var actual   string = actualFieldBody

			if expected != actual {
				t.Errorf("For test #%d, the actual value for the field-body is not what was expected.", testNumber)
				t.Logf("EXPECTED: %q", expected)
				t.Logf("ACTUAL:   %q", actual)
				t.Logf("DATA: %q", test.Data)
				t.Logf("EXPECTED-FIELD-NAME: %q", test.ExpectedFieldName)
				t.Logf("ACTUAL-FIELD-NAME:   %q", actualFieldName)
				continue
			}
		}
	}
}
