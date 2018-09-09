// Run the following command:
// go run mergeSort.go 5 3 4 2 1 7 6
package main

import (
	"flag"
	"fmt"
	"log"
	"strconv"
)

// main() receives command line arguments and returns a sorted array.
// See mergeSort() for merge-sort.
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

	mergeSort(argsInt, 0, len(argsInt)-1)

	for _, v := range argsInt {
		fmt.Printf("%d\n", v)
	}
}

func mergeSort(a []int, p int, r int) {

	// a Combine-Conquer algorithm
	if p < r {
		q := (p + r) / 2
		mergeSort(a, p, q)
		mergeSort(a, q+1, r)
		merge(a, p, q, r)
	}
}

func merge(a []int, p int, q int, r int) {
	b := make([]int, q-p+1)
	copy(b, a[p:q+1])
	c := make([]int, r-q)
	copy(c, a[q+1:r+1])
	// b and c has been sorted.

	// k and l is the argument of b and c respectively.
	k, l := 0, 0

	// We will merge two sorted arrays, b and c.
	for i := p; i <= r; i++ {
		if k < q-p+1 && l < r-q && b[k] < c[l] {
			a[i] = b[k]
			k++
		} else if k < q-p+1 && l < r-q && b[k] >= c[l] {
			a[i] = c[l]
			l++
		} else if k == q-p+1 {
			// all elements in b has been added in a.
			a[i] = c[l]
			l++
		} else if l == r-q {
			// all elements in c has been added in a.
			a[i] = b[k]
			k++
		}
	}
}
