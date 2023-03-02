// var var_name type = val
package main
import (
	"fmt"
)

func main() {
	var studentId int = 101   // type is declared
	var studentName = "Student" 	// type is inferred
	studentDept := "CSE" // type is inferred and can not be declared without val

	fmt.Println(studentId, studentName, studentDept)


	// Variable Declaration Without Initial Value
	var a string 	// default val = ""
	var b int 	// default val = 0
	var c bool 	// default val = false
	
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)


	// Value Assignment After Declaration
	var teacherName string
	teacherName = "John Doe"

	fmt.Println(teacherName)


	// Go Multiple Variable Declaration
	// If you use the type keyword, it is only possible to declare one type of variable per line.
	var q, w, e, r int = 1, 5, 6, 10

	fmt.Println(q,w,e,r)

	// If you are not used the type keyword, it is possible to declare multi type of variable per line.
	var p, o = 1, "po"
	l, k := 10, "LM"

	fmt.Println(p, o, l, k)

	// Go Variable Declaration in a Block for greater readability
	var (
		az int
		ax int = 30
		ac string = "LM"
	)

	fmt.Println(az, ax, ac)

}