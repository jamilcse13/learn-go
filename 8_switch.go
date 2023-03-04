package main
import ("fmt")

func main() {

	day := 3
	switch day {
	case 1:
		fmt.Println("Saturday")
	case 2:
		fmt.Println("Sunday")
	case 3:
		fmt.Println("Monday")
	case 4:
		fmt.Println("Tuesday")
	case 5:
		fmt.Println("Wednesday")
	case 6:
		fmt.Println("Thursday")
	case 7:
		fmt.Println("Friday")
	default:
		fmt.Println("Not a Weekday")
	}


	// The Multi-case
	switch day {
	case 1,3,5:
		fmt.Println("Odd weekday")
	case 2,4:
		fmt.Println("Even weekday")
	case 6,7:
		fmt.Println("Weekend")
	default:
		fmt.Println("Invalid day of day number")
	}
	
}