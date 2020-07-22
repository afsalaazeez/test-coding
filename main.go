package main

import (
	"fmt"
	"testcoding/myfunc"
)

func main() {

	fmt.Println("\nCalling function to move zeros to end of array: ")
	primes := []int{2, 0, 0, 3, 5, 7, 11, 13}
	fmt.Println(primes)
	myfunc.MoveZeros(primes)
	fmt.Println("after function call :\n", primes)

	in := []int{9, 3, 15, 20, 7}
	post := []int{9, 15, 7, 20, 3}
	fmt.Println("\ninorder: ", in)
	fmt.Println("postorder:", post)
	var root *myfunc.Node
	root = myfunc.CreateTree(in, post)
	println("Preorder of the constructed tree : ")
	myfunc.PreOrder(root)

	fmt.Println("\n\ncounting unique words in the passage :")
	passage := `ZBIO has come to campus for hiring today. ZBIO is a VC funded company in the enterprise infrastructure domain. This company is creating a platform to observe service to service communication`
	fmt.Println(passage)
	count := myfunc.CountWords(passage)
	fmt.Println("number of unique words: ", count)

	fmt.Println("Merge arrays:")
	arr1 := []int{1, 3, 4, 5}
	arr2 := []int{2, 4, 6, 8}
	fmt.Println("array 1:", arr1)
	fmt.Println("array 2:", arr2)
	myfunc.MergeArrays(arr1, arr2)

	fmt.Println("")
}
