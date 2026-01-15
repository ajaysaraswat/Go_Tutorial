package main

import "fmt"

func main() {
	fmt.Println("learning defer statement")
	defer fmt.Println("1......")
	fmt.Println("2......")
	defer fmt.Println("3......")
}