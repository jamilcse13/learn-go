## Varibalbe:
* Var Type:
    - int
    - float32
    - string
    - bool

* Used **var** keyword
* Used **"=** sign

* Difference Between var and := :

| var | := |
| --- | --- |
| Can be used inside and outside of functions | Can only be used inside functions |
| Variable declaration and value assignment can be done separately | Variable declaration and value assignment cannot be done separately (must be done in the same line) |


**Go Multiple Variable Declaration:**
In Go, it is possible to declare multiple variables in the same line. Example:
```
package main
import ("fmt")

func main() {
  var a, b, c, d int = 1, 3, 5, 7

  fmt.Println(a)
  fmt.Println(b)
  fmt.Println(c)
  fmt.Println(d)
}
```
If you use the type keyword, it is only possible to declare one type of variable per line.

## Constants:
* Constant Rules:
    - Constant names follow the same naming rules as variables
    - Constant names are usually written in uppercase letters (for easy identification and differentiation from variables)
    - Constants can be declared both inside and outside of a function
  
 ## Output:
- Go has three functions to output text:
  - Print()
  - Println()
  - Printf()

**Print():** 
- The Print() function prints its arguments with their default format. 
- If we want to print the arguments in new lines, we need to use \n.
- It is also possible to only use one Print() for printing multiple variables.
- Print() inserts a space between the arguments if neither are strings:

**Println():** The Println() function is similar to Print() with the difference that a whitespace is added between the arguments, and a newline is added at the end.

**Printf():** The Printf() function first formats its argument based on the given formatting verb and then prints them. Here we will use two formatting verbs:
- %v is used to print the value of the arguments
- %T is used to print the type of the arguments


## Array:
There are two ways to declare an array.
- With the var keyword:
  ```
  var array_name = [length]datatype{values}    // here length is defined
  var arr1 = [3]int{1,2,3}

  var array_name = [...]datatype{values}    // here length is inferred
  var arr1 = [...]int{1,2,3}
  ```
- With the := sign::
  ```
  array_name := [length]datatype{values}    // here length is defined
  arr2 := [5]int{4,5,6,7,8}
  array_name := [...]datatype{values}    // here length is inferred
  arr2 := [...]int{4,5,6,7,8}
  ```

**Array Initialization:**
```
arr1 := [5]int{} //not initialized
arr2 := [5]int{1,2} //partially initialized
arr3 := [5]int{1,2,3,4,5} //fully initialized
```

**Initialize Only Specific Elements:** It is possible to initialize only specific elements in an array.
```
arr1 := [5]int{1:10,2:40}
```

 The array above has 5 elements.

  - 1:10 means: assign 10 to array index 1 (second element).
  - 2:40 means: assign 40 to array index 2 (third element).


## Slice:
- Slices are similar to arrays, but are more powerful and flexible.
- Unlike arrays, the length of a slice can grow and shrink as you see fit.
- It has declared like Array
```
slice_name = []datatype{}

my_slice := []int{}
```
The code above declares an empty slice of 0 length and 0 capacity.

In Go, there are two functions that can be used to return the length and capacity of a slice:
  - **len()** function - returns the length of the slice (the number of elements in the slice)
  - **cap()** function - returns the capacity of the slice (the number of elements the slice can grow or shrink to)

* Create a Slice From an Array:
```
var my_array = [length]datatype{values}   // An array
my_slice := myarray[start:end]   // A slice made from the array
```

* Create a Slice With The **make()** Function:
```
slice_name := make([]type, length, capacity)
// If the capacity parameter is not defined, it will be equal to length.
```

* Append Elements To a Slice:
```
slice_name = append(slice_name, element1, element2, ...)
```

* Append One Slice To Another Slice:
```
slice3 := append(slice1, slice2...)
// Note: The '...' after slice2 is necessary when appending the elements of one slice to another.
// don't forget to use := for that
```

## Loop:
- The for loop is the only loop available in Go.
```
for statement1; statement2; statement3 {
   // code to be executed for each iteration
}
```

 The **Range** Keyword: 
- It returns both the index and the value.
  ```
  for index, value := array|slice|map {
    // code to be executed for each iteration
  }
  ```

- To only show the value or the index, you can omit the other output using an underscore (_).
  ```
  // it shows only value
  for _, value := array|slice|map {
    // code to be executed for each iteration
  }

  // it shows only index
  for index, _ := array|slice|map {
    // code to be executed for each iteration
  }
  ```

## Function:
- Create a Function:
  ```
  func FunctionName() {
    // code to be executed
  }
  ```

- Parameters and Arguments:
  ```
  func FunctionName(param1 type, param2 type, param3 type) {
    // code to be executed
  }
  ```

- Return Values:
  - If you want the function to return a value, you need to define the data type of the return value (such as int, string, etc), and also use the return keyword inside the function:
  ```
  func FunctionName(param1 type, param2 type) type {
    // code to be executed
    return output
  }
  ```

  - Named Return Values:
  ```
  func FunctionName(param1 type, param2 type) (name type) {
    // code to be executed
    name = param1 + param 2
    return
  }
  ```

## Structures:
A struct (short for structure) is used to create a collection of members of different data types, into a single variable.

While arrays are used to store multiple values of the same data type into a single variable, structs are used to store multiple values of different data types into a single variable.

A struct can be useful for grouping data together to create records.

- **Declare a Struct:** To declare a structure in Go, use the type and struct keywords:
  ```
  type struct_name struct {
    member1 datatype;
    member2 datatype;
    member3 datatype;
    ...
  }
```
