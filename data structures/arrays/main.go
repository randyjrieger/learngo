//Array Basics
// 0 based

//Arrays are fixed. The size is part of its defintion

package main

import "fmt" //for println

func main() {
	//an int array of size 5
	var myArray [5]int

	fmt.Println("Here are the starting contents of my array: %v", myArray)
	fmt.Println("\nOk, let's add some values!")

	//Add some values to our array. Keep the values the same data type
	myArray[0] = 1
	myArray[1] = 3
	myArray[2] = 5
	myArray[3] = 7
	myArray[4] = 9

	fmt.Println("\nHere are the contents of my 'Odd Numbers' array: %v", myArray)

}
