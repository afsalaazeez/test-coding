package myfunc

// MapArray function
func MapArray(arr1 []int, arr2 []int) map[int]int {
	m := make(map[int]int)
	for i := range arr2 {
		m[arr1[i]] = arr2[i]
	}
	return m
}
