# strmask
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
	maskedValue, _ := strmask.Apply("LLL-000-0", "ABC1234")
	fmt.Println(maskedValue) // out: ABC-123-4

	maskedValue, _ = strmask.Apply("LLL-000-0", "ABC-123-4")
	fmt.Println(maskedValue) // out: ABC-123-4

	maskedValue, _ = strmask.Apply("LLL-000-0", "ABC1234XYZ")
	fmt.Println(maskedValue) // out: ABC-123-4
}

```


## Mask Symbols
Symbol | Meaning
--- | ---
`0` | Requires a decimal digit at this position
`L` | Requires a unicode letter and digit at this position
`l` | Requires a unicode letter at this position
