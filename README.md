# go-netfield

**go-netfield** provides tools for parsing "net fields", for the Go programming language —
i.e., the type of fields found in HTTP headers & trailers, as well as SMTP headers.

Net-Fields look like this:
```
Content-Type: text/plain
Subject: Hello world!
Something: once
           twice
           thrice
           fource
Fruits:	apple banana cherry
```

Note that each line is terminated with a `"\r\n"`.

Also note that multi-line net-fields are possible.
To create a multi-line Net-Field, the beginning of the continuing field-body line must start with any type of linear-spacing.

## Documention

Online documentation, which includes examples, can be found at: http://godoc.org/github.com/reiver/go-netfield

[![GoDoc](https://godoc.org/github.com/reiver/go-netfield?status.svg)](https://godoc.org/github.com/reiver/go-netfield)

## Examples

```go
import "github.com/reiver/go-netfield"

// ...

fieldName, fieldBody, err := netfield.Parse(runescanner)
```

## Linear Spacing

In IETF RFC822, the linear-spacing characters are defined by the definition for "LWSP-char":
```
LWSP-char   =  SPACE / HTAB
```
This package expands the definition of "LWSP-char" to embrace Unicode.
And includes a number of other character.

* `U+0009` — horizontal tab (␉)
* `U+0020` — space (␠)
* `U+1680` — ogham space mark
* `U+180E` — mongolian vowel separator
* `U+2000` — en quad
* `U+2001` — em quad
* `U+2002` — en space
* `U+2003` — em space
* `U+2004` — three-per-em space
* `U+2005` — four-per-em space
* `U+2006` — six-per-em space
* `U+2007` — figure space
* `U+2008` — punctuation space
* `U+2009` — thin space
* `U+200A` — hair space
* `U+205F` — medium mathematical space
* `U+3000` — ideographic space
