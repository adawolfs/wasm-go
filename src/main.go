// Go http server
package main

import (
	"fmt"
	"os"

	"github.com/adawolfs/wasm-go/src/pokemon"
)

func main() {
	args := os.Args[1:]
	if len(args) == 0 {
		fmt.Println("Please provide a pokemon name")
		return
	}
	stats := pokemon.GetStats(args[0])
	fmt.Printf("%s: %+v\n", args[0], stats)
}
