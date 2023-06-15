package main

import (
	"fmt"
	"runtime"
)

func main() {
	runtime.GOMAXPROCS(2)

	messages := make(chan string)

	sayHelloTo := func(who string) {
		data := fmt.Sprintf("Hello %s", who)
		messages <- data
	}

	go sayHelloTo("john wick")
	go sayHelloTo("ethan hunt")
	go sayHelloTo("jason bourne")

	message1 := <-messages
	fmt.Println(message1)

	message2 := <-messages
	fmt.Println(message2)

	message3 := <-messages
	fmt.Println(message3)
}
