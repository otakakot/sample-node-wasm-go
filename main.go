package main

import "syscall/js"

var done = make(chan struct{})

func init() {
	js.Global().Set("golog", js.FuncOf(func(_ js.Value, args []js.Value) interface{} {
		defer func() {
			done <- struct{}{}
		}()

		println(args[0].String())

		return nil
	}))
}

func main() {
	js.Global().Get("console").Call("log", "Hello, WebAssembly!")
	<-done
}
