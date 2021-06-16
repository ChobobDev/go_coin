package main

import "fmt"

// const name string = "Bernie"
// const age int = 21

func addtion(a int, b int) int {
	return a + b
}

func something(number int, name string) (int, string) {
	return number, name
}

func popo(numbers ...int) int {
	var total int
	for _, item := range numbers {
		total += item
	}
	return total
}

func main() {
	// var fr_name string = "Jeong"
	// var species string = "gugu"
	fmt.Println("Blockchain study with Go lang")
	// fmt.Printf("name: %s age:%d \n", name, age)
	// fmt.Printf("이름 : %s 종: %s \n", fr_name, species)
	fmt.Println(addtion(1, 2))
	fmt.Println(something(1, "bernie"))
	sums := popo(2, 3, 4, 5, 6, 7)

	fmt.Println(sums)
}
