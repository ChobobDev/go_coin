package main

import (
	"github.com/ChobobDev/go_coin/cli"
	"github.com/ChobobDev/go_coin/db"
)

func main() {
	//nothing for now
	defer db.Close()
	cli.Start()
}
