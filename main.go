package main

import (
	"fmt"
	"sort-test/bubble_sort"
	"sort-test/merge_sort"
)

func main() {

	input2 := []int{3, 5, 1, 2, 88}
	fmt.Println("---ASC with ints---")
	fmt.Println("before:", input2)
	bubble_sort.Sort(input2, 1)
	fmt.Println("after:", input2)

	fmt.Println("---DESC with strings---")
	input3 := []string{"h", "c", "b", "z", "a"}
	fmt.Println("before:", input3)
	merge_sort.Sort(input3, -1)
	fmt.Println("after:", input3)

}
