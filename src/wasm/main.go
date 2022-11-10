package main

import (
	"fmt"
	"syscall/js"

	"github.com/adawolfs/wasm-go/src/pokemon"
)

func main() {
	wait := make(chan struct{}, 0)
	js.Global().Set("pokemonStats", js.FuncOf(pokemonStats))
	<-wait
}

func pokemonStats(this js.Value, args []js.Value) interface{} {
	stats := pokemon.GetStats(args[0].String())
	if stats.HP == 0 {
		return "Pokemon not found"
	}
	return fmt.Sprintf("HP: %d, Atk: %d, Def: %d, Spd: %d, Spc: %d", stats.HP, stats.Atk, stats.Def, stats.Spd, stats.Spc)
}
