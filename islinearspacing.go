package netfield

// isLinearSpacing returns whether a rune is linear-spacing or not.
//
// In IETF RFC822, "LWSP-char" is defined as:
//
//	LWSP-char   =  SPACE / HTAB
//
// This package expands the definition of "LWSP-char" to embrace Unicode.
// And includes a number of other character.
func isLinearSpacing(r rune) bool {

	switch r {
	case
		'\u0009', // horizontal tab
		'\u0020', // space
//		'\u00A0', // no-break space
		'\u1680', // ogham space mark
		'\u180E', // mongolian vowel separator
		'\u2000', // en quad
		'\u2001', // em quad
		'\u2002', // en space
		'\u2003', // em space
		'\u2004', // three-per-em space
		'\u2005', // four-per-em space
		'\u2006', // six-per-em space
		'\u2007', // figure space
		'\u2008', // punctuation space
		'\u2009', // thin space
		'\u200A', // hair space
//		'\u202F', // narrow no-break space
		'\u205F', // medium mathematical space
		'\u3000': // ideographic space
		return true
	default:
		return false
	}
}
