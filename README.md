# capture [![Build Status](https://travis-ci.org/mssola/capture.svg?branch=master)](https://travis-ci.org/mssola/capture) [![GoDoc](https://godoc.org/github.com/mssola/capture?status.png)](http://godoc.org/github.com/mssola/capture)

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

## License

Copyright &copy; 2015 Miquel Sabaté Solà

Permission is hereby granted, free of charge, to any person obtaining
a copy of this software and associated documentation files (the
"Software"), to deal in the Software without restriction, including
without limitation the rights to use, copy, modify, merge, publish,
distribute, sublicense, and/or sell copies of the Software, and to
permit persons to whom the Software is furnished to do so, subject to
the following conditions:

The above copyright notice and this permission notice shall be
included in all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND,
EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF
MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND
NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT HOLDERS BE
LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER IN AN ACTION
OF CONTRACT, TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN CONNECTION
WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.

