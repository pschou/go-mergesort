package mergesort_test

import (
	"context"
	"fmt"
	"io"
	"os"
	"strings"

	ms "github.com/pschou/go-mergesort"
)

func ExampleNew() {
	s := ms.New(context.TODO(), ms.ScanLines, ms.BytesCompareDedup,
		strings.NewReader("a\nc\nd\nz\n"),
		strings.NewReader("b\ne\nf\ng\nz\n"),
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

	s := ms.New(context.TODO(), ms.ScanLines, ms.BytesCompare, list...)
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

	s := ms.New(context.TODO(), ms.ScanLines, ms.BytesCompare, list...)

	out, _ := os.Create("out")
	for s.Scan() {
		fmt.Fprintf(out, "%s\n", s.Text())
	}
	// Output:
}
