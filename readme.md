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
