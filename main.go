package main

import (
	"github.com/ChobobDev/go_coin/cli"
	"github.com/ChobobDev/go_coin/db"
)

func main() {
	defer db.Close()
	cli.Start()
}
