package myfunc

// MoveZeros parse an integer array which will be
// updated with zero elements ending last
func MoveZeros(arr []int) {
	count := 0
	for i := 0; i < len(arr); i++ {
		if arr[i] != 0 {
			arr[count] = arr[i]
			count++
		}
	}
	for count < len(arr) {
		arr[count] = 0
		count++
	}
}
