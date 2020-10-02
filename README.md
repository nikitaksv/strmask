# strmask
Easy mask for string

## Install
```sh
go get github.com/nikitaksv/strmask
```

## Usage
```go
maskedValue, _ := strmask.Apply("LLL-000-0", "ABC1234")
fmr.Println(maskedValue) // out: ABC-123-4

maskedValue, _ = strmask.Apply("LLL-000-0", "ABC-123-4")
fmr.Println(maskedValue) // out: ABC-123-4 

maskedValue, _ = strmask.Apply("LLL-000-0", "ABC1234XYZ")
fmr.Println(maskedValue) // out: ABC-123-4 
```


## Mask Symbols
Symbol | Meaning
--- | ---
`0` | Requires a decimal digit at this position
`L` | Requires an unicode letter and digit at this position
`l` | Permits an unicode letter at this position
