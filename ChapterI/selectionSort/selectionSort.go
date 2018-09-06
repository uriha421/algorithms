package main

import (
	"flag"
	"fmt"
	"log"
	"strconv"
)

// main() receives command line arguments and returns a sorted array.
// See selectionSort() for selection-sort.
func main() {
  flag.Parse()

	var err error
  argsInt := make([]int, flag.NArg())

  for i, v := range flag.Args() {
		argsInt[i], err = strconv.Atoi(v)
		if err != nil {
			log.Fatal(err)
		}
	}

  selectionSort(argsInt)

  for _, v := range argsInt {
    fmt.Printf("%d\n", v)
  }

}


func selectionSort(a []int) {
  var k, l int
  for i := 0; i < len(a) - 1; i++ {
    k = a[i]
    l = i

    // k will be the least number in a[i+1:len(a)].
    // l will be the index of the least number in a[i+1:len(a)].
    for j := i + 1; j < len(a); j++ {
      if (a[i] > a[j]) {
        k = a[j]
        l = j
      }
    }

    // a[0:i+1] has been sorted.
    a[l] = a[i]
    a[i] = k
  }
}
