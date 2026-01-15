package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Welcome to the time study of golang")
	presentTime := time.Now()
	fmt.Println(presentTime)

	fmt.Println(presentTime.Format("01-02-2006 Monday")) //use the same no always 
}