package capture

import (
	"fmt"
	"os"
	"testing"
)

func TestCapturePanic(t *testing.T) {
	res := All(func() { panic("Panic!") })
	if res.Stdout != nil {
		t.Fatal("Wrong value for stdout")
	}
	if res.Stderr != nil {
		t.Fatal("Wrong value for stdout")
	}
	if res.Error.Error() != "panic: Panic!" {
		t.Fatal("Wrong error")
	}
}

func TestCapture(t *testing.T) {
	res := All(func() {
		fmt.Fprintf(os.Stdout, "stdout")
		fmt.Fprintf(os.Stderr, "stderr")
	})

	if res.Error != nil {
		t.Fatal("It shoud have failed")
	}
	if string(res.Stdout) != "stdout" {
		t.Fatal("Wrong stdout")
	}
	if string(res.Stderr) != "stderr" {
		t.Fatal("Wrong stderr")
	}
}

func ExampleAll() {
	res := All(func() { panic("Told ya!") })
	fmt.Printf("%v\n", res.Error)

	res = All(func() {
		fmt.Fprintf(os.Stdout, "Hello")
		fmt.Fprintf(os.Stderr, "World")
	})
	fmt.Printf("%s %s\n", res.Stdout, res.Stderr)

	// Output:
	// panic: Told ya!
	// Hello World
}
