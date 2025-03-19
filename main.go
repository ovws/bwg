package main

import (
	"fmt"
	"runtime"
)

func say(s string) {
	for i := 0; i < 10; i++ {
		runtime.Gosched()
		fmt.Println(s)
	}
}

func main() {
	go say("world")
	say("hello")
	fmt.Println("main end")
}
