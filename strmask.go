package strmask

import (
	"errors"
	"unicode"
)

const (
	Letter      = 'l'
	LetterDigit = 'L'
	Digit       = '0'
)

var ErrValue = errors.New("invalid value")

func Apply(mask, value string) (string, error) {
	var err error
	runes := make([]rune, len(mask))
	cursor := 0
	for i, mT := range mask {
		runes[i], err = getRune(mT, value, &cursor)
		if err != nil {
			return "", err
		}
	}

	str := string(runes)
	return str, nil
}

func getRune(maskType rune, value string, cursor *int) (rune, error) {
	if (*cursor) > len(value)-1 {
		return 0, ErrValue
	}

	r := rune(value[*cursor])

	if maskType == r {
		*cursor++
		return maskType, nil
	} else if isMaskType(maskType) {
		*cursor++
		if (maskType == LetterDigit && isLetterDigit(r)) ||
			(maskType == Letter && unicode.IsLetter(r)) ||
			(maskType == Digit && unicode.IsDigit(r)) {
			return r, nil
		}
		return getRune(maskType, value, cursor)
	} else {
		return maskType, nil
	}
}

func isLetterDigit(r rune) bool {
	return unicode.IsLetter(r) || unicode.IsDigit(r)
}

func isMaskType(r rune) bool {
	switch r {
	case LetterDigit, Letter, Digit:
		return true
	default:
		return false
	}
}
