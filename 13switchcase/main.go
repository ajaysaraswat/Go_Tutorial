package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main(){
	fmt.Println("switch case in golang")
	rand.Seed(time.Now().UnixNano())
	diceNumber := rand.Intn(6)  + 1//0-5 + 1 = 1 - 6
	fmt.Println(diceNumber)

	switch diceNumber {
	case 1:
		fmt.Println("you got 1 ")
		fallthrough // after executing this case it will go to next case
	case 2:
		fmt.Println("you got 2 ")
		fallthrough
	case 3:
		fmt.Println("you got 3 ")
		fallthrough
	case 4 : 
	    fmt.Println("you got 4 ")
		
	case 5:
		fmt.Println("you got 5 ")
		
	case 6 : 
	    fmt.Println("you got 6 ")
		
	default : 
	     fmt.Println("not dice ")
	}
}