package main

import "fmt"

func main()  {
	var name string = "John"

	age := 89
	score := 90.5

	if age > 18 {
		fmt.Println("You are an adult")
	} else {
		fmt.Println("You are a minor")
	}

	switch score {
	case 90:
		fmt.Println("You got an A")
	case 80:
		fmt.Println("You got a B")
	case 70:
		fmt.Println("You got a C")
	default:
		fmt.Println("You got a D")
	}

	fmt.Println("Hello, World!")
	fmt.Println(name, age, score)
}