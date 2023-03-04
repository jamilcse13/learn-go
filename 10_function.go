package main
import ("fmt")

func myMessage() {
	fmt.Println("I just got executed!")
}


// Parameters and Arguments
func familyName(fname string, age int) {
	fmt.Printf("Hello %d year old %s Broad\n", age, fname)
}


// Go Function Returns
func myFunction(x int, y int) int {
	return x + y
}


// Named Return Values
func myFunction1(x int, y int) (result int) {
	result = x + y
	// it will work like: return result
	return
}


// Multiple Return Values
func myFunction2(x int, y int, msg string) (result int, text string) {
	result = x + y
	text = msg + " World"
	return
}


// Recursion Functions
func testCount(x int) int {
	if x == 11 {
		return 0
	}

	fmt.Println(x)
	return testCount(x+1)
}


func factorial_recursion(x float64) (y float64) {
	if x > 0 {
		y = x * factorial_recursion(x-1)
	} else {
		y = 1
	}

	return
}


func main() {
	myMessage()
	familyName("Stuart", 10)
	familyName("Alex", 3)
	
	fmt.Println(myFunction(10, 3))
	fmt.Println(myFunction1(20, 30))
	
	total := myFunction(15, 25)
	fmt.Println(total)
	
	fmt.Println(myFunction2(3, 8, "Hello"))
	// store the value and then return
	a, b := myFunction2(3, 8, "Hello")
	fmt.Println(a,b)

	_, d := myFunction2(3, 8, "Hello")
	fmt.Println(d)
	fmt.Println()

	// recursion
	testCount(1)
	fmt.Println()

	fmt.Println(factorial_recursion(5))

}