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

* Go Multiple Variable Declaration:
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
