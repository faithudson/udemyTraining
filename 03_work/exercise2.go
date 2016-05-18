package main

import "fmt"


// This is function declaration for function "half"


func main() {

	half := func(faith int) (int, bool) {
	return faith / 2, faith%2 == 0
}
	diamond, even := half(12)
	fmt.Println(diamond, even)
}