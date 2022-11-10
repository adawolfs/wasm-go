package main

import (
	"fmt"
	"io/ioutil"

	wasmer "github.com/wasmerio/wasmer-go/wasmer"
)

func main() {
	wasmBytes, _ := ioutil.ReadFile("static/wasm.wasm")

	engine := wasmer.NewEngine()
	store := wasmer.NewStore(engine)

	// Compiles the module
	module, _ := wasmer.NewModule(store, wasmBytes)

	// Instantiates the module
	importObject := wasmer.NewImportObject()
	instance, _ := wasmer.NewInstance(module, importObject)

	// Gets the `sum` exported function from the WebAssembly instance.
	pokemonStats, _ := instance.Exports.GetFunction("pokemonStats")

	// Calls that exported function with Go standard values. The WebAssembly
	// types are inferred and values are casted automatically.
	result, _ := pokemonStats('c')

	fmt.Println(result)
}
