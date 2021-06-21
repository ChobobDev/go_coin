package main

import (
	"fmt"

	"github.com/ChobobDev/go_coin/blockchain"
)

func main() {
	chain := blockchain.GetBlockchain()
	chain.AddBlock("Cho Block")
	chain.AddBlock("Kim Block")
	chain.AddBlock("Lee Block")
	chain.AddBlock("Park Block")
	chain.AddBlock("Chae Block")
	chain.AddBlock("Lim Block")
	chain.AddBlock("Ham Block")
	chain.AddBlock("Oh Block")

	for _, block := range chain.AllBlocks() {
		fmt.Printf("Data : %s \n", block.Data)
		fmt.Printf("Hash : %s \n", block.Hash)
		fmt.Printf("Prev Hash : %s \n", block.PrevHash)
		fmt.Println("")
	}
}
