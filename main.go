package main

import (
	"fmt"
	"testcoding/myfunc"
)

func main() {

	fmt.Println("\n==Calling function to move zeros to end of array: ")
	array := []int{2, 0, 0, 3, 5, 7, 11, 13}
	fmt.Println("", array)
	myfunc.MoveZeros(array)
	fmt.Println("after function call :\n", array)

	fmt.Println("\n==Tree Construction")
	in := []int{9, 3, 15, 20, 7}
	post := []int{9, 15, 7, 20, 3}
	fmt.Println(" inorder: ", in)
	fmt.Println(" postorder:", post)
	var tree *myfunc.Node
	tree = myfunc.CreateTree(in, post)
	println("Preorder of the constructed tree : ")
	myfunc.PreOrder(tree)

	fmt.Println("\n\n==Counting unique words:")
	passage := `ZBIO has come to campus for hiring today. ZBIO is a VC funded company in the enterprise infrastructure domain. This company is creating a platform to observe service to service communication`
	fmt.Println(" ", passage)
	count := myfunc.CountWords(passage)
	fmt.Println("Unique words: ", count)

	fmt.Println("\n==Merge arrays:")
	arr1 := []int{1, 3, 4, 5}
	arr2 := []int{2, 4, 6, 8}
	fmt.Println(" array 1:", arr1)
	fmt.Println(" array 2:", arr2)
	myfunc.MergeArrays(arr1, arr2)

	fmt.Println("\n\n==LRU Cache :")
	obj := myfunc.LRUConstructor(3)
	obj.Add(1, 10)
	obj.Add(2, 20)
	fmt.Println(obj.Get(1))
	obj.Add(3, 30)
	obj.Add(4, 40)
	fmt.Println(obj.Get(2))
	obj.Add(1, 10)
	obj.Add(5, 50)
	obj.Add(7, 70)
	fmt.Println(obj.Get(5))
	fmt.Println(obj.Get(3))
	obj.Show()
	fmt.Println("Number of elements Cache : ", obj.Len())

	fmt.Println("\n==Key Mapping :")
	arr3 := []int{1, 3, 4, 5}
	arr4 := []int{2, 4, 6, 8}
	fmt.Printf(" array 1: %v \n array 2: %v \n ", arr3, arr4)
	fmt.Println(myfunc.MapArray(arr1, arr2))

	fmt.Println("\n==Prime number!")
	val := 24
	fmt.Printf(" value = %d is %v\n", val, myfunc.IsPrimeIter(1))
	fmt.Printf(" value = %d is %v\n", val, myfunc.IsPrimeRecur(1))

}
