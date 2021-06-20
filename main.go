package main

import (
	"fmt"

	"github.com/ChobobDev/go_coin/person"
)

func main() {
	bernie := person.Person{}
	bernie.SetDetails("Bernie", 21, "Korea")
	fmt.Println(bernie)
}
