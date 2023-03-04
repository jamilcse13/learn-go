package main

import (
	"fmt"
)

func main() {
	// Create a Slice With []datatype{values}
	mySlice1 := []int{}
	fmt.Println(len(mySlice1))
	fmt.Println(cap(mySlice1))
	fmt.Println(mySlice1)

	mySlice2 := []string{"Go", "Python", "C++"}
	fmt.Println(len(mySlice2))
	fmt.Println(cap(mySlice2))
	fmt.Println(mySlice2)
	fmt.Println();


	// Create a Slice From an Array
	arr1 := [6]int{10, 11, 12, 13, 14,15}
	my_slice := arr1[2:4]

	fmt.Printf("mySlice = %v\n", my_slice)
	fmt.Printf("length = %d\n", len(my_slice))
	fmt.Printf("capacity = %d\n", cap(my_slice))
	fmt.Println();


	// Create a Slice With The make() Function
	// slice_name := make([]type, length, capacity)
	my_slice2 := make([]int, 5, 10)
	fmt.Printf("mySlice2 = %v\n", my_slice2)
	fmt.Printf("length = %v\n", len(my_slice2))
	fmt.Printf("capacity = %v\n", cap(my_slice2))

	// with omitted capacity
	my_slice3 := make([]int, 5)
	fmt.Printf("mySlice3 = %v\n", my_slice3)
	fmt.Printf("length = %v\n", len(my_slice3))
	fmt.Printf("capacity = %v\n", cap(my_slice3))
	fmt.Println();


	test_slice := []int{10,20,30}
	fmt.Printf("testSlice = %v\n", test_slice)

	// Append Elements To a Slice
	// slice_name = append(slice_name, element1, element2, ...)
	test_slice = append(test_slice, 40, 50)
	fmt.Printf("testSlice with append = %v\n", test_slice)
	fmt.Printf("length = %v\n", len(test_slice))
	fmt.Printf("capacity = %v\n", cap(test_slice))
	fmt.Println();


	// Append One Slice To Another Slice
	// slice3 = append(slice1, slice2...)
	slice1 := []int{1,2,3}
	slice2 := []int{4,5,6}
	slice3 := append(slice1, slice2...)

	fmt.Printf("slice3=%v\n", slice3)
  	fmt.Printf("length=%d\n", len(slice3))
  	fmt.Printf("capacity=%d\n", cap(slice3))
	fmt.Println();


	//Change The Length of a Slice
	slice_arr1 := [6]int{9, 10, 11, 12, 13, 14} // An array
	fmt.Printf("Original array = %v\n", slice_arr1)
	myslice1 := slice_arr1[1:5] // Slice array
	fmt.Printf("myslice1 = %v\n", myslice1)
	fmt.Printf("length = %d\n", len(myslice1))
	fmt.Printf("capacity = %d\n", cap(myslice1))

	myslice1 = slice_arr1[1:3] // Change length by re-slicing the array
	fmt.Printf("myslice1 = %v\n", myslice1)
	fmt.Printf("length = %d\n", len(myslice1))
	fmt.Printf("capacity = %d\n", cap(myslice1))

	myslice1 = append(myslice1, 20, 21, 22, 23) // Change length by appending items
	fmt.Printf("myslice1 = %v\n", myslice1)
	fmt.Printf("length = %d\n", len(myslice1))
	fmt.Printf("capacity = %d\n", cap(myslice1))
	fmt.Println();


	// Memory Efficiency
	numbers := []int{1,2,3,4,5,6,7,8,9,10,11,12,13,14,15}
	// Original slice
	fmt.Printf("numbers = %v\n", numbers)
	fmt.Printf("length = %d\n", len(numbers))
	fmt.Printf("capacity = %d\n", cap(numbers))

	// Create copy with only needed numbers
	neededNumbers := numbers[:len(numbers)-10]
	numbersCopy := make([]int, len(neededNumbers))
	copy(numbersCopy, neededNumbers)

	fmt.Printf("numbersCopy = %v\n", numbersCopy)
	fmt.Printf("length = %d\n", len(numbersCopy))
	fmt.Printf("capacity = %d\n", cap(numbersCopy))
}