package vex

import (
	"errors"
	"strconv"
)

// Format v into a hex string with a leading digit indicating its length.
func Format(v uint64) string {
	s := strconv.FormatUint(v, 16)
	s = strconv.FormatUint(uint64(len(s))-1, 16) + s
	return s
}

// Parse s as a hex string with the leading digit indicating its length.
func Parse(s string) (uint64, error) {
	if len(s) <= 1 {
		return 0, errors.New("vex.Parse(): value must have length of at least 2")
	} else if !isHex(s[0]) {
		return 0, errors.New("vex.Parse(): invalid prefix")
	}

	v, err := strconv.ParseUint(s[1:], 16, 64)
	if err != nil {
		return 0, errors.New("vex.Parse(): invalid hex value")
	}
	return v, nil
}

// isHex returns true if character is a hex digit.
func isHex(ch byte) bool {
	return (ch >= '0' && ch <= '9') || (ch >= 'a' && ch <= 'f')
}
