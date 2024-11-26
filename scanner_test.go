package mergesort_test

import (
	"bufio"
	"context"
	"fmt"
	"io"
	"os"
	"strings"

	ms "github.com/pschou/go-mergesort"
)

func ExampleNew() {
	s := ms.New(context.TODO(), bufio.ScanLines,
		strings.NewReader("a\nc\nd\n"),
		strings.NewReader("b\ne\nf\ng\n"),
		strings.NewReader("f\nz\n"),
	)

	for s.Scan() {
		fmt.Println(s.Text())
	}
	// Output:
	// a
	// b
	// c
	// d
	// e
	// f
	// f
	// g
	// z
}

func ExampleNewWithoutScanFirst() {
	list := []io.Reader{
		strings.NewReader("b\ne\nf\ng\n"),
		strings.NewReader("f\nz\n"),
		strings.NewReader("a\nc\nd\n"),
		strings.NewReader("x\ny\n"),
	}

	s := ms.New(context.TODO(), bufio.ScanLines, list...)
	s.Scan()
	fmt.Println(s.Text())
	// Output:
	// a
}

func ExampleFileReader() {
	a, _ := os.Open("a")
	b, _ := os.Open("b")
	c, _ := os.Open("c")
	list := []io.Reader{a, b, c}

	s := ms.New(context.TODO(), bufio.ScanLines, list...)

	out, _ := os.Create("out")
	for s.Scan() {
		fmt.Fprintf(out, "%s\n", s.Text())
	}
	// Output:
}
