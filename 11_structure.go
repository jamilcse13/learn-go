package main
import ("fmt")

type Person struct {
	name string
	age int
	job string
	salary int
}

func main() {
	var pers1 Person
	var pers2 Person

	// Pers1 specification
	pers1.name = "John Doe"
	pers1.age = 40
	pers1.job = "Teacher"
	pers1.salary = 40000
	
	// Pers2 specification
	pers2.name = "Fulan Ibn Fulan"
	pers2.age = 30
	pers2.job = "ENgineer"
	pers2.salary = 50000


	// Access and print Pers1 info
	fmt.Println("Name: ", pers1.name)
	fmt.Println("Age: ", pers1.age)
	fmt.Println("Job: ", pers1.job)
	fmt.Println("Salary: ", pers1.salary)
	fmt.Println()
	
	// Access and print Pers1 info
	fmt.Println("Name: ", pers2.name)
	fmt.Println("Age: ", pers2.age)
	fmt.Println("Job: ", pers2.job)
	fmt.Println("Salary: ", pers2.salary)
	fmt.Println()


	// Pass Struct as Function Arguments
	printPerson(pers1)
	printPerson(pers2)
}


func printPerson(pers Person) {
	fmt.Println("Name: ", pers.name)
	fmt.Println("Age: ", pers.age)
	fmt.Println("Job: ", pers.job)
	fmt.Println("Salary: ", pers.salary)
	fmt.Println()
}