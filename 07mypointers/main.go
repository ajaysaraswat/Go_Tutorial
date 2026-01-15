package main

import "fmt"

func main() {
	fmt.Println("welcome to the pointers study of golang")
// 	var ptr *int
// 	fmt.Println("value of ptr is:", ptr)
	num := 23
	var ptr = &num
	fmt.Println("value of actual pointer is ",ptr)
	fmt.Println("value of actual pointer is ",*ptr)
	*ptr = *ptr + 2
	fmt.Println(num)
   }