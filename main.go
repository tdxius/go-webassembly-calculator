package main

import (
  "strconv"
  "syscall/js"
)

func add(this js.Value, i []js.Value) interface{} {
  value1 := js.Global().Get("document").Call("getElementById", i[0].String()).Get("value").String()
  value2 := js.Global().Get("document").Call("getElementById", i[1].String()).Get("value").String()

  int1, _ := strconv.Atoi(value1)
  int2, _ := strconv.Atoi(value2)

  js.Global().Get("document").Call("getElementById", i[2].String()).Set("value", int1 + int2)

  return true
}

func subtract(this js.Value, i []js.Value) interface{} {
  value1 := js.Global().Get("document").Call("getElementById", i[0].String()).Get("value").String()
  value2 := js.Global().Get("document").Call("getElementById", i[1].String()).Get("value").String()

  int1, _ := strconv.Atoi(value1)
  int2, _ := strconv.Atoi(value2)

  js.Global().Get("document").Call("getElementById", i[2].String()).Set("value", int1 - int2)

  return true
}

func registerCallbacks() {
  js.Global().Set("add", js.FuncOf(add))
  js.Global().Set("subtract", js.FuncOf(subtract))
}

func main() {
  c := make(chan struct{}, 0)

  println("Go WebAssembly Initialized")
  registerCallbacks()

  <-c
}
