//remember, our file must be part of a package
package main

//we need to import 2 packages
//fmt - for Println
//regexp - to allow us to work with regular expressions
import (
	"fmt"
	"regexp"
)

func main() {
	fmt.Println("Let's try matching strings against 'l([a-z]+).s'")
	match1, _ := regexp.MatchString(`l([a-z]+).s`, "lions")
	fmt.Printf("lions -> %v\n", match1)
	match2, _ := regexp.MatchString(`l([a-z]+).s`, "lillies")
	fmt.Printf("lillies -> %v\n", match2)
	match3, _ := regexp.MatchString(`l([a-z]+).s`, "lyle lovett")
	fmt.Printf("lyle lovett -> %v\n", match3)
	match4, _ := regexp.MatchString(`l([a-z]+).s`, "lit")
	fmt.Printf("lit -> %v\n", match4)
}
