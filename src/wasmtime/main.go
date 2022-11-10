package main

import (
	"fmt"
	"io/ioutil"

	"github.com/bytecodealliance/wasmtime-go"
)

func main() {
	wasmBytes, _ := ioutil.ReadFile("static/wasm.wasm")
	// Almost all operations in wasmtime require a contextual `store`
	// argument to share, so create that first
	store := wasmtime.NewStore(wasmtime.NewEngine())

	// Once we have our binary `wasm` we can compile that into a `*Module`
	// which represents compiled JIT code.
	module, _ := wasmtime.NewModule(store.Engine, wasmBytes)

	// Our `hello.wat` file imports one item, so we create that function
	// here.
	item := wasmtime.WrapFunc(store, func() {
		fmt.Println("Hello from Go!")
	})

	// Next up we instantiate a module which is where we link in all our
	// imports. We've got one import so we pass that in here.
	instance, _ := wasmtime.NewInstance(store, module, []wasmtime.AsExtern{item})

	// After we've instantiated we can lookup our `run` function and call
	// it.
	pokemonStats := instance.GetFunc(store, "pokemonStats")
	if pokemonStats == nil {
		panic("not a function")
	}
	_, _ = pokemonStats.Call(store)
}
