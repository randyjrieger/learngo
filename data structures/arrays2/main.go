//Array Basics
//Using functions to fill your Arrays
package main

import "fmt"

//(copyOfArray [5]int) - we are expecting to be passed an integer array of size 5
//[5] int - this 'extra' bit at the end of the signature says what we are returning
func populateArray(copyOfArray [5]int) [5]int {
	//Add some values to our array. Keep the values the same data type
	copyOfArray[0] = 1
	copyOfArray[1] = 3
	copyOfArray[2] = 5
	copyOfArray[3] = 7
	copyOfArray[4] = 9

	//Return the array because it was passed as a copy to this func
	return copyOfArray
}

func main() {
	//an int array of size 5
	var myArray [5]int

	fmt.Println("Here are the starting contents of my array: %v", myArray)
	fmt.Println("\nOk, let's add some values!")

	//It is important to know that when you pass an array to a function, it passes a COPY of the array, not a reference
	//So, if you change the array, the main func won't know!
	myArray = populateArray(myArray)

	fmt.Println("\nHere are the contents of my 'Odd Numbers' array: %v", myArray)
}
