package main

import (
	"fmt"
	"slices"
)

// slice -> dynamic
// most used construct in go
// + useful methods
func main() {
	//uninitialized slice is null
	var nums []int
	fmt.Println(nums)
	fmt.Println(len(nums))
	fmt.Println(nums == nil)

	//capacity -> maximum numbers of elements can fit in the slice
	var nam = make([]int, 2)
	// fmt.Println(cap(nam))

	nam = append(nam, 1)
	nam = append(nam, 2)
	nam = append(nam, 3)
	nam = append(nam, 4)
	fmt.Println(nam)
	fmt.Println(cap(nam))

	var nums1 = make([]int, 0, 5)
	nums1 = append(nums1, 2)
	var nums2 = make([]int, len(nums1))
	//copy function
	copy(nums2, nums1)
	fmt.Println(nums1, nums2)

	// slice operator

	var num = []int{1, 2, 3, 4, 5}
	fmt.Println(num[0:1])
	fmt.Println(num[:2])
	fmt.Println(num[1:])

	// slices
	var numk = []int{1, 2, 3}
	var numk2 = []int{1, 2, 4}

	fmt.Println(slices.Equal(numk, numk2))

	var nomk = [][]int{{1, 2, 3}, {4, 5, 6}}
	fmt.Println(nomk)

}
