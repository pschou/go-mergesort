# Parallel Merge Sorting Scanner

A bufio.Scanner like reader which reads from multiple io.Readers and returns an ordered list

```golang
import (
  ...
  ms "github.com/pschou/go-mergesort"
)

func ExampleNew() {
  s := ms.New(context.TODO(), bufio.ScanLines, ms.BytesCompareDedup,
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
  // g
  // z
}
```
