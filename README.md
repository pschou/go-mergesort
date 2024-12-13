# Parallel Merge Sorting Scanner

A bufio.Scanner like reader which reads from multiple io.Readers and returns an
ordered list.  Multiple go-routines and connecting channels are created to
speed up the sorting.  This way multiple file reads are done in parallel along
with with each node in the binary tree doing bottom up sorting.  The final
sorted channel is exposed as the standard Scan(), Bytes(), and Text() that one
is familiar with from bufio.Scanner.

An additional FilterFunc is provided to parallelize any computation which needs
to be done to the record before it is consumed.  An example of this may be a
conversion of record types, and an example of such use case is doing a parseInt
to create a 4 byte unsigned integer for data space optimization.

The CompareFunction can be useful in cases that records in the original file
may be case insensitive ordering, in which case the compare function can be
provided which ignores the case of the input byte slice to provide a ordering
resultant.

```golang
import (
  ...
  ms "github.com/pschou/go-mergesort"
)

func ExampleNew() {
  s := ms.New(context.TODO(),
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
