package cli

import (
	"flag"
	"fmt"
	"runtime"

	"github.com/ChobobDev/go_coin/explorer"
	"github.com/ChobobDev/go_coin/rest"
)

func usage() {
	fmt.Println("Welcome to GO lang blockchain")
	fmt.Println("Please use the following commands:")
	fmt.Println("-port=4000 : Set the PORT of the Server")
	fmt.Println("-mode=rest: Choose Between 'html' and 'rest'")
	// defer는 살리고 cli만 죽이기 위해 사용
	runtime.Goexit()
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
