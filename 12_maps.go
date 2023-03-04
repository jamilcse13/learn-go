package main
import ("fmt")

func main() {
	var a = map[string]string{"brand": "Ford", "model": "Mustang", "year": "1964"}
	b := map[string]int{"Oslo": 1, "Bergen": 2, "Trondheim": 3, "Stavanger": 4}

	fmt.Printf("a\t%v\n", a)
	fmt.Printf("b\t%v\n", b)
	fmt.Println()


	// Creating Maps Using make()Function:
	var c = make(map[string]string)		// The map is empty now
	var d = make(map[string]int)		// The map is empty now
	
	c["name"] = "John Doe"
	c["country"] = "Argentina"
	c["language"] = "Spanish"

	d["age"] = 33
	d["salary"] = 1200

	fmt.Printf("c\t%v\n", c)
	fmt.Printf("d\t%v\n", d)
	fmt.Println()


	// Creating an Empty Map
	var e = make(map[string]string)
	var f map[string]string

	fmt.Println(e)
	fmt.Println(e == nil)
	fmt.Println(f)
	fmt.Println(f == nil)
	fmt.Println()


	// Accessing Map Elements
	fmt.Println(c["name"])

	c["name"] = "Fulan"
	c["marital_status"] = "Married"

	fmt.Println(c)


	// Remove Element from Map:
	delete(c, "marital_status")
	fmt.Println(c)
	fmt.Println()


	var z = map[string]string{"brand": "Ford", "model": "Mustang", "year": "1964", "day": ""}
	// Check For Specific Elements in a Map
	val1, ok1 := z["brand"]
	val2, ok2 := z["color"]
	val3, ok3 := z["day"]
	_, ok4 := z["model"]

	fmt.Println(val1, ok1)
	fmt.Println(val2, ok2)
	fmt.Println(val3, ok3)
	fmt.Println(ok4)
	fmt.Println()

	y := z
	fmt.Println(z)
	fmt.Println(y)
	
	y["year"] = "1977"
	fmt.Println("After change to y:")

	fmt.Println(z)
	fmt.Println(y)
	fmt.Println()
	

	// Iterating Over Maps
	g := map[string]int{"two": 2, "one": 1, "three": 3, "four": 4}

	for k, v := range g {
		fmt.Printf("%v : %v, ", k, v)
	}
	fmt.Println()
	

	// Iterate Over Maps in a Specific Order
	var order = []string{}	// defining the order
	order = append(order, "one", "two", "three", "four")

	for k, v := range g {    // loop with no order
		fmt.Printf("%v : %v, ", k, v)
	}
	fmt.Println()

	for _, element := range order {   // loop with the defined order
		fmt.Printf("%v : %v, ", element, g[element])
	}
}