package main

import "github.com/ChobobDev/go_coin/blockchain"

func main() {
	chain := blockchain.GetBlockchain()
	chain.addBlock("Jeong Block")
	chain.addBlock("Cho Block")
	chain.addBlock("Kim Block")
	chain.addBlock("Lee Block")
	chain.addBlock("Park Block")
	chain.addBlock("Chae Block")
	chain.addBlock("Lim Block")
	chain.addBlock("Ham Block")
	chain.addBlock("Oh Block")
	chain.listBlocks()

}
