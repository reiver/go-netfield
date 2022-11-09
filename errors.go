package netfield

import (
	"github.com/reiver/go-fck"
)

const (
	errColonNotFound  = fck.Error("colon not found")
	errNilRuneScanner = fck.Error("nil rune-scanner")
	errRuneError      = fck.Error("rune error")
)
