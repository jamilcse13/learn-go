package main
import ("fmt")

var a = 20;
var b = 25;

func main() {

	if a > b {
		fmt.Printf("%d is greated than %d\n", a, b)
	} else {
		fmt.Printf("%d is greated than %d\n", b, a)
	}

	time := 8

	if time <= 10 {
		fmt.Println("Good Morning")
	} else if time <= 20 {
		fmt.Println("Good Day")		
	} else {
		fmt.Println("Good Evening")		
	}
}