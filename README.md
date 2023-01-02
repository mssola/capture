<p align="center">
  <a href="https://github.com/mssola/capture/actions?query=workflow%3ACI" title="CI status for the main branch"><img src="https://github.com/mssola/capture/workflows/CI/badge.svg" alt="Build Status for main branch" /></a>
  <a href="https://pkg.go.dev/github.com/mssola/capture" rel="nofollow"><img alt="GoDoc" src="https://pkg.go.dev/badge/github.com/mssola/capture" style="max-width:100%;"></a>
</p>

---

This is a super simple package that basically allows you to capture the output
from os.Stdout and os.Stderr in a safe and clean manner. This is how you should
use it:

```go
package main

import (
	"fmt"
	"os"

	"github.com/mssola/capture"
)

func willPanic() {
	panic("Told ya!")
}

func foo() {
	fmt.Fprintf(os.Stdout, "Hello")
	fmt.Fprintf(os.Stderr, "World")
}

func main() {
	res := capture.All(func() { willPanic() })
	fmt.Printf("%v\n", res.Error)
	// Output: "Panic: Told ya!"

	res = capture.All(func() { foo() })
	fmt.Printf("%s %s\n", res.Stdout, res.Stderr)
	// Output: "Hello World"
}
```

Copyright &copy; 2015-2023 Miquel Sabaté Solà, released under the MIT License.
