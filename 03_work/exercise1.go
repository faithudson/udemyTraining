package main

import "fmt"


// This is function declaration for function "half"



func half(faith int) (int, bool) {
	return faith / 2, faith%2 == 0
}

func main() {
	diamond, even := half(12)
	fmt.Println(diamond, even)
}