package myfunc

import (
	"fmt"
)

// MergeArrays comment
func MergeArrays(arr1 []int, arr2 []int) {
	var i, k, j int
	arr3 := make([]int, len(arr1)+len(arr2))
	for ; i < len(arr1) && j < len(arr2); k++ {
		if arr1[i] < arr2[j] {
			arr3[k] = arr1[i]
			i++
		} else {
			arr3[k] = arr2[j]
			j++
		}
	}
	for ; i < len(arr1); i++ {
		arr3[k] = arr1[i]
		k++
	}
	for ; j < len(arr2); j++ {
		arr3[k] = arr2[j]
		k++
	}
	fmt.Println("Merged array:")
	for j := range arr3 {
		fmt.Printf(" %d", arr3[j])
	}
}
