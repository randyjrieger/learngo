package main

import "fmt"

// this function is accessible from main() because
// it is declared in the same package as main()
func partingmessage() {
	fmt.Println("Text me later!\nPeace!")
}
