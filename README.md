# strmask
Easy mask for string

## Install
```sh
go get github.com/nikitaksv/strmask
```

## Usage

Playground: https://play.golang.org/p/BUxPS2c2JT6

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
