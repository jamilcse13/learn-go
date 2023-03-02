package main
import (
	"fmt"
)

func main() {
	var name_arr = [3]string{"jim","kim","tim"}
	num_arr := [3]int{1,2,3}

	fmt.Println(name_arr, num_arr)

	var test_arr = [...]int{1,2,3,5}
	var test_arr2 = [...]string{"loop","oop","wwe","asap"}

	fmt.Println(test_arr, test_arr2)


	// Access Elements of an Array
	fmt.Println(test_arr2[1])
	fmt.Println(name_arr[1])

	// Change Elements of an Array
	test_arr2[3] = "cam"
	fmt.Println(test_arr2)


	// Array Initialization
	arr1 := [5]int{} 	// not initialized
	arr2 := [5]int{1,2}  	// partially initialized
	arr3 := [5]int{1,2,3,4,5}  	// fully initialized

	fmt.Println(arr1)
	fmt.Println(arr2)
	fmt.Println(arr3)


	// Initialize Only Specific Elements
	arr4 := [5]int{1:10, 3:20}
	fmt.Println(arr4)


	// Find the Length of an Array
	fmt.Println(len(test_arr2))
}