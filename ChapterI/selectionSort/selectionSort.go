// Run the following command:
// go run selectionSort.go 5 3 4 2 1 7 6
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
	for i := 0; i < len(a)-1; i++ {

		// k will be the least number in a[i:len(a)].
		// l will be the index of the least number in a[i:len(a)].
		k = a[i]
		l = i
		for j := i + 1; j < len(a); j++ {
			if k > a[j] {
				k = a[j]
				l = j
			}
		}

		a[l] = a[i]
		a[i] = k
		// a[0:i] has been sorted.
	}
}
