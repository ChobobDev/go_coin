package main

import (
	"github.com/ChobobDev/go_coin/blockchain"
	"github.com/ChobobDev/go_coin/cli"
)

func main() {
	blockchain.Blockchain()
	cli.Start()
}
