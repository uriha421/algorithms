// Run the following command:
// go run insertionSort.go 5 3 4 2 1 7 6
package main

import (
	"flag"
	"fmt"
	"log"
	"strconv"
)

// main() receives command line arguments and returns a sorted array.
// See insertionSort() for insertion-sort.
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

	insertionSort(argsInt)

	for _, v := range argsInt {
		fmt.Printf("%d\n", v)
	}
}

func insertionSort(a []int) {
	var k, l int

	for i := 1; i < len(a); i++ {
		// a[:i] has been sorted.

		k = a[i]  // k will be inserted into a[:i], and a[:i+1] will be sorted.
		l = i - 1 // We will compare k with a[l], (l = i-1, ..., 0).
		for k < a[l] && l >= 0 {
			a[l+1] = a[l]
			a[l] = k
			l--
		}
	}
}
