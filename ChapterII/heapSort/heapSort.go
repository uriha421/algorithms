// Run the following command:
// go run heapSort.go 5 3 4 2 1 7 6
package main

import (
	"flag"
	"fmt"
	"log"
	"strconv"
)

// main() receives command line arguments and returns a sorted array.
// See heapSort() for heap-sort.
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

	heapSort(argsInt, len(argsInt)-1)

	for _, v := range argsInt {
		fmt.Printf("%d\n", v)
	}
}

func heapSort(a []int, heapSize int) {
  // We will sort a[] so that it will be a max-heap:
  // The number any parent has will be not less than the number its children have.
  buildMaxHeap(a, heapSize)

  for i := heapSize; i>0; i-- {
    // a[:i+1] has become max-heap, so a[0] is the largest in a[:i+1]
    // We will swap a[0] and a[i], so a[i:] will be sorted.
    swap(a, 0, i)

    // a[i:] will be removed from the heap-structure.
    heapSize--

    // The number any parent except a[0] has has become not less than the number its children have.
    // We will remake a max-heap.
    maxHeaptify(a, 0, heapSize)
  }
}

// We will make a[] become a max-heap.
func buildMaxHeap(a []int, heapSize int) {
  for i := heapSize/2-1; i>=0; i-- {
    maxHeaptify(a, i, heapSize)
  }
}

// If the number all children of a[i] have is not less than the number their parent has,
// then we will make a max-heap whose root is a[i].
func maxHeaptify(a []int, i int, heapSize int) {
  l := left(i)
  r := right(i)
  var maxIndex int

  if l<=heapSize && a[l]>a[i] {
    maxIndex = l
  } else {
    maxIndex = i
  }
  if r<=heapSize && a[r]>a[maxIndex] {
    maxIndex = r
  }
  if maxIndex != i {
    swap(a, i, maxIndex)
    maxHeaptify(a, maxIndex, heapSize)
  }
}

func left(i int) int {
  return 2*i+1
}

func right(i int) int {
  return 2*i+2
}

func swap(a []int, i int, j int) {
  tmp := a[i]
  a[i] = a[j]
  a[j] = tmp
}
