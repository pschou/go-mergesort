package mergesort_test

import (
	"context"
	"fmt"
	"io"
	"os"
	"strings"
	"unicode"

	ms "github.com/pschou/go-mergesort"
)

func ExampleNewWithDedup() {
	s := ms.New(context.TODO(),
		strings.NewReader("a\nc\nd\nz\n"),
		strings.NewReader("b\ne\nf\ng\nz\n"),
		strings.NewReader("f\nz\n"),
	)
	s.Split(ms.ScanLines)
	s.Compare(ms.BytesCompareDedup)

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

func ExampleNewFilter() {
	list := []io.Reader{
		strings.NewReader("b\ne\nf\ng\n"),
		strings.NewReader("d\nf\nz\n"),
		strings.NewReader("a\nc\nd\n"),
		strings.NewReader("x\ny\n"),
	}

	s := ms.New(context.TODO(), list...)
	// s.Split(ms.ScanLines)
	s.Compare(ms.BytesCompareDedup)
	s.Filter(func(in []byte) ([]byte, func()) {
		in[0] = byte(unicode.ToUpper(rune(in[0])))
		return in, nil
	})

	for s.Scan() {
		fmt.Println(s.Text())
	}
	// Output:
	// A
	// B
	// C
	// D
	// E
	// F
	// G
	// X
	// Y
	// Z
}

func ExampleNew() {
	list := []io.Reader{
		strings.NewReader("b\ne\nf\ng\n"),
		strings.NewReader("f\nz\n"),
		strings.NewReader("a\nc\nd\n"),
		strings.NewReader("x\ny\n"),
	}

	s := ms.New(context.TODO(), list...)
	// s.Split(ms.ScanLines)
	// s.Compare(ms.BytesCompare)

	s.Scan()
	fmt.Println(s.Text())
	// Output:
	// a
}

func ExampleFileReaderWriter() {
	a, _ := os.Open("a")
	b, _ := os.Open("b")
	c, _ := os.Open("c")
	list := []io.Reader{a, b, c}

	s := ms.New(context.TODO(), list...)
	// s.Split(ms.ScanLines)
	// s.Compare(ms.BytesCompare)

	out, _ := os.Create("out")
	for s.Scan() {
		fmt.Fprintf(out, "%s\n", s.Text())
	}
	// Output:
}
