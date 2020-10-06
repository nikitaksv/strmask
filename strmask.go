package strmask

import (
	"errors"
	"unicode"
)

const (
	LetterDigit = '#'
	Digit       = '0'

	LowerLetterDigit = 'L'
	LowerLetter      = 'l'

	UpperLetterDigit = 'U'
	UpperLetter      = 'u'
)

var ErrValue = errors.New("invalid value")

func Apply(mask, value string) (string, error) {
	valueRunes := []rune(value)
	runes := make([]rune, 0, len(mask))
	cursor := 0

	var r rune
	for _, mT := range mask {
		if isMaskType(mT) {
			r = getRune(mT, valueRunes, &cursor)
			if r == -1 {
				return string(runes), ErrValue
			}
		} else {
			r = mT
		}
		runes = append(runes, r)
	}

	return string(runes), nil
}

func getRune(maskType rune, value []rune, cursor *int) rune {
	if (*cursor) > len(value)-1 {
		return -1
	}

	r := value[*cursor]
	*cursor++

	// # - Requires a unicode letter or digit at this position
	// 0 - Requires a digit at this position
	if (maskType == LetterDigit && isLetterOrDigit(r)) || (maskType == Digit && unicode.IsDigit(r)) {
		return r
	}

	// L - Requires a unicode letter (will be lower) or digit at this position
	// l - Requires a unicode letter (will be lower) at this position
	if (maskType == LowerLetterDigit && isLetterOrDigit(r)) || (maskType == LowerLetter && unicode.IsLetter(r)) {
		return unicode.ToLower(r)
	}

	// U - Requires a unicode letter (will be upper) or digit at this position
	// u - Requires a unicode letter (will be upper) at this position
	if (maskType == UpperLetterDigit && isLetterOrDigit(r)) || (maskType == UpperLetter && unicode.IsLetter(r)) {
		return unicode.ToUpper(r)
	}

	return getRune(maskType, value, cursor)
}

func isLetterOrDigit(r rune) bool {
	return unicode.IsLetter(r) || unicode.IsDigit(r)
}

func isMaskType(r rune) bool {
	switch r {
	case LetterDigit,
		Digit,
		LowerLetterDigit,
		LowerLetter,
		UpperLetterDigit,
		UpperLetter:
		return true
	default:
		return false
	}
}
