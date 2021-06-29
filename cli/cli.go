package cli

import (
	"flag"
	"fmt"
	"os"

	"github.com/ChobobDev/go_coin/explorer"
	"github.com/ChobobDev/go_coin/rest"
)

func usage() {
	fmt.Println("Welcome to GO lang blockchain")
	fmt.Println("Please use the following commands:")
	fmt.Println("-port=4000 : Set the PORT of the Server")
	fmt.Println("-mode=rest: Choose Between 'html' and 'rest'")
	os.Exit(0)
}

func Start() {
	port := flag.Int("port", 4000, "Set port of server")

	mode := flag.String("mode", "rest", "Choose Between 'html' and 'rest'")

	flag.Parse()

	switch *mode {
	case "rest":
		rest.Start(*port)
	case "html":
		explorer.Start(*port)
	default:
		usage()
	}

	fmt.Println(*port, *mode)
}
