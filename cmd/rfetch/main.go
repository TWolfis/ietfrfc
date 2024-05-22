package main

// fetches rfc's
import (
	"flag"
	"fmt"
	"os"

	"github.com/TWolfis/ietfrfc"
)

var (
	n int
)

func main() {
	flag.IntVar(&n, "num", 1, "rfc number to fetch")
	flag.Parse()

	rfc, err := ietfrfc.Get(n)
	if err != nil {
		fmt.Println(err)
		os.Exit(0)
	}

	fmt.Println(rfc)
}
