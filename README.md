# obfuscate: Go functions for obfuscating sensitive data

[![GoDoc](https://godoc.org/github.com/networkteam/obfuscate?status.svg)](https://godoc.org/github.com/networkteam/obfuscate)
[![Build Status](https://github.com/networkteam/obfuscate/workflows/run%20tests/badge.svg)](https://github.com/networkteam/obfuscate/actions?workflow=run%20tests)
[![Go Report Card](https://goreportcard.com/badge/github.com/networkteam/obfuscate)](https://goreportcard.com/report/github.com/networkteam/obfuscate)

## Features

* Mask email addresses

## Scope

Log personal data like email addresses for debugging purposes.

## Limitations

You should not rely on this package for security.
Sensitive data should not be stored or stored in a secure way in the first place.

## Install

```sh
go get github.com/networkteam/obfuscate
```

## Usage

```go
package main

import (
	"fmt"

	"github.com/networkteam/obfuscate"
)

func main() {
	fmt.Println(obfuscate.EmailAddressPartially("john.doe@example.com"))
	// Output:
	// j*n.d*e@e*e.com
}
```