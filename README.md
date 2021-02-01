# strmask

[![Godoc Reference](https://godoc.org/github.com/nikitaksv/strmask?status.svg)](http://godoc.org/github.com/nikitaksv/strmask)
[![Coverage Status](https://coveralls.io/repos/github/nikitaksv/strmask/badge.svg?branch=main)](https://coveralls.io/github/nikitaksv/strmask?branch=main)
[![FOSSA Status](https://app.fossa.com/api/projects/git%2Bgithub.com%2Fnikitaksv%2Fstrmask.svg?type=shield)](https://app.fossa.com/projects/git%2Bgithub.com%2Fnikitaksv%2Fstrmask?ref=badge_shield)

Easy mask for string

## Install
```sh
go get github.com/nikitaksv/strmask
```

## Usage

```go
package main

import (
	"fmt"

	"github.com/nikitaksv/strmask"
)

func main() {
	// Ex. phone number
	maskedValue, _ := strmask.Apply("+0 (000) 000-00-00", "12345678900")
	fmt.Println(maskedValue) // out: +1 (234) 567-89-00

	// Ex. phone number with err
	maskedValue, err := strmask.Apply("+0 (000) 000-00-00", "12345")
	if err != nil {
		fmt.Println(err) // out: invalid value
	}
	fmt.Println(maskedValue) // but maskedValue out: +1 (234) 5

	// Ex. ID
	maskedValue, _ = strmask.Apply("####-####-####-####", "1234qwerasdfzxcv")
	fmt.Println(maskedValue) // out: 1234-qwer-asdf-zxcv

	// Ex. Passport
	maskedValue, _ = strmask.Apply("00 0000000", "123456789")
	fmt.Println(maskedValue) // out: 12 3456789
	
	// Ex. code
	maskedValue, _ = strmask.Apply("G-UUUUUU", "q1w2e3")
	fmt.Println(maskedValue) // out: G-Q1W2E3
}
```


## Mask Symbols
Symbol | Meaning
--- | ---
`#` | Requires a unicode letter or digit at this position
`0` | Requires a digit at this position
`L` | Requires a unicode letter (will be lower) or digit at this position
`l` | Requires a unicode letter (will be lower) at this position
`U` | Requires a unicode letter (will be upper) or digit at this position
`u` | Requires a unicode letter (will be upper) at this position
