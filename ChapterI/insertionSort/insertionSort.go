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
// See selectionSort() for insertion-sort.
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
		k = a[i]
		l = i - 1
		for k < a[l] && l >= 0 {
			a[l+1] = a[l]
			a[l] = k
			l--
		}
	}
}
