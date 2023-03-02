package main
import (
	"fmt"
)

func main() {
	// Print()
	var i, j string = "Hello","World"

	fmt.Print(i, j)
	
	// include new line
	fmt.Print("\n")
	fmt.Print(i, "\n")
	fmt.Print(j, "\n")
	
	fmt.Print(i, "\n", j)
	fmt.Print("\n")
	
	// add a white space
	fmt.Print(i, " ", j)
	fmt.Print("\n")
	
	// Print() inserts a space between the arguments if neither are strings:
	var a, b = 10, 20
	fmt.Print(a, b)
	fmt.Print("\n \n")
	
	
	
	// Println()
	fmt.Println(i, j)
	fmt.Println(a, b)
	fmt.Print("\n \n")
	


	// Printf()
	fmt.Printf("i has velaue: %v and type %T \n", i, i)
	fmt.Printf("a has velaue: %v and type %T", a, a)

}