package main

// fetches rfc's
import (
	"flag"
	"fmt"
	"math/rand/v2"
	"os"

	"github.com/TWolfis/ietfrfc"
)

var (
	number int
	short  bool
)

func main() {
	// generate randum number
	rnRfc := rand.IntN(10000)
	flag.IntVar(&number, "num", rnRfc, "rfc number to fetch")
	flag.BoolVar(&short, "short", false, "short output")

	// parse flags
	flag.Parse()

	rfc, err := ietfrfc.Get(number)
	if err != nil {
		fmt.Println(err)
		os.Exit(0)
	}

	if short {
		fmt.Println(rfc.Title)
		fmt.Println("RFC", rfc.Number)
		fmt.Println(rfc.Abstract)
		fmt.Println(rfc.Author)
		fmt.Println(rfc.URL)

		os.Exit(0)
	}

	fmt.Println(rfc)
}
