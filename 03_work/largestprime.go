package main

import "fmt"

func main() {
    small := 1
    big := 9
//    current_highest := -1 

    for i := small; i <= big; i++ {
         // This for loop will loop over 9 times. starting at 1 and ending at 9. and i is the variable here. 
    	// Which is why it prints 9 times. Do you get that? i do
        //fmt.Println(i)

    	// Do you get why this only prints 4 times? 2,4,6,8
        if i%2 == 0 {
	        fmt.Println(i)
        }

    }
}