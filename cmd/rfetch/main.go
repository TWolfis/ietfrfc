package main

// fetches rfc's
import (
	"flag"
	"fmt"
	"log"
	"math/rand/v2"

	"github.com/TWolfis/ietfrfc"
)

var (
	n int
)

func main() {
	flag.IntVar(&n, "n", 0, "rfc number")
	flag.Parse()

	if n == 0 {
		n = rand.IntN(10000)
	}

	rfc, err := ietfrfc.Get(n)
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println(rfc)
}
