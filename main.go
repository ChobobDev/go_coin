package main

import (
	"github.com/ChobobDev/go_coin/explorer"
	"github.com/ChobobDev/go_coin/rest"
)

func main() {
	go explorer.Start(3000)
	rest.Start(5000)
}
