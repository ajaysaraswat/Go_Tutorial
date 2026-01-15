package main

import "fmt"

func main() {
	fmt.Println("welcome to the array in golangs")
	// define the array format
	var fruitList [4]string
	fruitList[0] = "apple"
	fruitList[1] = "tamato"
	fruitList[3] = "peach"

	fmt.Println("fruitlist is ", fruitList)
	fmt.Println("fruitlist is ", len(fruitList))
}