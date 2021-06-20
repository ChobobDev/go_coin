package main

import (
	"crypto/sha256"
	"fmt"
)

type block struct {
	data     string
	hash     string
	prevHash string
}

func main() {
	jeongBlock := block{"Jeong dulgi Block", "", ""}
	hash := sha256.Sum256([]byte(jeongBlock.data + jeongBlock.prevHash))
	hexHash := fmt.Sprintf("%x", hash)
	jeongBlock.hash = hexHash
	fmt.Println(jeongBlock)

}
