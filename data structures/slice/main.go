package main

import "fmt"

func main() {
	fmt.Println("Golang. It's a Slice!")

	//Declaring an Array of customer Ids
	customerIds := [5]int{2445, 1212, 1415, 2187, 1626}

	//Declaring a Slice which will reference a section of the array
	//Here, we use 1:3, which will reference elements 1 to 2.
	//The last element, 3, is excluded.
	var selectedCustomers []int = customerIds[1:3]

	fmt.Println("Selected Customer IDs: ", selectedCustomers)
}
