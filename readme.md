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