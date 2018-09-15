// Run the following command:
// go run quickSort.go 5 3 4 2 1 7 6
package main

import (
	"flag"
	"fmt"
	"log"
	"strconv"
)

// main() receives command line arguments and returns a sorted array.
// See quickSort() for quick-sort.
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

	quickSort(argsInt, 0, len(argsInt)-1)

	for _, v := range argsInt {
		fmt.Printf("%d\n", v)
	}
}

func quickSort(a []int, p int, r int) {
	// a divede-conquer algorithm
	if p < r {
		q := partition(a, p, r) // divide
		quickSort(a, p, q-1) // conquer
		quickSort(a, q+1, r) // conquer
	}
}

// We will divide a[p:r+1] into a[p:q] and a[q+1:r+1] such that
// a[q] is not less than all numbers in a[p:q],
// and a[q] is not more than all numbers in a[q+1:r+1].
// partition returns the pivot, q.
func partition(a []int, p int, r int) int {
	// low will move from p to q.
	// It will be the case that a[low] <= a[q], low = 0, 1, ..., q-1.
	low := p

	// We will compare a[i] and a[r], i = p, p+1, ..., r-1.
	// a[r] will be the pivot, a[q].
	for i := p; i < r; i++ {
		if a[i] < a[r] {
			swap(a, low, i)
			low++
		}
	}

	swap(a, low, r)

	return low
}

func swap(a []int, i int, j int) {
	tmp := a[i]
	a[i] = a[j]
	a[j] = tmp
}
