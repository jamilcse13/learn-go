package main
import ("fmt")

func main() {
	for i:=0; i<5; i++ {
		if i == 3 {
			break
			// continue
		}
		fmt.Println(i)
	}


	// Nested Loops
	adj := [2]string{"big", "testy"}
	fruits := [3]string{"apple", "orange", "banana"}
	for i:=0; i<len(adj); i++ {
		for j:=0; j<len(fruits); j++ {
			fmt.Println(adj[i], fruits[j])
		}
	}


	// The Range Keyword
	for ind, val := range fruits {
		fmt.Printf("%v\t%v\n", ind, val)
	}

	// show only value
	for _, val := range fruits {
		fmt.Printf("%v\n", val)
	}

	// show only index
	for ind, _ := range fruits {
		fmt.Printf("%v\n", ind)
	}
}