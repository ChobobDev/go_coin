package main

import "fmt"

func main() {
	foods := []string{"Pizza", "Pasta", "Noodle", "Hamburger"}
	for _, food := range foods {
		fmt.Println(food)
	}

	fmt.Printf("%v\n", foods)
	foods = append(foods, "French Fries")
	fmt.Printf("%v\n", foods)
}
