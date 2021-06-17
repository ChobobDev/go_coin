package main

import "fmt"

func main() {
	x := 16
	fmt.Printf("%b \n", x)
	fmt.Printf("%o \n", x)
	fmt.Printf("%x \n", x)
	fmt.Printf("%U \n", x)

	xAsBinary := fmt.Sprintf("%b", x)
	xAsOct := fmt.Sprintf("%o", x)
	xAsHex := fmt.Sprintf("%x", x)
	xAsUnicode := fmt.Sprintf("%U", x)

	fmt.Println(x, xAsBinary)
	fmt.Println(x, xAsOct)
	fmt.Println(x, xAsHex)
	fmt.Println(x, xAsUnicode)
}
