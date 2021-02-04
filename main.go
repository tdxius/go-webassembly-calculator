package main

import (
  "syscall/js"
)

func add(this js.Value, i []js.Value) interface{} {
  return js.ValueOf(i[0].Int() + i[1].Int())
}

func sub(this js.Value, i []js.Value) interface{} {
  return js.ValueOf(i[0].Int() - i[1].Int())
}

func registerCallbacks() {
  js.Global().Set("add", js.FuncOf(add))
  js.Global().Set("sub", js.FuncOf(sub))
}

func main() {
  c := make(chan struct{}, 0)

  println("Go WebAssembly Initialized")
  registerCallbacks()

  <-c
}
