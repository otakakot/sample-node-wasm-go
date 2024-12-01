package main

import "syscall/js"

var done = make(chan struct{})

func golog(_ js.Value, args []js.Value) any {
	defer func() {
		done <- struct{}{}
	}()

	println(args[0].String())

	js.Global().Get("console").Call("log", args[0].String())

	return nil
}

func init() {
	js.Global().Set("golog", js.FuncOf(golog))
}

func main() {
	<-done
}
