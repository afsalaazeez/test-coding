# Go Language tutorial problems
### 1.	Given an array, write a function to move all 0's to the end of it while maintaining the relative order of the non-zero elements.

 Example:
 ```
 Input: [0,1,0,3,12]
 Output: [1,3,12,0,0]
 ```

### 2.	Given inorder and postorder traversal of a tree, construct the binary tree.

Note:
You may assume that duplicates do not exist in the tree.

Example:
```
  inorder = [9,3,15,20,7], postorder = [9,15,7,20,3]
```
Return the following binary tree:
```
    3
   / \
  9  20
    /  \
   15   7
```
### 3.	Count the number of unique words (which are not repeated) in a given passage.

	Note:
	Ignore . , : ’ ” etc

Example:
“ZBIO has come to campus for hiring today. ZBIO is a VC funded company in the enterprise infrastructure domain. This company is creating a platform to observe service to service communication” 
```
  Total count: 31
  Unique count: 18
```
### 4.	Given two sorted arrays, merge them in a sorted manner. 
```
  Example:
  Input: arr1[] = { 1, 3, 4, 5}, arr2[] = {2, 4, 6, 8}
  Output: arr3[] = {1, 2, 3, 4, 4, 5, 6, 8}
```
### 5.	Implement a program to show data structure with LRU implementation 
* Instantiate the data structure for a fixed size
* Provide add(), get(), show() and len() methods
* Get() item should be considered as item access/use for LRU to be updated
* Demonstrate through test code that repeated calls to add() for more than the fixed size leads to removing the elements as per the LRU to accommodate the new elements


### 6.	Given 2 arrays containing an identical number of items, iterate through them and create a hash/map with an element from the first array as key and element from the second array as value.


### 7.	Write logic to determine whether a number is prime or not. Implement this logic once using iteration and another time as recursion.


### 8.	Hash or map is by default a key store without order/sequence. Implement an ordered map as your own data structure, so that a retrieve on this data structure will return the key, value tuple in the same order as they were populated.