package main

import "fmt"

func main() {
	fmt.Println("Learning git commands....")
	oddArray := []string{"1","3","5","7"}
	for index,odd := range oddArray {
		fmt.Printf("%v odd no is %v \n",index,odd)
	}
}
