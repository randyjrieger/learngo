package main

import (
	"fmt"
	"time"
)

func counting(name string) {
	for x := 1; x <= 3; x++ {
		fmt.Println(name, ": says", x)
	}
}

func main() {
	//Synchronously Calling A func
	go counting("Asynchronous")

	fmt.Println("I'm not waiting for you!")

	time.Sleep(time.Second)
}
