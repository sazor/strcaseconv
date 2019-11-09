# strcaseconv

## Overview [![GoDoc](https://godoc.org/github.com/sazor/strcaseconv?status.svg)](https://godoc.org/github.com/sazor/strcaseconv) [![Go Report Card](https://goreportcard.com/badge/github.com/sazor/strcaseconv)](https://goreportcard.com/report/github.com/sazor/strcaseconv) [![License](https://img.shields.io/github/license/sazor/strcaseconv)](https://github.com/sazor/strcaseconv) [![Go version](https://img.shields.io/badge/go-v1.10-blue)](https://github.com/sazor/strcaseconv)

String case converter inspired by stdlib `strings` package. Fast and simple. Conversions make only 1 allocation of resulting string if changes are needed, otherwise, original string is returned.

Use with caution. Original string is expected to be the exact case specified in method name. No validation or intermediate conversion included. No additional processing like trimming included. ASCII only at the moment. Such limitations are applied to keep it fast and simple with the least possible amount of allocations. Designed to be used in strictly predictable cases such as database records to structs mapping.  

## Install

```
go get github.com/sazor/strcaseconv
```

## Example

```go
strcaseconv.SnakeToUpperCamel("foo_bar") // FooBar
strcaseconv.SnakeToLowerCamel("foo_bar") // fooBar
strcaseconv.KebabToUpperCamel("foo-bar") // FooBar
strcaseconv.KebabToLowerCamel("foo-bar") // fooBar
strcaseconv.CamelToSnake("FooBar") // foo_bar
strcaseconv.CamelToKebab("fooBar") // foo-bar
strcaseconv.SnakeToKebab("foo_bar") // foo-bar
strcaseconv.KebabToSnake("foo-bar") // foo_bar
```

## License

MIT.
