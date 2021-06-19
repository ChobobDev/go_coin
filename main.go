package main

import "fmt"

type person struct {
	name        string
	age         int
	nationality string
}

func (p person) greeting() {
	fmt.Printf("Hello! My name is %s and I'm %d yers old, and I'm from %s", p.name, p.age, p.nationality)

}

func main() {
	bernie := person{"Bernie", 21, "Korea"}
	bernie.greeting()
}
